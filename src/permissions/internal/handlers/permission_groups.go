package handlers

import (
	"concierge/permissions/internal/database"
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPermissionGroups(ctx *gin.Context) {
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	permissions, err := db.GetPermissionGroups()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"groups": permissions})
}

func GetPermissionGroupById(ctx *gin.Context) {
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	id, ok := getPathParam(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	group, err := db.GetPermissionGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"group": group})
}

func PostPermissionGroups(ctx *gin.Context) {
	var group models.PermissionGroup
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err := ctx.ShouldBindJSON(&group); err != nil {
		fmt.Print(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if group.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Permission name is required"})
		return
	}

	if err := db.(database.Database).CreatePermissionGroup(&group); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add permission"})
		return
	}

	ctx.JSON(http.StatusCreated, group)
}
