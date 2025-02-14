package main

import (
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
	server := api.NewAPIServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
