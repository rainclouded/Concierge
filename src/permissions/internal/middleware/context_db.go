package middleware

import (
	"concierge/permissions/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetDb(db database.Database) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if db == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "System cannot connect to the database. Contact System Administrator"})
			return
		}
		ctx.Set("db", db)
		ctx.Next()
	}
}

func GetDb(ctx *gin.Context) (database.Database, bool) {
	dbInterface, exists := ctx.Get("db")
	if !exists || dbInterface == nil {
		return nil, false
	}

	db, ok := dbInterface.(database.Database)
	if !ok {
		return nil, false
	}

	return db, true
}
