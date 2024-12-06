package middleware

import (
	"concierge/permissions/internal/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetAccountClient sets the AccountClient in the Gin context
// Args:
//   accCl: The AccountClient instance to be set in the context
// Returns:
//   A Gin middleware function that sets the AccountClient in the context
func SetAccountClient(accCl client.AccountClient) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Check if the AccountClient is nil
		if accCl == nil {
			// If the client is nil, respond with an error message
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "System cannot connect to the Account server. Contact System Administrator"})
			return
		}
		// Set the AccountClient in the context
		ctx.Set("AccountClient", accCl)
		// Continue with the next middleware or handler
		ctx.Next()
	}
}

// GetAccountClient retrieves the AccountClient from the Gin context
// Args:
//   ctx: The Gin context containing the AccountClient
// Returns:
//   (client.AccountClient, bool): The AccountClient and a boolean indicating if it was found
func GetAccountClient(ctx *gin.Context) (client.AccountClient, bool) {
	// Retrieve the AccountClient from the context
	accCl, exists := ctx.Get("AccountClient")
	// Check if the AccountClient exists and is not nil
	if !exists || accCl == nil {
		return nil, false
	}

	// Type assert the AccountClient from the context
	client, ok := accCl.(client.AccountClient)
	if !ok {
		return nil, false
	}

	// Return the AccountClient and true if found
	return client, true
}
