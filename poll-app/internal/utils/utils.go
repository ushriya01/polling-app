package utils

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("your-secret-key")

// ParseToken parses and validates the JWT token from the request header
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	// Extract the token from the Authorization header
	tokenParts := strings.Split(tokenString, " ")
	if len(tokenParts) != 2 {
		return nil, errors.New("invalid token format")
	}

	tokenString = tokenParts[1]

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the token signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return SigningKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Verify the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
