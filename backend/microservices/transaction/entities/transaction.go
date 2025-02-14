package entities

type Transaction struct {
	TransactionID int64  `gorm:"primaryKey;autoIncrement"`
	UserEmail     string `gorm:"type:varchar(255);column:user_email"`
	Amount        int    `gorm:"column:amount"`
	CardNumber    string `gorm:"type:varchar(16);column:cardnumber"`
	PaymentMethod string `gorm:"type:varchar(255);column:method"`
	CVV           string `gorm:"type:varchar(3);column:cvv"`
	Status        string `gorm:"type:enum('Pending the Payment', 'Paid', 'Declined');column:status"`
}