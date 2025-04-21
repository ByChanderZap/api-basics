package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl                string
	Port                 string
	JWTExpirationInHours int64
	JWTSecret            string
}

var Envs = initConfig()

func initConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default or environment variables")
	}
	return Config{
		DBUrl:                getEnv("DB_URL", "FUCK"),
		Port:                 getEnv("PORT", ":8080"),
		JWTExpirationInHours: getEnvAsInt("JWT_EXPIRATION_IN_HOURS", 168),
		JWTSecret:            getEnv("JWT_SECRET", "not_that_secret_secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
