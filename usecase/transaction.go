package usecase

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/jedib0t/go-pretty/v6/table"
)

type TransactionUsecase struct {
	Repo repository.ITransactionRepository
	Ctx  context.Context
}

type ITransactionUsecase interface {
	LoadCSV(filepath string) ([]domain.Transaction, error)
	BulkImport([]domain.Transaction) error
	BulkCompute([]domain.Transaction) []domain.Transaction
	BulkUpdate([]domain.Transaction) error
	LoadTransactions(filepath string)
}

// PrintTable prints a table of transactions using go-pretty table.
func (t TransactionUsecase) printTable(transactions []domain.Transaction) {
	// Create a new table
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)

	// Define the table header
	tw.AppendHeader(table.Row{
		"MATERIAL_ID",
		"TRANSACTION_TYPE",
		"REFERENCE",
		"RECEIVED",
		"ISSUED",
		"RETURNED",
		"BALANCE",
		"UNIT_PRICE",
		"TOTAL_PRICE",
		"CUMULATIVE",
	})

	// Populate the table body with data from the transactions array
	for _, transaction := range transactions {
		// Add data to the table body
		tw.AppendRow([]interface{}{
			transaction.MaterialId,
			transaction.TransactionType,
			transaction.ReferenceNo,
			transaction.Received,
			transaction.Issued,
			transaction.Returns,
			transaction.Balance,
			transaction.UnitPrice,
			transaction.TotalPrice,
			transaction.Cumulative,
		})
	}

	// Set table style
	tw.SetStyle(table.StyleBold)

	// Render the table
	tw.Render()
}

func (t TransactionUsecase) LoadCSV(filepath string) ([]domain.Transaction, error) {
	csvReader := repository.NewCSVReader(filepath)
	transactions, err := repository.ReadCSV[domain.Transaction](csvReader)
	if err != nil {
		return []domain.Transaction{}, fmt.Errorf("unable to parse the csv file, %s", err)
	}
	t.printTable(transactions)
	return transactions, nil
}

func (tr TransactionUsecase) BulkImport(transactions []domain.Transaction) error {
	err := tr.Repo.BulkCreate(transactions)
	if err != nil {
		return fmt.Errorf("unable to bulk create transaction, %s ", err)
	}
	return nil
}
func sanitize(value string) string {
	value = strings.Trim(value, " ")
	if value == "" {
		return "0"
	}
	return value
}
func getTransactionType(t domain.Transaction) string {
	var transactionType = t.TransactionType
	if transactionType == "" {
		transactionType = "CREDIT"
		if t.Received+t.Returns < t.Issued {
			transactionType = "DEBIT"
		}
	}
	return transactionType
}

func (tr TransactionUsecase) FindLinkedTransaction(transactionId, materialId string) error {
	var transactions []domain.Transaction
	var err error
	if transactionId != "" {
		transactions, err = tr.Repo.FetchSubsequentTransactions(transactionId)
		if err != nil {
			return fmt.Errorf("error Fetching transactions with transaction id %s, %v", transactionId, err)
		}
	}
	if materialId != "" {
		transactions, err = tr.Repo.Fetch(domain.ComputeFilter{MaterialId: transactionId})
		if err != nil {
			return fmt.Errorf("error Fetching transactions with transaction id %s, %v", transactionId, err)
		}
	}
	tr.printTable(transactions)
	return nil
}

func (tr TransactionUsecase) BulkUpdate(transactions []domain.Transaction) error {
	return tr.Repo.BulkUpdate(transactions)
}
func (tr TransactionUsecase) BulkCompute(transactions []domain.Transaction) []domain.Transaction {
	if transactions[0].Balance == 0 {
		transactions[0].Balance = transactions[0].Received
	}
	transactions[0].Cumulative = sanitize(transactions[0].UnitPrice)
	transactions[0].TransactionType = getTransactionType(transactions[0])
	transactions[0].UnitPrice = sanitize(transactions[0].UnitPrice)
	oldUnitPrice, err := strconv.ParseFloat(sanitize(transactions[0].UnitPrice), 64)
	if err != nil {
		panic(err)
	}
	transactions[0].TotalPrice = fmt.Sprintf("%.2f", domain.CalculateTotalPrice(
		transactions[0].Received, transactions[0].Returns, transactions[0].Issued, oldUnitPrice))
	for i := 1; i < len(transactions); i++ {
		v := transactions[i]
		unitPrice, err := strconv.ParseFloat(sanitize(v.UnitPrice), 64)
		if err != nil {
			panic(err)
		}
		lastUnitPrice, err := strconv.ParseFloat(sanitize(transactions[i-1].UnitPrice), 64)
		if err != nil {
			log.Fatalf("issues converting string to float for lastUnitPrice, %v", err)
		}
		lastCumulative, err := strconv.ParseFloat(sanitize(transactions[i-1].Cumulative), 64)
		if err != nil {
			log.Fatalf("issues converting string to float for lastCumulative, %v", err)
		}
		var lastBalance = transactions[i-1].Balance
		transactions[i].Balance = domain.CalculateBalance(lastBalance, v.Received, v.Returns, v.Issued)
		transactions[i].UnitPrice = fmt.Sprint(
			domain.CalculateUnitPrice(
				getTransactionType(v),
				unitPrice,
				lastCumulative,
			),
		)

		newUnitPrice, err := strconv.ParseFloat(transactions[i].UnitPrice, 64)
		if err != nil {
			panic(err)
		}
		transactions[i].TotalPrice = fmt.Sprintf("%.2f",
			domain.CalculateTotalPrice(
				v.Received,
				v.Returns,
				v.Issued,
				newUnitPrice,
			),
		)
		transactions[i].Cumulative = fmt.Sprint(
			domain.CalculateWeightedAverage(
				getTransactionType(v),
				lastUnitPrice,
				float64(lastBalance),
				float64(v.Received),
				float64(v.Returns),
				domain.CalculateUnitPrice(
					getTransactionType(v),
					unitPrice,
					lastCumulative,
				),
				float64(transactions[i].Balance),
			),
		)
		transactions[i].TransactionType = getTransactionType(transactions[i])
	}
	tr.printTable(transactions)

	return transactions
}
