package router

import (
	"fmt"
	"net/http"

	"github.com/dru-go/noah-toolbox/adapter/handler/util"
	"github.com/dru-go/noah-toolbox/usecase"
)

func ImportHandler(usecase usecase.IMaterialUsecase) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		filePath, err := util.CSVUploadHandler(r, "import")
		if err != nil {
			fmt.Fprintf(w, "Error Uploading the file, %s", err.Error())
			return
		}
		usecase.BulkImport(filePath, r.FormValue("category"))
	}
}
