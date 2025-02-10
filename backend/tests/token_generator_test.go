package tests

import (
	"enbektap/utils"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	username := "testuser"

	tokenString, err := utils.GenerateJWT(username)
	assert.NoError(t, err, "Error = nil")
	assert.NotEmpty(t, tokenString, "Token shouldnt be empty")

	token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return utils.GetSecretKey(), nil
	})
	assert.NoError(t, err, "Error = nil")
	assert.True(t, token.Valid, "Token should be valid")

	claims, ok := token.Claims.(*utils.Claims)
	assert.True(t, ok, "Claims should be of type *Claims")
	assert.Equal(t, username, claims.Username, "Username in claims should match the provided username")
}
