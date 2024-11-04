package config

import (
	"concierge/permissions/internal/database"
	"fmt"
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
	dbImplementation := os.Getenv("PERMISSION_DB_IMPLEMENTATION")

	fmt.Printf("Loading DB")

	if dbImplementation == "MARIADB" {
		mariaDb, _ := database.NewMariaDB(loadConnectionString(), false)
		if mariaDb != nil {
			fmt.Printf("DB Connected")
			newDb = mariaDb
		} else {
			fmt.Printf("Not connected")
		}
	}

	//fallback on newMock if maria fails or db implementation is MOCK
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

func loadConnectionString() string {
	un := os.Getenv("PERMISSION_DB_USERNAME")
	pw := os.Getenv("PERMISSION_DB_PASSWORD")
	host := os.Getenv("PERMISSION_DB_HOST")
	port := os.Getenv("PERMISSION_DB_PORT")
	name := os.Getenv("PERMISSION_DB_NAME")
	cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", un, pw, host, port, name)
	return cs
}
