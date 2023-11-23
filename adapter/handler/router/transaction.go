package router

import (
	"net/http"

	"github.com/dru-go/noah-toolbox/usecase"
)

func ImportTransactionHandler(usecase usecase.ITransactionUsecase) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Import from file"))
	}
}
func ComputeTransactionHandler(usecase usecase.ITransactionUsecase) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Import from file"))
	}
}
