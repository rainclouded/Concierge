package config

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

// LoadEncrypAlgo loads the encryption algorithm for JWT signing from the environment variable "JWT_SIGNING_METHOD".
// It returns the appropriate jwt.SigningMethod based on the value of the method variable.
// Args:
//     None
// Returns:
//     jwt.SigningMethod: The signing method used for JWT signing, like HS256, RS256, or ECDSA384.
func LoadEncrypAlgo() jwt.SigningMethod {
	method := os.Getenv("JWT_SIGNING_METHOD")

	switch method {
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	case "RS256":
		return jwt.SigningMethodRS256
	case "ECDSA384":
		return jwt.SigningMethodES384
	default:
		return jwt.SigningMethodES384
	}
}

// LoadSessionExp loads the session expiration time from the environment variable "SESSION_EXPIRATION".
// If the value is empty or invalid, it returns the default value of 60 minutes. The minimum expiration time is 10 minutes.
// Args:
//     None
// Returns:
//     int: The session expiration time in minutes.
func LoadSessionExp() int {
	expStr := os.Getenv("SESSION_EXPIRATION")

	if expStr != "" {
		if exp, err := strconv.Atoi(expStr); err == nil {
			return max(10, exp) // Ensure expiration time is at least 10 minutes
		}
	}

	return 60 // Default expiration time is 60 minutes
}

// LoadPrivateKey loads the private key for JWT signing from the environment variable "JWT_PRIVATE_KEY".
// If the key is not found in the environment variable, a default private key is used.
// Args:
//     None
// Returns:
//     *ecdsa.PrivateKey: The parsed ECDSA private key.
func LoadPrivateKey() *ecdsa.PrivateKey {
	pkStr := os.Getenv("JWT_PRIVATE_KEY")
	if pkStr == "" {
		pkStr = `-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDC4czoxahGqOAy2eCbsNjyEfFCsRItQ+G00whfrCbJQfsEDFN3HiSO5InXH8ZqjfmGgBwYFK4EEACKhZANiAATrXPwqQbsF+yKhRyYwxNNtnSEdHyTMhcg9hymgueps48dc9Izg9gKwtuFpPO7DSwBIMxx1IRmrAXDeSudfAcoSncgPmiXa+PiqnEPNl2XhPR029Z5EwIYtkYA9XPrM4Pg=
-----END EC PRIVATE KEY-----`
	}
	pk, err := ParseECDSAPrivateKeyFromPEM(pkStr)
	if err != nil {
		return nil
	}
	return pk
}

// LoadPublicKey loads the public key for JWT verification from the environment variable "JWT_PUBLIC_KEY".
// If the key is not found, it generates the public key from the loaded private key.
// Args:
//     None
// Returns:
//     *ecdsa.PublicKey: The parsed ECDSA public key.
func LoadPublicKey() *ecdsa.PublicKey {
	publicKeyString := os.Getenv("JWT_PUBLIC_KEY")
	if publicKeyString == "" {
		privateKey := LoadPrivateKey()

		return &privateKey.PublicKey
	}
	key, err := ParseECDSAPublicKeyFromPEM(publicKeyString)
	if err != nil {
		return nil
	}
	return key
}

// ParseECDSAPrivateKeyFromPEM parses an ECDSA private key from the provided PEM-encoded string.
// Args:
//     pemStr (string): The PEM-encoded private key string.
// Returns:
//     *ecdsa.PrivateKey: The parsed ECDSA private key.
//     error: An error if the parsing fails, otherwise nil.
func ParseECDSAPrivateKeyFromPEM(pemStr string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil || block.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode ECDSA private key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// ParseECDSAPublicKeyFromPEM parses an ECDSA public key from the provided PEM-encoded string.
// Args:
//     pemStr (string): The PEM-encoded public key string.
// Returns:
//     *ecdsa.PublicKey: The parsed ECDSA public key.
//     error: An error if the parsing fails, otherwise nil.
func ParseECDSAPublicKeyFromPEM(pemStr string) (*ecdsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode ECDSA public key")
	}

	publicKeyRSA, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return publicKeyRSA.(*ecdsa.PublicKey), nil
}
