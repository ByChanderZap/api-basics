package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewPostgresStorage(dbUrl string) (*sql.DB, error) {
	con, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	return con, nil
}
