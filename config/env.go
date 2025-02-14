package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		DBUrl: getEnv("DB_URL", "FUCK"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
