package middleware

import (
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/models"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ParseSignedMessage(sessionKey string) (*models.SessionKeyData, error) {
	publicKey, err := config.LoadPublicKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(sessionKey, func(t *jwt.Token) (interface{}, error) {
		return publicKey, nil
	}, jwt.WithValidMethods([]string{"ES384"}))

	if err != nil {
		fmt.Printf("Invalid key: %s", sessionKey)
		return nil, fmt.Errorf("failed to parse sessionKey: %s", err.Error())
	}

	if !token.Valid {
		return nil, fmt.Errorf("sessionKey signature was not valid")
	}
	claims := token.Claims.(jwt.MapClaims)
	var sessionData models.SessionKeyData
	data, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sessionKey: %s", err.Error())
	}

	err = json.Unmarshal(data, &sessionData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sessionKey: %s", err.Error())
	}

	return &sessionData, nil
}

func SignMessage(sessionData *models.SessionKeyData) (string, error) {
	signingMethod := config.LoadEncrypAlgo()
	claims := jwt.MapClaims{
		"accountId":         sessionData.AccountID,
		"accountName":       sessionData.AccountName,
		"permissionVersion": sessionData.PermissionVersion,
		"permissionString":  sessionData.PermissionString,
		"exp":               jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.LoadSessionExp()))),
	}

	token := jwt.NewWithClaims(signingMethod, claims)

	pk, err := config.LoadPrivateKey()
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(pk)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func PermissionSliceToPermissionString(permissions []*models.Permission) []int {
	N := config.LoadPermissionPerIndex() //Controls # of permissions per list element
	slice := []int{0}
	for _, permission := range permissions {
		index := permission.ID / N
		value := int(math.Pow(2, float64(permission.ID%N)))
		for i := len(slice); i < index+1; i++ {
			slice = append(slice, 0)
		}
		slice[index] += value
	}
	return slice
}

func GetPublicKeyPEM() (string, error) {
	publicKey, err := config.LoadPublicKey()
	if err != nil {
		return "", err
	}

	derKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derKey,
	}

	return string(pem.EncodeToMemory(block)), nil
}

func GetAPIKeyFromCtx(ctx *gin.Context) string {
	value := ctx.Request.Header[config.LoadSessionKeyHeader()]
	if len(value) == 0 {
		return ""
	}
	return value[0]
}
