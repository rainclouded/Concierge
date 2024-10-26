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
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	permissions, err := db.GetPermissions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, middleware.Format("Permissions retreived successfully", permissions))
}

// expects router.GET("/permissions/:id", GetPermission)
func GetPermissionById(ctx *gin.Context) {
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
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

func PostPermission(ctx *gin.Context) {
	var permission models.PermissionPostRequest
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	//https://gin-gonic.com/docs/examples/binding-and-validation/
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
