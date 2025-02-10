package controllers

import (
	"enbektap/entities"
	"enbektap/services"
	"enbektap/utils"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type VacancyController struct {
	VacancyService *services.VacancyService
	UserService    *services.UserService
}

type Options struct {
	FilterBy string `json:"filterBy"`
	SortBy   string `json:"sortBy"`
}

func (c *VacancyController) CreateVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "Bearer null" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is empty"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetSecretKey()), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JWT"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
		return
	}
	email, exists := claims["username"].(string)
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	_, err = c.UserService.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	var vacancy entities.Vacancy
	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := c.VacancyService.CreateVacancy(vacancy); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vacancy"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Vacancy created successfully"})
}

func (c *VacancyController) GetVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "Bearer null" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is empty"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetSecretKey()), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JWT"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
		return
	}
	email, exists := claims["username"].(string)
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	_, err = c.UserService.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	id := uint(id64)
	vacancy, err := c.VacancyService.GetVacancy(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	ctx.JSON(http.StatusOK, vacancy)
}

func (c *VacancyController) GetAllVacancies(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "Bearer null" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is empty"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetSecretKey()), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JWT"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
		return
	}
	email, exists := claims["username"].(string)
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	_, err = c.UserService.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	filterBy := ctx.DefaultQuery("jobType", "none")
	sortBy := ctx.DefaultQuery("sort", "none")
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize := 9

	vacancies, total, err := c.VacancyService.GetAllVacancies(filterBy, sortBy, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vacancies"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"vacancies":   vacancies,
		"total":       total,
		"currentPage": page,
		"pageSize":    pageSize,
		"totalPages":  int(math.Ceil(float64(total) / float64(pageSize))),
	})
}

func (c *VacancyController) UpdateVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPut {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	id := uint(id64)
	var vacancy entities.Vacancy
	if err := ctx.ShouldBindJSON(&vacancy); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "Bearer null" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is empty"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetSecretKey()), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JWT"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
		return
	}
	email, exists := claims["username"].(string)
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	admin, err := c.UserService.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if admin.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Lack of permission"})
	}

	if err := c.VacancyService.UpdateVacancy(id, vacancy); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update vacancy"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vacancy updated successfully"})
}

func (c *VacancyController) DeleteVacancy(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodDelete {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}

	idStr := ctx.DefaultQuery("id", "")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing id parameter"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "Bearer null" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token is empty"})
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.GetSecretKey()), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JWT"})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
		return
	}
	email, exists := claims["username"].(string)
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not Found"})
		return
	}

	admin, err := c.UserService.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if admin.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Lack of permission"})
	}

	id := uint(id64)
	if err := c.VacancyService.DeleteVacancy(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete vacancy"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vacancy deleted successfully"})
}
