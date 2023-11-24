package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dru-go/noah-toolbox/adapter/handler/util"
	"github.com/dru-go/noah-toolbox/domain"
	"github.com/dru-go/noah-toolbox/usecase"
)

func ImportTransactionHandler(usecase usecase.ITransactionUsecase) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		filePath, err := util.CSVUploadHandler(r, "import")
		if err != nil {
			fmt.Fprintf(w, "Error Uploading the file, %s", err.Error())
			return
		}
		transactions, err := usecase.LoadCSV(filePath)
		if err != nil {
			fmt.Fprintf(w, "Error Loading the transactions, %s", err.Error())
			return
		}
		if err = usecase.BulkImport(transactions); err != nil {
			fmt.Fprintf(w, "Error Importing the transactions, %s", err.Error())
			return
		}
		fmt.Fprintf(w, "Successfully Saved ... %v records", len(transactions))
	}
}

type RequestBody struct {
	Material     string   `json:"material,omitempty"`
	Project      string   `json:"project,omitempty"`
	Company      string   `json:"company,omitempty"`
	Transactions []string `json:"transactions,omitempty"`
}

func ComputeTransactionHandler(usecase usecase.ITransactionUsecase) HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body RequestBody

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Fprintf(w, "Error Decoding request body %s", err)
			return
		}
		// NOTE Add input validation
		transactions, err := usecase.Find(domain.ComputeFilter{MaterialId: body.Material, Project: body.Project, Company: body.Company, Ids: body.Transactions})
		if err != nil {
			fmt.Fprintf(w, "Error finding the transactions %s", err)
			return
		}
		transactions = usecase.BulkCompute(transactions)
		if r.URL.Query().Has("save") {
			fmt.Fprintf(w, "Saving to database ... %v records", len(transactions))
			usecase.BulkUpdate(transactions)
			return
		}
		util.WriteResponse[domain.Transactions](w, transactions)
	}
}
