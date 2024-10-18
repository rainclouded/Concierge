package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPathParam(ctx *gin.Context, paramKey string) (int, bool) {
	strValue := ctx.Param(paramKey)

	// Convert the ID to an integer
	value, err := strconv.Atoi(strValue)
	if err != nil {
		return -1, false
	}
	return value, true
}
