package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type YourStruct struct {
	DATE        string `csv:"DATE,omitempty"`
	REFERENCE   string `csv:"REFERENCE,omitempty"`
	UNIT_PRICE  string `csv:"UNITPRICE,omitempty"`
	RECEIVED    string `csv:"RECEIVED,omitempty"`
	ISSUED      string `csv:"ISSUED,omitempty"`
	RETURN      string `csv:"RETURN,omitempty"`
	BALANCE     string `csv:"BALANCE,omitempty"`
	TOTAL_PRICE string `csv:"TOTALPRICE,omitempty"`
	REMARK      string `csv:"REMARK,omitempty"`
}

func TestRead(t *testing.T) {
	// Create a CSVReader for the test file
	csvReader := NewCSVReader("/home/dera/Downloads/alumuniumprofileforign.csv")

	// Call the Read function
	result, err := ReadCSV[YourStruct](csvReader)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}
