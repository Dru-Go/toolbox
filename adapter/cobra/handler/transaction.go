package handler

import (
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/dru-go/noah-toolbox/usecase"
)

type TransactionHandler struct {
	Usecase usecase.TransactionUsecase
}

func NewTransaction(usecase usecase.TransactionUsecase) TransactionHandler {
	// here we provide the options for Transactions
	// 1. Create Transaction
	return TransactionHandler{
		Usecase: usecase,
	}
}

func (TransactionHandler) GetLastTransaction(transactionId string) domain.Transaction {
	return domain.Transaction{}
}
func (TransactionHandler) BulkCreate() []domain.Transaction {
	return []domain.Transaction{}
}
func (TransactionHandler) ComputeTransactionPrice(lastTransaction, currentTransaction domain.Transaction) domain.Transaction {
	return domain.Transaction{}
}
