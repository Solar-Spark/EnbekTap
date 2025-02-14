package services

import (
	"transactions/entities"
	"transactions/infra"
)

type TransactionService struct{
	Repo infra.TransactionRepository
}

func(s *TransactionService) CreateTransaction(transaction entities.Transaction)error{
	return s.Repo.CreateTransaction(transaction)
}

func(s *TransactionService) ReadTransaction(id uint)(entities.Transaction, error){
	return s.Repo.ReadTransaction(id)
}

func(s *TransactionService) ReadTransactions(email string)([]entities.Transaction, error){
	return s.Repo.ReadTransactions(email)
}

func(s *TransactionService) DeleteTransaction(id uint)error{
	return s.Repo.DeleteTransaction(id)
}