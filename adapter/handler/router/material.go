package router

import (
	"net/http"

	"github.com/dru-go/noah-toolbox/usecase"
)

func ImportHandler(usecase usecase.IMaterialUsecase) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Import from file"))
	}
}
