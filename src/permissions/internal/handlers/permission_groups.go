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
	var groupReq models.PermissionGroupRequest
	var group = &models.PermissionGroup{}
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err := ctx.ShouldBindJSON(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if groupReq.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Permission name is required"})
		return
	}

	if groupReq.TemplateID > 0 {
		template, err := db.GetPermissionGroupById(groupReq.TemplateID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("could not find template group %d", groupReq.TemplateID)})
			return
		}

		group = template.DeepCopy()
	}

	group.Merge(&groupReq)

	if err := db.CreatePermissionGroup(group); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add permission"})
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func PatchPermissionGroups(ctx *gin.Context) {
	var groupReq models.PermissionGroupRequest

	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if err := ctx.ShouldBindJSON(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	group, err := db.GetPermissionGroupById(groupReq.TemplateID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("permission group %d not found", groupReq.TemplateID)})
		return
	}

	group.Merge(&groupReq)

	err = db.UpdatePermissionGroup(group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func GetPermissionGroupAccount(ctx *gin.Context) {
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

	members, err := db.GetGroupMembers(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("permission group %d not found", id)})
		return
	}

	group, err := db.GetPermissionGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("permission group %d not found", id)})
	}

	ctx.JSON(http.StatusOK, gin.H{"permission-group-id": group, "accounts": members})
}

func PutPermissionGroupAccount(ctx *gin.Context) {
	var groupReq models.PermissionGroupRequest

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

	if err := ctx.ShouldBindJSON(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	group, err := db.GetPermissionGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("permission group %d not found", id)})
		return
	}

	groupReq.TemplateID = -1
	groupReq.Name = ""
	groupReq.Description = ""
	groupReq.Permissions = nil

	group.Merge(&groupReq)
}
