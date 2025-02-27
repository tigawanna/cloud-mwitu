package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	SQLitePath string // Path to SQLite database file
	Port       string // Server port
}

func GetEnv() Env {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error loading .env file")
		return Env{
			SQLitePath: "./data/cloud-mwitu.db",
			Port:       "8080",
		}
	}

	return Env{
		SQLitePath: getEnvWithDefault("SQLITE_PATH", "./data/cloud-mwitu.db"),
		Port:       getEnvWithDefault("PORT", "8080"),
	}
}

// getEnvWithDefault returns environment variable value or default if not set
func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
