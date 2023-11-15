package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:admin@/cookbook")
	if err != nil {
		return nil, err
	}

	return db, nil
}
