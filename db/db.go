package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// func NewPostgresStorage(dbUrl string) (*sql.DB, error) {
// 	con, err := sql.Open("postgres", dbUrl)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return con, nil
// }

func NewPostgresStorage(dbUrl string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	// Customize connection settings if needed
	config.MaxConns = 10
	config.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	return pool, nil
}
