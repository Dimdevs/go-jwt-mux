package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jeypac/go-jwt-mux/config"
	"github.com/jeypac/go-jwt-mux/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request header or cookie
		tokenString := extractToken(r)

		// Parse the JWT token
		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.GetPublicKey(), nil
		})

		if err != nil {
			// Handle token parsing errors
			handleTokenError(w, err)
			return
		}

		if !token.Valid {
			// Handle invalid tokens
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// If the token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) string {
	// Your logic to extract the token from the request header or cookie
	// For example, from the Authorization header or a cookie named "token"
	authorizationHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		return strings.TrimPrefix(authorizationHeader, "Bearer ")
	}

	cookie, err := r.Cookie("token")
	if err == nil {
		return cookie.Value
	}

	return ""
}

func handleTokenError(w http.ResponseWriter, err error) {
	// Handle token parsing errors
	response := map[string]string{"error": err.Error(), "message": "Unauthorized"}
	helper.ResponseJSON(w, http.StatusUnauthorized, response)
}
