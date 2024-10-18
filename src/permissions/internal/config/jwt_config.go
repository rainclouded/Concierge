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

func LoadSessionExp() int {
	expStr := os.Getenv("SESSION_EXPIRATION")

	if expStr != "" {
		if exp, err := strconv.Atoi(expStr); err == nil {
			return max(10, exp)
		}
	}

	return 60
}

// default publicKey := "MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAE61z8KkG7BfsioUcmMMTTbZ0hHR8kzIXIPYcpoLnqbOPHXPSM4PYCsLbhaTzuw0sASDMcdSEZqwFw3krnXwHKEp3ID5ol2vj4qpxDzZdl4T0dNvWeRMCGLZGAPVz6zOD4"
func LoadPrivateKey() (*ecdsa.PrivateKey, error) {
	pkStr := os.Getenv("JWT_PRIVATE_KEY")
	if pkStr == "" {
		pkStr = `-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDC4czoxahGqOAy2eCbsNjyEfFCsRItQ+G00whfrCbJQfsEDFN3HiSO5InXH8ZqjfmGgBwYFK4EEACKhZANiAATrXPwqQbsF+yKhRyYwxNNtnSEdHyTMhcg9hymgueps48dc9Izg9gKwtuFpPO7DSwBIMxx1IRmrAXDeSudfAcoSncgPmiXa+PiqnEPNl2XhPR029Z5EwIYtkYA9XPrM4Pg=
-----END EC PRIVATE KEY-----`
	}

	return ParseECDSAPrivateKeyFromPEM(pkStr)
}

func LoadPublicKey() (*ecdsa.PublicKey, error) {
	publicKeyString := os.Getenv("JWT_PUBLIC_KEY")
	if publicKeyString == "" {
		privateKey, err := LoadPrivateKey()
		if err != nil {
			return nil, err
		}

		return &privateKey.PublicKey, nil
	}

	return ParseECDSAPublicKeyFromPEM(publicKeyString)
}

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
