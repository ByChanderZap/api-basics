package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ByChanderZap/api-basics/cmd/api"
	"github.com/ByChanderZap/api-basics/config"
	"github.com/ByChanderZap/api-basics/db"
)

func main() {
	fmt.Println(config.Envs.DBUrl)
	db, err := db.NewPostgresStorage(config.Envs.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Connected")
}
