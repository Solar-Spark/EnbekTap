package utils

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
)

func TokenParse(authHeader string) (string, error) {
	if authHeader == "" {
		err := errors.New("Authorization header missing")
		return "", err
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSecretKey()), nil
	})
	if err != nil || !token.Valid {
		err := errors.New("Invalid or expired JWT Token")
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := errors.New("Failed to parse claims")
		return "", err
	}
	email, exists := claims["username"].(string)
	if !exists {
		err := errors.New("Email claim not found in token")
		return "", err
	}
	return email, nil
}
