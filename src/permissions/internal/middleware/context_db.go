package middleware

import (
	"concierge/permissions/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetDb sets the Database connection in the Gin context
// Args:
//   db: The Database instance to be set in the context
// Returns:
//   A Gin middleware function that sets the Database instance in the context
func SetDb(db database.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if the database connection is nil
		if db == nil {
			// If the database connection is nil, respond with an error message
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "System cannot connect to the database. Contact System Administrator"})
			return
		}
		// Set the Database instance in the context
		ctx.Set("db", db)
		// Continue with the next middleware or handler
		ctx.Next()
	}
}

// GetDb retrieves the Database connection from the Gin context
// Args:
//   ctx: The Gin context containing the Database instance
// Returns:
//   (database.Database, bool): The Database instance and a boolean indicating if it was found
func GetDb(ctx *gin.Context) (database.Database, bool) {
	// Retrieve the Database instance from the context
	dbInterface, exists := ctx.Get("db")
	// Check if the Database instance exists and is not nil
	if !exists || dbInterface == nil {
		return nil, false
	}

	// Type assert the Database instance from the context
	db, ok := dbInterface.(database.Database)
	if !ok {
		return nil, false
	}

	// Return the Database instance and true if found
	return db, true
}
