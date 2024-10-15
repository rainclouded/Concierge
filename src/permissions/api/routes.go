package api

import (
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/database"
	"concierge/permissions/internal/handlers"
	"concierge/permissions/internal/middleware"

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

	router.POST("/access-keys", handlers.PostAccessKey)

	router.GET("/permissions", handlers.GetPermissions)
	router.GET("/permissions/:id", handlers.GetPermissionById)
	router.POST("/permissions", handlers.PostPermission)

	router.GET("/permission-groups", handlers.GetPermissionGroups)
	router.GET("/permission-groups/:id", handlers.GetPermissionById)
	router.POST("/permission-groups", handlers.PostPermissionGroups)

	return router
}
