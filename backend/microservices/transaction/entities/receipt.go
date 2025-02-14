package entities

import "time"

type Receipt struct {
	CompanyName   string
	TransactionID int64
	OrderDate     time.Time
	Items         []Item
	CustomerName  string
	PaymentMethod string
	TotalAmount   int
}

type Item struct {
	Name     string
	UnitPrice int
}