package handlers

import (
	"concierge/permissions/internal/constants"
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPermissions retrieves all permissions from the database
// Args:
//   ctx: The Gin context containing the request and response objects
// Returns:
//   A JSON response with the status and list of permissions or an error message
func GetPermissions(ctx *gin.Context) {
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	permissions, err := db.GetPermissions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, middleware.Format("Permissions retrieved successfully", permissions))
}

// GetPermissionById retrieves a permission by its ID
// Args:
//   ctx: The Gin context containing the request and response objects
// Returns:
//   A JSON response with the status and the requested permission or an error message
func GetPermissionById(ctx *gin.Context) {
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	jwt, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	if !jwt.HasPermissionByName(ctx, constants.CanViewPermissions) {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Missing permission to view permissions", nil))
		return
	}

	id, ok := getPathParam(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid id parameter", nil))
		return
	}

	permissions, err := db.GetPermissionById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, middleware.Format(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, middleware.Format("Permission found successfully", permissions))
}

// PostPermission creates a new permission in the database
// Args:
//   ctx: The Gin context containing the request and response objects
// Returns:
//   A JSON response with the status and the created permission or an error message
func PostPermission(ctx *gin.Context) {
	var permission models.PermissionPostRequest
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	jwt, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	if !jwt.HasPermissionByName(ctx, constants.CanEditPermissions) {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Missing permission to edit permissions", nil))
		return
	}

	// Binding the request body to the permission struct
	if err := ctx.ShouldBindJSON(&permission); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid input", nil))
		return
	}

	newPerm, err := db.CreatePermission(permission.Name)
	if err != nil {
		ctx.JSON(http.StatusConflict, middleware.Format(fmt.Sprintf("Cannot create permission '%s' already exists", permission.Name), nil))
		return
	}

	ctx.JSON(http.StatusCreated, middleware.Format("Permission created successfully", newPerm))
}
