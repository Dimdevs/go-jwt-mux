package config

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY *ecdsa.PrivateKey

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}

func init() {
	var err error
	JWT_KEY, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err) // Handle the error according to your application's needs
	}
}
