package middleware

import (
	"concierge/permissions/internal/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetAccountClient(accCl client.AccountClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if accCl == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "System cannot connect to the Account server. Contact System Administrator"})
			return
		}
		ctx.Set("AccountClient", accCl)
		ctx.Next()
	}
}

func GetAccountClient(ctx *gin.Context) (client.AccountClient, bool) {
	accCl, exists := ctx.Get("AccountClient")
	if !exists || accCl == nil {
		return nil, false
	}

	client, ok := accCl.(client.AccountClient)
	if !ok {
		return nil, false
	}

	return client, true
}
