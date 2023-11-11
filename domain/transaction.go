package domain

type Transaction struct {
	Id              string  `json:"id,omitempty"`
	MaterialId      string  `json:"material_id,omitempty"`
	Project         string  `json:"project,omitempty"`
	Company         string  `json:"company,omitempty"`
	TransactionType string  `json:"transaction_type,omitempty"`
	ReferenceNo     string  `json:"reference_no,omitempty"`
	Units           float32 `json:"units,omitempty"`
	Quantity        float32 `json:"quantity,omitempty"`
	Received        float32 `json:"received,omitempty"`
	Issued          float32 `json:"issued,omitempty"`
	Returns         float32 `json:"return,omitempty"`
	Balance         float32 `json:"balance,omitempty"`
	Source          string  `json:"source,omitempty"`
	Destination     string  `json:"destination,omitempty"`
	TinNumber       string  `json:"tin_number,omitempty"`
	Status          string  `json:"status,omitempty"`
	UnitPrice       float32 `json:"unit_price,omitempty"`
	TotalPrice      float32 `json:"total_price,omitempty"`
	Cumulative      float32 `json:"cumulative,omitempty"`
	Remark          string  `json:"remark,omitempty"`
	CreatedAt       string  `json:"created_at,omitempty"`
	UpdatedAt       string  `json:"updated_at,omitempty"`
	DeletedAt       string  `json:"deleted_at,omitempty"`
}

func CalculateBalance(lastBalance, currentReceived, currentReturned, currentIssued float32) float32 {
	if currentReceived > currentIssued {
		return lastBalance + currentReceived + currentReturned
	}
	return lastBalance - currentIssued
}

func CalculateUnitPrice(transactionType string, unitPrice, cumulative float32) float32 {
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
	currentBalance float32,
) float32 {
	if lastUnitPrice == 0 {
		return currentUnitPrice
	}
	return ((lastUnitPrice * lastBalance) + ((currentReceived + currentReturned) * currentUnitPrice)) / currentBalance
}
