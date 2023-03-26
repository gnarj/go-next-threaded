package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateDBConnection() (*sql.DB, error) {
	connStr := "postgresql://postgres@localhost:5432/todos?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}
