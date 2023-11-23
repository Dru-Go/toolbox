package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/dru-go/noah-toolbox/domain"
	"github.com/dru-go/noah-toolbox/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Fetch(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}

	type args struct {
		filter domain.ComputeFilter
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Testing with nil filters",
			args: args{
				filter: domain.ComputeFilter{},
			},
		},
		{
			name: "Testing with ids filters",
			args: args{
				filter: domain.ComputeFilter{
					Ids: []string{"0ef10c8d-1404-4dfe-baf9-23ccfa5f5af00ef10c8d-1404-4dfe-baf9-23ccfa5f5af0"},
				},
			},
		},
		{
			name: "Testing with materialId filters",
			args: args{
				filter: domain.ComputeFilter{
					MaterialId: "CN3952",
				},
			},
		},
		{
			name: "Testing with date filters",
			args: args{
				filter: domain.ComputeFilter{
					Date: domain.DateFilter{
						StartDate: "2023-10-18 13:11:40",
						EndDate:   "2023-11-04 04:58:57",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Db: db,
			}
			_, err := repo.Fetch(tt.args.filter)
			assert.Nil(t, err)
		})
	}
}

func TestRepository_LastTransactionFetchSubsequentTransactions(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Get Subsequent Transaction: Test with id",
			args: args{id: "19c7b88e-ad80-427f-a15f-826f0b445294"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Db: db,
			}
			got, err := repo.FetchSubsequentTransactions(tt.args.id)
			fmt.Println(utils.PrettyPrint(got))
			assert.Equal(t, len(got), 3)
			assert.Nil(t, err)
		})
	}
}

func TestRepository_BulkCreate(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type args struct {
		transactions []domain.Transaction
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test batch transaction",
			args: args{
				transactions: []domain.Transaction{
					{
						Id:              "4dcf8cda-f2d2-4b8d-a386-7d6fdb8071bk",
						MaterialId:      "CN000019",
						Project:         "cd0a2fdc-d4d0-4d1c-b6c2-d699eb178ede",
						Company:         "78030b4b-2155-4991-be9d-68f200c5a99d",
						TransactionType: "CREDIT",
						ReferenceNo:     "cd0a2fdc",
						Received:        30,
						Balance:         0,
						Source:          "Architecture",
						Destination:     "Project #1",
						UnitPrice:       "900",
						TotalPrice:      "0",
						Cumulative:      "999",
						CreatedAt:       "2023-11-04 03:53:42",
					},
					{
						Id:              "4dcf8cda-f2d2-4b8d-a386-7d6fdb8071bl",
						MaterialId:      "CN000019",
						Project:         "cd0a2fdc-d4d0-4d1c-b6c2-d699eb178ede",
						Company:         "78030b4b-2155-4991-be9d-68f200c5a99d",
						TransactionType: "CREDIT",
						ReferenceNo:     "cd0a2fdd",
						Received:        30,
						Balance:         0,
						Source:          "Architecture",
						Destination:     "Project #1",
						UnitPrice:       "908",
						TotalPrice:      "0",
						Cumulative:      "999",
						CreatedAt:       "2023-11-04 03:53:42",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Db: db,
			}
			err := repo.BulkCreate(tt.args.transactions)
			assert.Nil(t, err)
		})
	}
}

func TestRepository_BulkUpdate(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type args struct {
		transactions []domain.Transaction
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Testing bulk update",
			args: args{
				transactions: []domain.Transaction{
					{
						Id:              "15523e03-d2b7-424f-821c-63441fb1b74d",
						MaterialId:      "EL7274",
						Project:         "cd0a2fdc-d4d0-4d1c-b6c2-d699eb178ede",
						Company:         "78030b4b-2155-4991-be9d-68f200c5a99d",
						TransactionType: "CREDIT",
						ReferenceNo:     "cd0a2fdc",
						Received:        30,
						Balance:         10,
						Source:          "Architecture",
						Destination:     "Project #1",
						UnitPrice:       "900",
						TotalPrice:      "0",
						Cumulative:      "999",
						CreatedAt:       "2023-11-04 03:53:42",
					},
					{
						Id:              "199668b0-dfdb-4862-95b1-48ed1dfe9642",
						MaterialId:      "EL7274",
						Project:         "cd0a2fdc-d4d0-4d1c-b6c2-d699eb178ede",
						Company:         "78030b4b-2155-4991-be9d-68f200c5a99d",
						TransactionType: "CREDIT",
						ReferenceNo:     "cd0a2fdd",
						Received:        30,
						Balance:         20,
						Source:          "Architecture",
						Destination:     "Project #1",
						UnitPrice:       "908",
						TotalPrice:      "0",
						Cumulative:      "999",
						CreatedAt:       "2023-11-04 03:53:42",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := Repository{
				Db: db,
			}
			err := repo.BulkUpdate(tt.args.transactions)
			assert.Nil(t, err)
		})
	}
}
