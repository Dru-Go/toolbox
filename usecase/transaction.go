package usecase

import (
	"context"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/dru-go/noah-toolbox/domain"
)

type TransactionUsecase struct {
	Repo repository.ITransactionRepository
	Ctx  context.Context
}

type ITransactionUsecase interface {
	LoadCSV(filepath string) ([]domain.Transaction, error)
	BulkImport([]domain.Transaction) error
	BulkCompute(domain.ComputeFilter) []domain.Transaction
}
