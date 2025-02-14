package infra

import (
	"transactions/entities"

	"gorm.io/gorm"
)

type TransactionRepo struct{
	DB *gorm.DB
}

type TransactionRepository interface{
	CreateTransaction(transaction entities.Transaction) error
	ReadTransaction(id uint)(entities.Transaction, error)
	ReadTransactions(email string)([]entities.Transaction, error)
	DeleteTransaction(id uint)error
}

func(repo *TransactionRepo) CreateTransaction(transaction entities.Transaction)error{
	return repo.DB.Create(&transaction).Error
}

func(repo *TransactionRepo) ReadTransaction(id uint)(entities.Transaction, error){
	var transaction entities.Transaction
	err := repo.DB.First(&transaction, id).Error
	return transaction, err
}

func(repo *TransactionRepo) ReadTransactions(email string)([]entities.Transaction, error){
	var transactions []entities.Transaction

	query := repo.DB.Model(&entities.Transaction{}).Where("user_email = ?", email)
	err := query.Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func(repo *TransactionRepo) DeleteTransaction(id uint)error{
	return repo.DB.Delete(&entities.Transaction{}, id).Error
}