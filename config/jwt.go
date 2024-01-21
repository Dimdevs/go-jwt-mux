package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY *ecdsa.PrivateKey

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}

func init() {
	var err error
	JWT_KEY, err = generateECDSAKey()
	if err != nil {
		log.Fatal("Error generating ECDSA key:", err)
	}
}

func generateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func GetPublicKey() *ecdsa.PublicKey {
	return &JWT_KEY.PublicKey
}
