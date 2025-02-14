package controllers

import (
	enbektapServices "enbektap/services"
	enbektapUtils "enbektap/utils"
	"log"
	"net/http"
	"strconv"
	"time"
	"transactions/dto"
	"transactions/entities"
	"transactions/services"
	"transactions/utils"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{
	TransactionService *services.TransactionService
	UserService *enbektapServices.UserService
}

func(c *TransactionController) CreateTransaction(ctx *gin.Context){
	if ctx.Request.Method != http.MethodPost{
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	var transactionDTO dto.TransactionDTO
	var transaction entities.Transaction
	if err := ctx.ShouldBindJSON(&transactionDTO); err != nil{
		log.Fatal(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if transactionDTO.UserEmail == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is empty"})
		return
	}
	if transactionDTO.Amount < 0{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price number"})
		return
	}
	message, ok := utils.PriceCheck(transactionDTO.Amount)
	if !ok{
		ctx.JSON(http.StatusPaymentRequired, gin.H{"error": message})
		return
	}
	transaction.UserEmail = transactionDTO.UserEmail
	transaction.Amount = transactionDTO.Amount
	transaction.Status = "pending"
	if err := c.TransactionService.CreateTransaction(transaction); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	paymentmethod := transactionDTO.PaymentMethod + " " + utils.MaskCard(transactionDTO.CardNumber)

	receipt := entities.Receipt{
		CompanyName:   "EnbekTap",
		TransactionID: transaction.TransactionID,
		OrderDate:     time.Now(),
		CustomerName:  transactionDTO.CardName,
		PaymentMethod: paymentmethod,
		TotalAmount:   transaction.Amount,
		Items: []entities.Item{
			{Name: "Donation", UnitPrice: 5000},
		},
	}

	err := utils.GenerateReceiptPDF(receipt, "receipt.pdf")
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate receipt"})
	}
	attachments := []string{"receipt.pdf"}
	enbektapUtils.SendEmailWAtt(transaction.UserEmail, message, attachments)
	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction created: " + message, "status": transaction.Status, "id": transaction.TransactionID, "amount": transaction.Amount, "payment_method": paymentmethod, "date": receipt.OrderDate})
}

func(c *TransactionController) ReadTransaction(ctx *gin.Context){
	if ctx.Request.Method != http.MethodGet{
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	idStr := ctx.DefaultQuery("id", "")
	if idStr == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Transaction ID is empty"})
		return
	}

	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}
	id := uint(id64)
	transaction, err := c.TransactionService.ReadTransaction(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

func(c *TransactionController) ReadTransactions(ctx *gin.Context){
	if ctx.Request.Method != http.MethodGet{
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Invalid request method"})
		return
	}
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is required"})
		return
	}
	transaction, err := c.TransactionService.ReadTransactions(email)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}

	var transactions []gin.H
	for _, t := range transaction {
		transactions = append(transactions, gin.H{
			"TransactionID": t.TransactionID,
			"Status":        t.Status,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func(c *TransactionController) DeleteTransaction(ctx *gin.Context){
	if ctx.Request.Method != http.MethodDelete{
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
	if err := c.TransactionService.DeleteTransaction(id); err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "transaction deleted"})
}