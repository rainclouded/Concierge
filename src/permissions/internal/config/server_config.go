package config

import (
	"concierge/permissions/internal/database"
	"os"
	"strconv"
)

type ServerConfig struct {
	ServerPort string
}

func LoadConfig() (ServerConfig, error) {
	config := ServerConfig{
		ServerPort: "8080", // Set your desired port here
	}
	return config, nil
}

func LoadDB() (db database.Database) {
	var newDb database.Database
	dbImplementation := os.Getenv("DB_IMPLEMENTATION")

	if dbImplementation == "MARIADB" {
		newDb = database.NewMariaDB()
	}

	//fallback on newMock if maria fails
	if newDb == nil {
		newDb = database.NewMockDB()
	}

	return newDb
}

func LoadPermissionPerIndex() int {
	valueStr := os.Getenv("PERMISSIONS_PER_INDEX")
	if valueStr != "" {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return max(1, value)
		}
	}

	return 30
}

func LoadSessionKeyHeader() string {
	valueStr := os.Getenv("SESSION_KEY_HEADER")
	if valueStr == "" {
		valueStr = "X-API-Key"
	}

	return valueStr
}
