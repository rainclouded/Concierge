package api

import (
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/database"
	"concierge/permissions/internal/handlers"
	"concierge/permissions/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultNewRouter() *gin.Engine {
	db := config.LoadDB()
	router := NewRouter(db)
	return router
}

func NewRouter(db database.Database) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.SetDb(db))

	router.POST("/sessions", handlers.PostSessionKey)
	router.GET("/sessions/", handlers.ParseSessionKey)
	router.GET("/sessions/public-key", handlers.GetPublicKey)

	router.GET("/permissions/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	router.GET("/permissions", handlers.GetPermissions)
	router.GET("/permissions/:id", handlers.GetPermissionById)
	router.POST("/permissions", handlers.PostPermission)
	router.PATCH("/permissions/:id", handlers.PatchPermission)

	router.GET("/permission-groups", handlers.GetPermissionGroups)
	router.GET("/permission-groups/:id", handlers.GetPermissionById)
	router.POST("/permission-groups", handlers.PostPermissionGroups)
	router.PATCH("/permission-groups", handlers.PatchPermissionGroups)

	router.GET("/permission-groups/:id/accounts", handlers.GetPermissionGroupAccount)
	router.PUT("/permission-groups/:id/accounts", handlers.PutPermissionGroupAccount)

	return router
}
