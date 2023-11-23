package usecase

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/stretchr/testify/assert"
)

func TestTransactionUsecase_LoadCSV(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type fields struct {
		Repo repository.ITransactionRepository
		Ctx  context.Context
	}
	type args struct {
		filepath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Testing load transaction csv file",
			fields: fields{
				Repo: repository.NewRepository(db),
				Ctx:  context.Background(),
			},
			args: args{
				filepath: "/home/dera/Downloads/alumuniumprofileforign.csv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TransactionUsecase{
				Repo: tt.fields.Repo,
				Ctx:  tt.fields.Ctx,
			}
			got, err := tr.LoadCSV(tt.args.filepath)
			assert.Nil(t, err)
			assert.NotNil(t, got)
		})
	}
}

func TestTransactionUsecase_BulkCompute(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type fields struct {
		Repo repository.ITransactionRepository
		Ctx  context.Context
	}
	type args struct {
		transactions []domain.Transaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []domain.Transaction
	}{
		{
			name: "Test Transaction Computation #1",
			fields: fields{
				Repo: repository.NewRepository(db),
				Ctx:  context.Background(),
			},
			args: args{
				transactions: []domain.Transaction{
					{
						MaterialId:      "CN000019",
						TransactionType: "CREDIT",
						ReferenceNo:     "ASDASDASD",
						Received:        50,
						Returns:         0,
						Issued:          0,
						Balance:         50,
						Cumulative:      "230",
						UnitPrice:       "230",
						TotalPrice:      "11500",
					},
					{
						MaterialId:      "CN000019",
						TransactionType: "CREDIT",
						ReferenceNo:     "ASDASDASD",
						Received:        100,
						Returns:         0,
						Issued:          0,
						Balance:         0,
						Cumulative:      "0",
						UnitPrice:       "235",
						TotalPrice:      "0",
					},
					{
						MaterialId:      "CN000019",
						TransactionType: "DEBIT",
						ReferenceNo:     "ASDASDASD",
						Received:        0,
						Returns:         0,
						Issued:          30,
						Balance:         0,
						Cumulative:      "0",
						UnitPrice:       "0",
						TotalPrice:      "0",
					},
					{
						MaterialId:      "CN000019",
						TransactionType: "CREDIT",
						ReferenceNo:     "ASDASDASD",
						Received:        10,
						Returns:         0,
						Issued:          0,
						Balance:         0,
						Cumulative:      "0",
						UnitPrice:       "232",
						TotalPrice:      "0",
					},
					{
						MaterialId:      "CN000019",
						TransactionType: "DEBIT",
						ReferenceNo:     "ASDASDASD",
						Received:        0,
						Returns:         0,
						Issued:          30,
						Balance:         0,
						Cumulative:      "0",
						UnitPrice:       "0",
						TotalPrice:      "0",
					},
					{
						MaterialId:      "CN000019",
						TransactionType: "CREDIT",
						ReferenceNo:     "ASDASDASD",
						Received:        20,
						Returns:         0,
						Issued:          0,
						Balance:         0,
						Cumulative:      "0",
						UnitPrice:       "200",
						TotalPrice:      "0",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TransactionUsecase{
				Repo: tt.fields.Repo,
				Ctx:  tt.fields.Ctx,
			}

			got := tr.BulkCompute(tt.args.transactions)
			assert.NotNil(t, got)
			// assert.Equal(t, got[len(got)-1].Balance, 150)
			// assert.Equal(t, got[len(got)-1].TotalPrice, "23500")
		})
	}
}

func TestTransactionUsecase_BulkImport(t *testing.T) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		panic(err)
	}
	type fields struct {
		Repo repository.ITransactionRepository
		Ctx  context.Context
	}
	type args struct {
		transactions []domain.Transaction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Testing bulk import to a data store",
			fields: fields{
				Repo: repository.NewRepository(db),
				Ctx:  context.Background(),
			},
			args: args{
				transactions: []domain.Transaction{
					{
						MaterialId:      "CN000019",
						TransactionType: "CREDIT",
						ReferenceNo:     "ASDASDASD",
						Received:        50,
						Returns:         0,
						Issued:          0,
						Balance:         50,
						Cumulative:      "230",
						UnitPrice:       "230",
						TotalPrice:      "11500",
					},
					{
						MaterialId:      "CN000019",
						TransactionType: "CREDIT",
						ReferenceNo:     "ASDASDASD",
						Received:        100,
						Returns:         0,
						Issued:          0,
						Balance:         0,
						Cumulative:      "0",
						UnitPrice:       "235",
						TotalPrice:      "0",
					},
					{
						MaterialId:      "CN000019",
						TransactionType: "DEBIT",
						ReferenceNo:     "ASDASDASD",
						Received:        0,
						Returns:         0,
						Issued:          30,
						Balance:         0,
						Cumulative:      "0",
						UnitPrice:       "0",
						TotalPrice:      "0",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TransactionUsecase{
				Repo: tt.fields.Repo,
				Ctx:  tt.fields.Ctx,
			}
			err := tr.BulkImport(tt.args.transactions)
			assert.Nil(t, err)
		})
	}
}
