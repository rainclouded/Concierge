package middleware

import (
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/models"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWT_Context struct {
	privateKey              *ecdsa.PrivateKey
	publicKey               *ecdsa.PublicKey
	encryptionAlgo          jwt.SigningMethod
	PermissionPerIndex      int
	sessionHeader           string
	permissionNames         map[string]int
	permissionNamesCacheExp time.Time
}

func NewJWT() *JWT_Context {
	publicKey := config.LoadPublicKey() //TODO handle public key is nil
	pk := config.LoadPrivateKey()       //TODO handle private key is nil
	signingMethod := config.LoadEncrypAlgo()
	N := config.LoadPermissionPerIndex()
	sessionKeyHeaderName := config.LoadSessionKeyHeader()

	return &JWT_Context{
		privateKey:              pk,
		publicKey:               publicKey,
		encryptionAlgo:          signingMethod,
		PermissionPerIndex:      N,
		sessionHeader:           sessionKeyHeaderName,
		permissionNames:         map[string]int{},
		permissionNamesCacheExp: time.Now(),
	}
}

func SetJWTContex(jwtCtx *JWT_Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if jwtCtx == nil {
			ctx.JSON(http.StatusInternalServerError, Format("System cannot connect to JWT module", nil))
			return
		}
		ctx.Set("jwt_ctx", jwtCtx)
		ctx.Next()
	}
}

func GetJWTContext(ctx *gin.Context) (*JWT_Context, bool) {
	jwtCtxInterface, exists := ctx.Get("jwt_ctx")
	if !exists {
		return nil, false
	}

	jwtCtx, ok := jwtCtxInterface.(*JWT_Context)
	return jwtCtx, ok
}

func (jwtCtx *JWT_Context) ParseSignedMessage(sessionKey string) (*models.SessionKeyData, error) {
	token, err := jwt.Parse(sessionKey, func(t *jwt.Token) (interface{}, error) {
		return jwtCtx.publicKey, nil
	}, jwt.WithValidMethods([]string{"ES384"}))

	if err != nil {
		return nil, fmt.Errorf("failed to parse sessionKey: %s", err.Error())
	}

	if !token.Valid {
		return nil, fmt.Errorf("sessionKey signature was not valid")
	}
	claims := token.Claims.(jwt.MapClaims)

	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return nil, fmt.Errorf("sessionKey has expired")
		}
	} else {
		return nil, fmt.Errorf("expiration claim not found")
	}

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

func (jwtCtx *JWT_Context) SignMessage(sessionData *models.SessionKeyData) (string, error) {
	claims := jwt.MapClaims{
		"accountId":         sessionData.AccountID,
		"accountName":       sessionData.AccountName,
		"permissionVersion": sessionData.PermissionVersion,
		"permissionString":  sessionData.PermissionString,
		"exp":               jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.LoadSessionExp()))),
	}

	token := jwt.NewWithClaims(jwtCtx.encryptionAlgo, claims)

	signedToken, err := token.SignedString(jwtCtx.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (jwtCtx *JWT_Context) PermissionSliceToPermissionString(permissions []*models.Permission) []int {
	slice := []int{0}
	for _, permission := range permissions {
		index := permission.ID / jwtCtx.PermissionPerIndex
		// value := int(math.Pow(2, float64(permission.ID%jwtCtx.PermissionPerIndex)))
		value := int(1 << (permission.ID % jwtCtx.PermissionPerIndex))
		for i := len(slice); i < index+1; i++ {
			slice = append(slice, 0)
		}
		slice[index] += value
	}
	return slice
}

func (jwtCtx *JWT_Context) GetPublicKeyPEM() (string, error) {
	derKey, err := x509.MarshalPKIXPublicKey(jwtCtx.publicKey)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derKey,
	}

	return string(pem.EncodeToMemory(block)), nil
}

func (jwtCtx *JWT_Context) GetAPIKeyFromCtx(ctx *gin.Context) string {
	return ctx.GetHeader(jwtCtx.sessionHeader)
}

func (jwtCtx *JWT_Context) GetPermissionStateFromPermString(permId int, permString []int) bool {
	index := permId / jwtCtx.PermissionPerIndex
	value := permString[index] & (1 << (permId % jwtCtx.PermissionPerIndex))
	return (value > 0)
}

func (jwtCtx *JWT_Context) HasPermissionByName(ctx *gin.Context, permName string) bool {
	apiKey, err := jwtCtx.ParseSignedMessage(jwtCtx.GetAPIKeyFromCtx(ctx))
	if err != nil {
		return false
	}

	if time.Now().After(jwtCtx.permissionNamesCacheExp) {
		db, ok := GetDb(ctx)
		if !ok {
			return false
		}

		permissions, err := db.GetPermissions()
		if err != nil {
			return false
		}

		jwtCtx.permissionNames = make(map[string]int)
		for _, perm := range permissions {
			jwtCtx.permissionNames[perm.Name] = perm.ID
		}
	}

	id, found := jwtCtx.permissionNames[permName]
	if !found {
		return false
	}

	return jwtCtx.GetPermissionStateFromPermString(id, apiKey.PermissionString)
}
