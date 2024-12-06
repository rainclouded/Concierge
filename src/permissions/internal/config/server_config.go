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

// LoadConfig loads the server configuration, including the server port.
// Args:
//     None
// Returns:
//     ServerConfig: The configuration for the server, including the port.
//     error: Any error encountered during loading (not used here).
func LoadConfig() (ServerConfig, error) {
	config := ServerConfig{
		ServerPort: "8080", // Set your desired port here
	}
	return config, nil
}

// LoadDB loads the database configuration based on the environment variable "PERMISSION_DB_IMPLEMENTATION".
// It attempts to connect to MariaDB if the implementation is specified as "MARIADB", and falls back to a mock database if it fails or is not specified.
// Args:
//     None
// Returns:
//     database.Database: The database instance (MariaDB or mock database).
func LoadDB() (db database.Database) {
	var newDb database.Database
	dbImplementation := os.Getenv("PERMISSION_DB_IMPLEMENTATION")

	fmt.Println("Loading DB")

	if dbImplementation == "MARIADB" {
		mariaDb, _ := database.NewMariaDB(loadConnectionString(), false)
		if mariaDb != nil {
			fmt.Println("DB Connected")
			newDb = mariaDb
		} else {
			fmt.Println("Not connected")
		}
	}

	// Fallback to newMockDB if MariaDB fails or if the DB implementation is MOCK.
	if newDb == nil {
		newDb = database.NewMockDB()
	}

	return newDb
}

// LoadPermissionPerIndex loads the number of permissions to be displayed per index from the environment variable "PERMISSIONS_PER_INDEX".
// If the value is empty or invalid, it returns a default of 30.
// Args:
//     None
// Returns:
//     int: The number of permissions to display per index.
func LoadPermissionPerIndex() int {
	valueStr := os.Getenv("PERMISSIONS_PER_INDEX")
	if valueStr != "" {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return max(1, value)
		}
	}

	return 30 // Default to 30 if no valid value is found
}

// LoadSessionKeyHeader loads the header name for the session key from the environment variable "SESSION_KEY_HEADER".
// If not set, it defaults to "X-API-Key".
// Args:
//     None
// Returns:
//     string: The session key header name.
func LoadSessionKeyHeader() string {
	valueStr := os.Getenv("SESSION_KEY_HEADER")
	if valueStr == "" {
		valueStr = "X-API-Key"
	}

	return valueStr
}

// loadConnectionString generates the connection string for the database using environment variables for username, password, host, port, and database name.
// Args:
//     None
// Returns:
//     string: The generated connection string for the database.
func loadConnectionString() string {
	un := os.Getenv("PERMISSION_DB_USERNAME")
	pw := os.Getenv("PERMISSION_DB_PASSWORD")
	host := os.Getenv("PERMISSION_DB_HOST")
	port := os.Getenv("PERMISSION_DB_PORT")
	name := os.Getenv("PERMISSION_DB_NAME")
	cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", un, pw, host, port, name)
	return cs
}
