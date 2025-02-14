package controllers

import (
	"bytes"
	"enbektap/dto"
	"enbektap/services"
	"enbektap/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type MicroserviceController struct{
	UserService *services.UserService
}

func(c *MicroserviceController) CreateTransaction(ctx *gin.Context){
	if ctx.Request.Method != http.MethodPost{
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	url := "http://localhost:8081/createtransaction"
	var receivingdata dto.TransactionRequest
	if err := ctx.ShouldBindJSON(&receivingdata); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if receivingdata.CardNumber == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Card Number is empty"})
		return
	}
	if receivingdata.CardName == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Card Name is empty"})
		return
	}
	if receivingdata.CVV == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "CVV is empty"})
		return
	}
	if receivingdata.Amount < 0{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price number"})
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

	sendingdata := dto.TransactionSender{
		UserEmail: email,
		Amount: receivingdata.Amount,
		PaymentMethod: receivingdata.PaymentMethod,
		CardNumber: receivingdata.CardNumber,
		CardName: receivingdata.CardName,
		CVV: receivingdata.CVV,
	}
	jsonData, err := json.Marshal(sendingdata)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encode data"})
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction creation request sent"})
}

func(c *MicroserviceController) GetAllTransactions(ctx *gin.Context){
	if ctx.Request.Method != http.MethodGet{
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	url := "http://localhost:8081/transactions"
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
			return
	}
	q := req.URL.Query()
	q.Add("email", email)
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
			return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
			return
	}
	type APIResponse struct {
			Transactions []dto.TransactionResponse `json:"transactions"`
	}

	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON"})
			return
	}

	transactions := apiResponse.Transactions

var filteredTransactions []gin.H
for _, t := range transactions {
    filteredTransactions = append(filteredTransactions, gin.H{
        "transaction_id": t.TransactionID,
        "status":         t.Status,
    })
}

	ctx.JSON(http.StatusOK, filteredTransactions)
}