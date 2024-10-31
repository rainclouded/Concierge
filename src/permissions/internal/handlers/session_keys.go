package handlers

import (
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSessionKey(ctx *gin.Context) {
	var loginRequest models.LoginAttempt
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format("Invalid input", nil))
		return
	}

	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
		return
	}

	accCli, ok := middleware.GetAccountClient(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
		return
	}

	jwtCtx, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
	}

	account, err := accCli.PostLoginAttempt(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusNotFound, middleware.Format(err.Error(), nil))
		return
	}

	permissions, err := db.GetPermissionForAccountId(account.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.Format(err.Error(), nil))
		return
	}

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

	ctx.JSON(http.StatusOK, middleware.Format("Session key successfully created", gin.H{"sessionKey": sessionKey}))
}

func ParseSessionKey(ctx *gin.Context) {
	jwtCtx, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
	}

	sessionKey := jwtCtx.GetAPIKeyFromCtx(ctx)
	if sessionKey == "" {
		ctx.JSON(http.StatusUnauthorized, middleware.Format("Not logged in!", nil))
		return
	}
	data, err := jwtCtx.ParseSignedMessage(sessionKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, middleware.Format(err.Error(), nil))
		return
	}

	response := models.SessionKeyDataResponse{
		AccountID:        data.AccountID,
		AccountName:      data.AccountName,
		PermissionString: jwtCtx.GetSessionPermissions(ctx),
	}

	ctx.JSON(http.StatusOK, middleware.Format("Session key successfully read", gin.H{"sessionData": response}))
}

func GetPublicKey(ctx *gin.Context) {
	jwtCtx, ok := middleware.GetJWTContext(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error!", nil))
		return
	}

	pem, err := jwtCtx.GetPublicKeyPEM()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, middleware.Format("Internal server error", nil))
		return
	}
	ctx.JSON(http.StatusOK, middleware.Format("Public key fetched", gin.H{"publicKey": pem}))
}
