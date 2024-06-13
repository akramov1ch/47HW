package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "password=password user=postgres dbname=users sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
