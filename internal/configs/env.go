package configs

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	SQLitePath  string // Path to SQLite database file
	Port        string // Server port
	DatabaseDSN string // Optional - only needed if you want to override the constructed DSN
	Debug       bool   // Debug mode (true/false)
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
		SQLitePath:  getEnvWithDefault("SQLITE_PATH", "./data/cloud-mwitu.db"),
		Port:        getEnvWithDefault("PORT", "8080"),
		DatabaseDSN: fmt.Sprintf("file:%s?cache=shared", getEnvWithDefault("SQLITE_PATH", "./data/cloud-mwitu.db")),
		Debug:       getBoolEnvWithDefault("DEBUG", false),
	}
}

func getBoolEnvWithDefault(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return boolValue
}

// getEnvWithDefault returns environment variable value or default if not set
func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
