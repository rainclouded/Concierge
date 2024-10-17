package middleware

import (
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/models"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignMessage(accountID int, accountName string, permissionVersion int, permissionString []int) (string, error) {
	signingMethod := config.LoadEncrypAlgo()
	claims := jwt.MapClaims{
		"account-id":         accountID,
		"account-name":       accountName,
		"permission-version": permissionVersion,
		"permission-string":  permissionString,
		"exp":                jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.LoadSessionExp()))),
	}

	token := jwt.NewWithClaims(signingMethod, claims)

	pk, err := ParseECDSAPrivateKeyFromPEM(config.LoadPrivateKey())
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(pk)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseECDSAPrivateKeyFromPEM(pemStr string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing the private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse EC private key: %v", err)
	}

	return privateKey, nil
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
