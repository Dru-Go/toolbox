package domain

type Transaction struct {
	Id              string `json:"id,omitempty"`
	MaterialId      string `json:"material_id,omitempty"`
	Project         string `json:"project,omitempty"`
	Company         string `json:"company,omitempty"`
	TransactionType string `json:"transaction_type,omitempty"`
	ReferenceNo     string `json:"reference_no,omitempty"`
	Units           int    `json:"units,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	Received        int    `json:"received,omitempty"`
	Issued          int    `json:"issued,omitempty"`
	Returns         int    `json:"return,omitempty"`
	Balance         int    `json:"balance,omitempty"`
	Source          string `json:"source,omitempty"`
	Destination     string `json:"destination,omitempty"`
	TinNumber       string `json:"tin_number,omitempty"`
	Status          string `json:"status,omitempty"`
	UnitPrice       string `json:"unit_price,omitempty"`
	TotalPrice      string `json:"total_price,omitempty"`
	Cumulative      string `json:"cumulative,omitempty"`
	Remark          string `json:"remark,omitempty"`
	CreatedAt       string `json:"created_at,omitempty"`
	UpdatedAt       string `json:"updated_at,omitempty"`
	DeletedAt       string `json:"deleted_at,omitempty"`
}

type DateFilter struct {
	StartDate string
	EndDate   string
}
type IDS []string

type ComputeFilter struct {
	MaterialId   string
	Ids          IDS
	Date         DateFilter
	Transactions []Transaction
}

func CalculateBalance(lastBalance, currentReceived, currentReturned, currentIssued float64) float64 {
	if currentReceived > currentIssued {
		return lastBalance + currentReceived + currentReturned
	}
	return lastBalance - currentIssued
}

func CalculateUnitPrice(transactionType string, unitPrice, cumulative float64) float64 {
	if transactionType == "DEBIT" {
		return cumulative
	}
	return unitPrice
}

func CalculateWeightedAverage(
	transaction_type string,
	lastUnitPrice,
	lastBalance,
	currentReceived,
	currentReturned,
	currentUnitPrice,
	currentBalance float64,
) float64 {
	if lastUnitPrice == 0 {
		return currentUnitPrice
	}
	return ((lastUnitPrice * lastBalance) + ((currentReceived + currentReturned) * currentUnitPrice)) / currentBalance
}
