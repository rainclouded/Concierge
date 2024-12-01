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

// JWT_Context holds the context for managing JWT authentication and permissions
// It stores keys, encryption methods, and permission handling data
type JWT_Context struct {
	privateKey              *ecdsa.PrivateKey   // The private key used for signing JWT tokens
	publicKey               *ecdsa.PublicKey    // The public key used for verifying JWT tokens
	encryptionAlgo          jwt.SigningMethod   // The encryption algorithm used for signing JWT
	PermissionPerIndex      int                  // The number of permissions per index in the permission string
	sessionHeader           string               // The header name used for session keys
	permissionNames         map[string]int      // A map of permission names to their IDs
	permissionNamesCacheExp time.Time            // The expiration time for the permission names cache
}

// NewJWT creates a new instance of JWT_Context
// It loads the public key, private key, encryption algorithm, and other config settings
// Returns:
//   *JWT_Context: A new JWT_Context instance
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

// SetJWTContex sets the JWT context in the Gin middleware
// Args:
//   jwtCtx: The JWT_Context instance to be set in the request context
// Returns:
//   gin.HandlerFunc: A Gin middleware function to set JWT_Context
func SetJWTContex(jwtCtx *JWT_Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if jwtCtx == nil {
			ctx.JSON(http.StatusInternalServerError, Format("System cannot connect to JWT module", nil))
			return
		}
		// Set the JWT context in the Gin request context
		ctx.Set("jwt_ctx", jwtCtx)
		// Continue to the next middleware or handler
		ctx.Next()
	}
}

// GetJWTContext retrieves the JWT context from the Gin request context
// Args:
//   ctx: The Gin context containing the JWT context
// Returns:
//   (*JWT_Context, bool): The JWT_Context instance and a boolean indicating if it was found
func GetJWTContext(ctx *gin.Context) (*JWT_Context, bool) {
	jwtCtxInterface, exists := ctx.Get("jwt_ctx")
	if !exists {
		return nil, false
	}

	jwtCtx, ok := jwtCtxInterface.(*JWT_Context)
	return jwtCtx, ok
}

// ParseSignedMessage parses and validates the session key from the request
// Args:
//   sessionKey: The session key to be parsed and validated
// Returns:
//   (*models.SessionKeyData, error): The parsed session key data or an error if parsing fails
func (jwtCtx *JWT_Context) ParseSignedMessage(sessionKey string) (*models.SessionKeyData, error) {
	println(`sessionKey: ` + sessionKey)
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

	// Check if the sessionKey has expired
	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			return nil, fmt.Errorf("sessionKey has expired")
		}
	} else {
		return nil, fmt.Errorf("expiration claim not found")
	}

	// Deserialize claims into SessionKeyData model
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

// SignMessage signs the session data into a JWT token
// Args:
//   sessionData: The session data to be signed into a JWT token
// Returns:
//   (string, error): The signed JWT token or an error if signing fails
func (jwtCtx *JWT_Context) SignMessage(sessionData *models.SessionKeyData) (string, error) {
	claims := jwt.MapClaims{
		"accountId":         sessionData.AccountID,
		"accountName":       sessionData.AccountName,
		"permissionVersion": sessionData.PermissionVersion,
		"permissionString":  sessionData.PermissionString,
		"exp":               jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(config.LoadSessionExp()))),
	}

	token := jwt.NewWithClaims(jwtCtx.encryptionAlgo, claims)

	// Sign the JWT token with the private key
	signedToken, err := token.SignedString(jwtCtx.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// PermissionSliceToPermissionString converts a slice of permissions to a permission string representation
// Args:
//   permissions: A slice of Permission models to be converted
// Returns:
//   []int: The permission string representation as an array of integers
func (jwtCtx *JWT_Context) PermissionSliceToPermissionString(permissions []*models.Permission) []int {
	slice := []int{0}
	for _, permission := range permissions {
		index := permission.ID / jwtCtx.PermissionPerIndex
		value := int(math.Pow(2, float64(permission.ID%jwtCtx.PermissionPerIndex)))
		print(fmt.Sprintf("%s: Total: %d, adding: %d\n", permission.Name, value, slice[index]))

		// Ensure the slice is large enough for the index
		for i := len(slice); i < index+1; i++ {
			slice = append(slice, 0)
		}
		slice[index] += value
	}
	fmt.Printf("Total: %d\n", slice[0])
	return slice
}

// GetPublicKeyPEM retrieves the public key in PEM format
// Returns:
//   (string, error): The public key in PEM format or an error if retrieval fails
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

// GetAPIKeyFromCtx retrieves the session key from the request header
// Args:
//   ctx: The Gin context containing the session key in the header
// Returns:
//   string: The session key retrieved from the request header
func (jwtCtx *JWT_Context) GetAPIKeyFromCtx(ctx *gin.Context) string {
	return ctx.GetHeader(jwtCtx.sessionHeader)
}

// GetPermissionStateFromPermString checks if the permission ID is set in the permission string
// Args:
//   permId: The permission ID to be checked
//   permString: The permission string to check against
// Returns:
//   bool: True if the permission is present, false otherwise
func (jwtCtx *JWT_Context) GetPermissionStateFromPermString(permId int, permString []int) bool {
	index := permId / jwtCtx.PermissionPerIndex
	value := permString[index] & (1 << (permId % jwtCtx.PermissionPerIndex))
	return (value > 0)
}

// HasPermissionByName checks if the user has the specified permission by name
// Args:
//   ctx: The Gin context containing the session key
//   permName: The name of the permission to check
// Returns:
//   bool: True if the user has the permission, false otherwise
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

		// Update the permission names cache
		jwtCtx.permissionNames = make(map[string]int)
		for _, perm := range permissions {
			jwtCtx.permissionNames[perm.Name] = perm.ID
		}
	}

	// Get the permission ID for the specified permission name
	id, found := jwtCtx.permissionNames[permName]
	if !found {
		return false
	}

	// Check if the user has the specified permission
	return jwtCtx.GetPermissionStateFromPermString(id, apiKey.PermissionString)
}

// GetSessionPermissions retrieves all permissions for the user's session
// Args:
//   ctx: The Gin context containing the session key
//   sessionData: The session data for the current user
// Returns:
//   []string: A list of permissions assigned to the user
func (jwtCtx *JWT_Context) GetSessionPermissions(ctx *gin.Context, sessionData *models.SessionKeyData) []string {
	apiKey, err := jwtCtx.ParseSignedMessage(jwtCtx.GetAPIKeyFromCtx(ctx))
	if err != nil {
		return []string{}
	}

	permissions := []string{}
	if time.Now().After(jwtCtx.permissionNamesCacheExp) {
		db, ok := GetDb(ctx)
		if !ok {
			return permissions
		}

		permissionsVal, err := db.GetPermissionForAccountId(sessionData.AccountID)
		if err != nil {
			return permissions
		}

		// Update the permission names cache
		jwtCtx.permissionNames = make(map[string]int)
		for _, perm := range permissionsVal {
			jwtCtx.permissionNames[perm.Name] = perm.ID
		}
	}

	// Add permissions to the result list
	for key, value := range jwtCtx.permissionNames {
		if jwtCtx.GetPermissionStateFromPermString(value, apiKey.PermissionString) {
			permissions = append(permissions, key)
		}
	}

	return permissions
}
