package domain

import (
	"math"
)

type Transaction struct {
	Id              string `csv:"ID,omitempty"`
	MaterialId      string `csv:"MATERIAL_ID,omitempty"`
	Project         string `csv:"PROJECT,omitempty"`
	Company         string `csv:"COMPANY,omitempty"`
	TransactionType string `csv:"TRANSACTION_TYPE,omitempty"`
	ReferenceNo     string `csv:"REFERENCE,omitempty"`
	Units           int    `csv:"UNITS,omitempty"`
	Quantity        int    `csv:"QUANTITY,omitempty"`
	Received        int    `csv:"RECEIVED,omitempty"`
	Issued          int    `csv:"ISSUED,omitempty"`
	Returns         int    `csv:"RETURN,omitempty"`
	Balance         int    `csv:"BALANCE,omitempty"`
	Source          string `csv:"SOURCE,omitempty"`
	Destination     string `csv:"DESTINATION,omitempty"`
	TinNumber       string `csv:"TIN_NUMBER,omitempty"`
	Status          string `csv:"STATUS,omitempty"`
	UnitPrice       string `csv:"UNIT_PRICE,omitempty"`
	TotalPrice      string `csv:"TOTAL_PRICE,omitempty"`
	Cumulative      string `csv:"CUMULATIVE,omitempty"`
	Remark          string `csv:"REMARK,omitempty"`
	CreatedAt       string `csv:"CREATED_AT,omitempty"`
	UpdatedAt       string `csv:"UPDATED_AT,omitempty"`
	DeletedAt       string `csv:"DELETED_AT,omitempty"`
}

type Transactions []Transaction

type DateFilter struct {
	StartDate string
	EndDate   string
}
type IDS []string

type ComputeFilter struct {
	Id           string
	MaterialId   string
	Company      string
	Project      string
	Ids          IDS
	Date         DateFilter
	Transactions []Transaction
}

func CalculateTotalPrice(currentReceived, currentReturned, currentIssued int, currentUnitPrice float64) float64 {
	if currentReceived > 0 || currentReturned > 0 {
		return float64(currentReceived+currentReturned) * currentUnitPrice
	}
	return float64(currentIssued) * currentUnitPrice
}

func CalculateBalance(lastBalance, currentReceived, currentReturned, currentIssued int) int {
	if currentReceived > currentIssued || currentReturned > currentIssued {
		return lastBalance + currentReceived + currentReturned
	}
	return lastBalance - currentIssued
}

func CalculateUnitPrice(transactionType string, unitPrice, lastCumulative float64) float64 {
	if unitPrice == 0 {
		return lastCumulative
	}
	if unitPrice != lastCumulative && transactionType == "DEBIT" {
		return lastCumulative
	}
	return unitPrice
}

func CalculateWeightedAverage(
	transactionType string,
	lastUnitPrice,
	lastBalance,
	currentReceived,
	currentReturned,
	currentUnitPrice,
	currentBalance float64,
) float64 {
	if lastUnitPrice == 0 || transactionType == "DEBIT" {
		if lastUnitPrice == currentUnitPrice {
			return lastUnitPrice
		}
		return currentUnitPrice
	}
	return roundFloat(((lastUnitPrice*lastBalance)+((currentReceived+currentReturned)*currentUnitPrice))/currentBalance, 2)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
