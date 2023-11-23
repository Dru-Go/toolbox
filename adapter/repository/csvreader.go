package repository

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/jszwec/csvutil"
)

type CSVReader struct {
	filePath string
}

func NewCSVReader(filepath string) CSVReader {
	return CSVReader{filePath: filepath}
}

// ReadCSV is a generic function that reads data from a CSV file and deserializes it using a provided serializer.
func ReadCSV[T any](c CSVReader) ([]T, error) {
	file, err := os.OpenFile(c.filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)
	// DebugWriter(file)
	dec, err := csvutil.NewDecoder(reader)
	if err != nil {
		log.Fatal(err)
	}
	var data []T
	for {
		var u T
		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(utils.PrettyPrint(u))
		data = append(data, u)
	}
	return data, nil
}
