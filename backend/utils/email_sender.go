package utils

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	SMTPHost    string
	SMTPPort    int
	SenderEmail string
	SenderPass  string
}

func NewEmailConfig() *EmailConfig {
	return &EmailConfig{
		SMTPHost:    "smtp.gmail.com",
		SMTPPort:    587,
		SenderEmail: "erni100105@gmail.com",
		SenderPass:  "truf wxhe yztg apvp",
	}
}

func SendEmail(recipientEmail, message string) error {
	config := NewEmailConfig()

	m := gomail.NewMessage()
	m.SetHeader("From", config.SenderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Verification Code")
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(
		config.SMTPHost,
		config.SMTPPort,
		config.SenderEmail,
		config.SenderPass,
	)

	d.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         config.SMTPHost,
	}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
