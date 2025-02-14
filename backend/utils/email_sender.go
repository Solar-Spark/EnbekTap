package utils

import (
	"crypto/tls"
	"fmt"
	"os"
	"path/filepath"

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

func SendEmailWAtt(recipientEmail string, message string, attachments []string) error {
	config := NewEmailConfig()
	m := gomail.NewMessage()
	m.SetHeader("From", config.SenderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Congratulations")
	m.SetBody("text/plain", message)

	if err := attachFiles(m, attachments); err != nil {
		return fmt.Errorf("failed to attach files: %w", err)
	}

	d := gomail.NewDialer(config.SMTPHost, config.SMTPPort, config.SenderEmail, config.SenderPass)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func attachFiles(m *gomail.Message, attachments []string) error {
	for _, file := range attachments {
		// Проверяем, существует ли файл
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("attachment not found: %s", file)
		}

		m.Attach(file, gomail.SetHeader(map[string][]string{
			"Content-Disposition": {fmt.Sprintf(`attachment; filename="%s"`, filepath.Base(file))},
		}))
	}
	return nil
}