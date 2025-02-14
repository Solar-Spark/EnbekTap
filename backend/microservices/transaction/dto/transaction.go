package dto

type TransactionDTO struct {
	UserEmail     string `json:"email" binding:"email"`
	CardNumber    string `json:"card" binding:"required"`
	CVV           string `json:"cvv" binding:"required"`
	CardName      string `json:"name" binding:"required"`
	PaymentMethod string `json:"method" binding:"required"`
	Amount        int    `json:"amount" binding:"required"`
}