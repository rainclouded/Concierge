package handlers

import (
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostSessionKey creates a session key for a user after successful login
// Args:
//   ctx: The Gin context containing the request and response objects
// Returns:
//   A JSON response containing the session key or an error message
func PostSessionKey(ctx *gin.Context) {
	var loginRequest models.LoginAttempt
	// Bind the incoming JSON request to the loginRequest struct
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid input", nil))
		return
	}

	// Retrieve the database client
	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
		return
	}

	// Retrieve the account client
	accCli, ok := middleware.GetAccountClient(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
		return
	}

	// Retrieve the JWT context
	jwtCtx, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
	}

	// Attempt login with the provided credentials
	account, err := accCli.PostLoginAttempt(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusNotFound, middleware.Format(err.Error(), nil))
		return
	}

	// Retrieve permissions for the account
	permissions, err := db.GetPermissionForAccountId(account.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.Format(err.Error(), nil))
		return
	}

	// Generate the session key for the logged-in account
	sessionKey, err := jwtCtx.SignMessage(
		&models.SessionKeyData{
			AccountID:         account.ID,
			AccountName:       account.Name,
			PermissionVersion: 1,
			PermissionString:  jwtCtx.PermissionSliceToPermissionString(permissions),
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.Format(err.Error(), nil))
		return
	}

	// Return the session key in the response
	ctx.JSON(http.StatusOK, middleware.Format("Session key successfully created", gin.H{"sessionKey": sessionKey}))
}

// ParseSessionKey parses and returns the session key information
// Args:
//   ctx: The Gin context containing the request and response objects
// Returns:
//   A JSON response containing the session data or an error message
func ParseSessionKey(ctx *gin.Context) {
	// Retrieve the JWT context
	jwtCtx, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
	}

	// Get the session key from the request context
	sessionKey := jwtCtx.GetAPIKeyFromCtx(ctx)
	if sessionKey == "" {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Not logged in!", nil))
		return
	}

	// Parse the session key
	data, err := jwtCtx.ParseSignedMessage(sessionKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	// Prepare the session data response
	response := models.SessionKeyDataResponse{
		AccountID:        data.AccountID,
		AccountName:      data.AccountName,
		PermissionString: jwtCtx.GetSessionPermissions(ctx, data),
	}

	// Return the session data in the response
	ctx.JSON(http.StatusOK, middleware.Format("Session key successfully read", gin.H{"sessionData": response}))
}

// GetPublicKey fetches and returns the public key for signing/verifying JWTs
// Args:
//   ctx: The Gin context containing the request and response objects
// Returns:
//   A JSON response containing the public key or an error message
func GetPublicKey(ctx *gin.Context) {
	// Retrieve the JWT context
	jwtCtx, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
		return
	}

	// Retrieve the public key in PEM format
	pem, err := jwtCtx.GetPublicKeyPEM()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}

	// Return the public key in the response
	ctx.JSON(http.StatusOK, middleware.Format("Public key fetched", gin.H{"publicKey": pem}))
}
