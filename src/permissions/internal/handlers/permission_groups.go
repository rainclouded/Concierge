package handlers

import (
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPermissionGroups(ctx *gin.Context) {
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	permissions, err := db.GetPermissionGroups()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, middleware.Format("Permission groups retreived successfully", permissions))
}

func GetPermissionGroupById(ctx *gin.Context) {
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

	group, err := db.GetPermissionGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, middleware.Format(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, middleware.Format("Permission group retreived successfully", group))
}

func PostPermissionGroups(ctx *gin.Context) {
	var groupReq models.PermissionGroupRequest
	var group = &models.PermissionGroup{}
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	if err := ctx.ShouldBindJSON(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid input", nil))
		return
	}

	if groupReq.Name == "" {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Permission name is required", nil))
		return
	}

	if len(groupReq.MembersRemove) > 0 {
		ctx.JSON(http.StatusBadRequest, middleware.Format("cannot remove members when creating a new group", nil))
		return
	}

	if err := db.CreatePermissionGroup(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func PatchPermissionGroups(ctx *gin.Context) {
	var groupReq models.PermissionGroupRequest

	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	if err := ctx.ShouldBindJSON(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid input", nil))
		return
	}

	id, ok := getPathParam(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid id parameter", nil))
		return
	}

	group, err := db.GetPermissionGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, middleware.Format(fmt.Sprintf("permission group %d not found", id), nil))
		return
	}

	err = db.UpdatePermissionGroup(id, &groupReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.Format(err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, group)
}
