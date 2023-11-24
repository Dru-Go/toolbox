package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResponse[T any](w http.ResponseWriter, response T) {
	data, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, "Error encoding the data result %s", err)
		return
	}
	w.Write(data)
}
