package db

import (
	"database/sql"
	"log"

	"github.com/ByChanderZap/api-basics/cmd/database"
	_ "github.com/lib/pq"
)

func NewPostgresStorage(dbUrl string) (*database.Queries, error) {
	con, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	db := database.New(con)
	return db, nil
}
