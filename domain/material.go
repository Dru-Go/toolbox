package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type Material struct {
	Id          string `csv:"ID,omitempty"`
	MaterialId  string `csv:"MATERIALID,omitempty"`
	Name        string `csv:"NAME,omitempty"`
	Description string `csv:"DESCRIPTION,omitempty"`
	Category    string `csv:"CATEGORY,omitempty"`
	Measurement string `csv:"MEASUREMENT,omitempty"`
	CreatedAt   string `csv:"CREATEDAT,omitempty"`
}

func RandInt(min, max int) []byte {
	unique := min + rand.Intn(max-min)
	return []byte(fmt.Sprint(unique))
}
func CountMaterialWithCategory(db *sql.DB, category string) (int, error) {
	if !Validate(category) {
		return 0, errors.New("the provided category does not exist")
	}
	// Prepare the SQL statement for counting rows
	query := "SELECT COUNT(*) FROM materials where category = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Execute the query and retrieve the result
	var count int
	err = stmt.QueryRow(category).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count, nil
}
