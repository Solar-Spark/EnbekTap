package controllers

import (
	"enbektap/dto"
	"enbektap/entities"
	"enbektap/services"
	"enbektap/utils"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var VerificationCodes = make(map[string]string)

type UserController struct {
	Service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Logout(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Log out successfully"})
}

func (c *UserController) Profile(ctx *gin.Context) {
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

	user, err := c.Service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Success",
		"FullName": user.FullName,
		"Email":    user.Email,
		"Role":     user.Role,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	var login dto.Login
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	user, err := c.Service.GetUserByEmail(login.Login)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	token, err := utils.GenerateJWT(login.Login)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create JWT code"})
		return
	}
	ctx.SetCookie("access_token", token, 3600*24, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Log in success", "jwt": token})
}

func (c *UserController) Signup(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodPost {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}
	if code, ok := VerificationCodes[user.Email]; !ok || code != user.VerificationCode {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Verified = true
	user.Password = string(hashedPassword)
	if err := c.Service.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (c *UserController) SendCode(ctx *gin.Context) {
	var codeData struct {
		Email string `json:"email"`
	}
	if err := ctx.ShouldBindJSON(&codeData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	code := utils.GenerateCode()
	fmt.Println(code)
	fmt.Println(codeData.Email)
	utils.SendEmail(codeData.Email, code)
	VerificationCodes[codeData.Email] = code
	ctx.JSON(http.StatusOK, gin.H{"message": "Code sent successfully"})
}

func (c *UserController) CreateUser(ctx *gin.Context) {
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

	admin, err := c.Service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if admin.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Lack of permission"})
	}
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := c.Service.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	if ctx.Request.Method != http.MethodGet {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid requset method"})
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

	admin, err := c.Service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if admin.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Lack of permission"})
	}
	id := uint(id64)
	user, err := c.Service.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
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

	admin, err := c.Service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if admin.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Lack of permission"})
	}
	users, err := c.Service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
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
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
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

	admin, err := c.Service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if admin.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Lack of permission"})
	}
	if err := c.Service.UpdateUser(id, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
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

	admin, err := c.Service.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if admin.Role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Lack of permission"})
	}
	id := uint(id64)
	if err := c.Service.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted Successfully"})
}
