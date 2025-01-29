package entities

type User struct {
	UserID           int64  `gorm:"primaryKey;autoIncrement"`
	Email            string `gorm:"type:varchar(255);column:email"`
	Password         string `gorm:"type:varchar(255);column:password"`
	Role             string `gorm:"type:enum('admin', 'user');column:role"`
	FullName         string `gorm:"type:varchar(255);column:full_name"`
	Verified         bool   `gorm:"column:verified;default:false"`
	VerificationCode string `gorm:"type:varchar(255);column:verification_token"`
}
