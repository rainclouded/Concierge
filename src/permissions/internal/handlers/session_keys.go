package handlers

import (
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/middleware"
	"concierge/permissions/internal/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSessionKey(ctx *gin.Context) {
	var loginRequest models.LoginAttempt
	requestBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	err = json.Unmarshal(requestBody, &loginRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	db, ok := middleware.GetDb(ctx)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	accCli := config.LoadAccountEndpoint()
	account, err := accCli.PostLoginAttempt(loginRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	permissions, err := db.GetPermissionForAccountId(account.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sessionKey, err := middleware.SignMessage(
		&models.SessionKeyData{
			AccountID:         account.ID,
			AccountName:       account.Name,
			PermissionVersion: 1,
			PermissionString:  middleware.PermissionSliceToPermissionString(permissions),
		},
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("session-key", sessionKey, config.LoadSessionExp()*60, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"session-key": sessionKey})
}

func ParseSessionKey(ctx *gin.Context) {
	sessionKey := middleware.GetSessionKey(ctx)
	if sessionKey == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing session key"})
		return
	}

	data, err := middleware.ParseSignedMessage(sessionKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"session-data": data})
}

func GetPublicKey(ctx *gin.Context) {
	pem, err := middleware.GetPublicKeyPEM()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
	ctx.JSON(http.StatusOK, gin.H{"public-key": pem})
}
