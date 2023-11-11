package usecase

import (
	"context"

	"github.com/dru-go/noah-toolbox/adapter/repository"
	"github.com/dru-go/noah-toolbox/domain"
)

type TransactionUsecase struct {
	Repo repository.Transaction
	Ctx  context.Context
}

type Usecase interface {
	FetchLastTransaction() domain.Transaction
	BulkImport()
	Compute()
}

func (TransactionUsecase) FetchLastTransaction() domain.Transaction {
	panic("not implemented") // TODO: Implement
}

func (TransactionUsecase) BulkImport() {
	panic("not implemented") // TODO: Implement
}

func (TransactionUsecase) Compute() {
	panic("not implemented") // TODO: Implement
}
