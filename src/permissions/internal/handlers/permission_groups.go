package handlers

import (
	"concierge/permissions/internal/constants"
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPermissionGroups retrieves all permission groups from the database.
// It checks if the user has the necessary permissions to view the groups.
// Args:
//    ctx: The Gin context object that holds request and response data.
// Returns:
//    None. It writes the response directly to the client based on the outcome.
func GetPermissionGroups(ctx *gin.Context) {
	// Fetch the database object from the context
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Get the JWT context for permission checks
	jwt, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Check if the user has the permission to view permission groups
	if !jwt.HasPermissionByName(ctx, constants.CanViewPermissionGroups) {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Missing permission to view permission groups", nil))
		return
	}


	accountIdFilter := ctx.DefaultQuery("account-id", "-1")
	accountNumber, err := strconv.Atoi(accountIdFilter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	var permissions []*models.PermissionGroup

	if accountNumber >= 0 {
		permissions, err = db.GetPermissionGroupsByAccount(accountNumber)
	} else {
		permissions, err = db.GetPermissionGroups()
	}


	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	// Return the permission groups in the response
	ctx.JSON(http.StatusOK, middleware.Format("Permission groups retrieved successfully", permissions))
}

// GetPermissionGroupById retrieves a permission group by its ID from the database.
// It checks if the user has the necessary permissions to view the group.
// Args:
//    ctx: The Gin context object containing request and response data.
// Returns:
//    None. The response is written directly to the client.
func GetPermissionGroupById(ctx *gin.Context) {
	// Fetch the database object from the context
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Get the JWT context for permission checks
	jwt, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Check if the user has permission to view permission groups
	if !jwt.HasPermissionByName(ctx, constants.CanViewPermissionGroups) {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Missing permission to view permission groups", nil))
		return
	}

	// Get the ID parameter from the URL
	id, ok := getPathParam(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid id parameter", nil))
		return
	}

	// Retrieve the permission group by ID from the database
	group, err := db.GetPermissionGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, middleware.Format(err.Error(), nil))
		return
	}

	// Return the permission group details in the response
	ctx.JSON(http.StatusOK, middleware.Format("Permission group retrieved successfully", group))
}

// PostPermissionGroups creates a new permission group in the database.
// It checks if the user has the necessary permissions and validates the input data.
// Args:
//    ctx: The Gin context object containing request and response data.
// Returns:
//    None. The response is written directly to the client.
func PostPermissionGroups(ctx *gin.Context) {
	var groupReq models.PermissionGroupRequest

	// Fetch the database object from the context
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Get the JWT context for permission checks
	jwt, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Check if the user has permission to edit permission groups
	if !jwt.HasPermissionByName(ctx, constants.CanEditPermissionGroups) {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Missing permission to edit permission groups", nil))
		return
	}

	// Bind the JSON input to the PermissionGroupRequest struct
	if err := ctx.ShouldBindJSON(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid input", nil))
		return
	}

	// Validate that the group name is not empty
	if groupReq.Name == "" {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Permission group name is required", nil))
		return
	}

	// Prevent removing members when creating a new group
	if len(groupReq.MembersRemove) > 0 {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Cannot remove members when creating a new group", nil))
		return
	}

	// Create the new permission group in the database
	if err := db.CreatePermissionGroup(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	// Return a success response
	ctx.JSON(http.StatusCreated, middleware.Format("Group created successfully", nil))
}

// PatchPermissionGroups updates an existing permission group in the database.
// It checks if the user has the necessary permissions and validates the input data.
// Args:
//    ctx: The Gin context object containing request and response data.
// Returns:
//    None. The response is written directly to the client.
func PatchPermissionGroups(ctx *gin.Context) {
	var groupReq models.PermissionGroupRequest

	// Fetch the database object from the context
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Get the JWT context for permission checks
	jwt, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Check if the user has permission to edit permission groups
	if !jwt.HasPermissionByName(ctx, constants.CanEditPermissionGroups) {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Missing permission to edit permission groups", nil))
		return
	}

	// Bind the JSON input to the PermissionGroupRequest struct
	if err := ctx.ShouldBindJSON(&groupReq); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid input", nil))
		return
	}

	// Get the ID parameter from the URL
	id, ok := getPathParam(ctx, "id")
	if !ok {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid id parameter", nil))
		return
	}

	// Check if the permission group exists
	_, err := db.GetPermissionGroupById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, middleware.Format(fmt.Sprintf("Permission group %d not found", id), nil))
		return
	}

	// Update the permission group in the database
	err = db.UpdatePermissionGroup(id, &groupReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.Format(err.Error(), nil))
		return
	}

	// Return a success response
	ctx.JSON(http.StatusOK, middleware.Format("Group updated successfully", nil))
}
