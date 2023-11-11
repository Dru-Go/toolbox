package repository

import (
	"database/sql"
)

func ConnectDb(connection string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "mysql://root:admin@localhost:3306/cookbook")
	if err != nil {
		return nil, err
	}

	return db, nil
}
