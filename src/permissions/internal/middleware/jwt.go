package middleware

import (
	"concierge/permissions/internal/config"
	"concierge/permissions/internal/models"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWT_Context struct {
	privateKey         *ecdsa.PrivateKey
	publicKey          *ecdsa.PublicKey
	encryptionAlgo     jwt.SigningMethod
	PermissionPerIndex int
	sessionHeader      string
}

func NewJWT() *JWT_Context {
	publicKey := config.LoadPublicKey() //TODO handle public key is nil
	pk := config.LoadPrivateKey()       //TODO handle private key is nil
	signingMethod := config.LoadEncrypAlgo()
	N := config.LoadPermissionPerIndex()
	sessionKeyHeaderName := config.LoadSessionKeyHeader()

	return &JWT_Context{
		privateKey:         pk,
		publicKey:          publicKey,
		encryptionAlgo:     signingMethod,
		PermissionPerIndex: N,
		sessionHeader:      sessionKeyHeaderName,
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
		value := int(math.Pow(2, float64(permission.ID%jwtCtx.PermissionPerIndex)))
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
