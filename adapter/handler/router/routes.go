package router

import (
	"net/http"

	"github.com/dru-go/noah-toolbox/usecase"
	"github.com/gorilla/mux"
)

func SetupRouter(MaterialUsecase usecase.IMaterialUsecase, TransactionUsecase usecase.ITransactionUsecase) {
	router := mux.NewRouter()
	m := router.PathPrefix("/material").Subrouter()
	m.HandleFunc("import", ImportHandler(MaterialUsecase))
	t := router.PathPrefix("/transaction").Subrouter()
	t.HandleFunc("import", ImportTransactionHandler(TransactionUsecase))
	t.HandleFunc("compute", ComputeTransactionHandler(TransactionUsecase))
	http.ListenAndServe(":3400", router)
}

type HandleFunc func(w http.ResponseWriter, r *http.Request)
