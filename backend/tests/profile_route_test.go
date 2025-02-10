package tests

import (
	"enbektap/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Fail on emprty token
func TestProfileRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	uc := controllers.NewUserController()
	r := gin.Default()
	auth := r.Group("/auth")
	auth.GET("/profile", uc.Profile)
	req, err := http.NewRequest(http.MethodGet, "/auth/profile", nil)
	assert.NoError(t, err)
	req.Header.Set("Authorization", "Bearer ") //empty token
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Contains(t, rr.Body.String(), "Failed to parse JWT")
}
