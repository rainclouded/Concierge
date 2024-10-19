package handlers

import (
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPermissions(ctx *gin.Context) {
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	permissions, err := db.GetPermissions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"permissions": permissions})
}

// expects router.GET("/permissions/:id", GetPermission)
func GetPermissionById(ctx *gin.Context) {
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

	permissions, err := db.GetPermissionById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"permission": permissions})
}

func PostPermission(ctx *gin.Context) {
	var permission models.Permission
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	//https://gin-gonic.com/docs/examples/binding-and-validation/
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if permission.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Permission name is required"})
		return
	}

	if err := db.CreatePermission(&permission); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add permission"})
		return
	}

	ctx.JSON(http.StatusCreated, permission)
}

func PatchPermission(ctx *gin.Context) {
	var permReq models.Permission
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

	if err := ctx.ShouldBindJSON(&permReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	permission, err := db.GetPermissionById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("permission %d not found", id)})
	}

	if permReq.Name != "" {
		permission.Name = permReq.Name
	}

	permission.Value = permReq.Value

	db.UpdatePermission(permission)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, permission)
}
