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

func PostAccessKey(ctx *gin.Context) {
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

	accessKey, err := middleware.SignMessage(account.ID, account.Name, 1, middleware.PermissionSliceToPermissionString(permissions))

	ctx.SetCookie("session-key", accessKey, config.LoadSessionExp()*60, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"session-key": accessKey})
}
