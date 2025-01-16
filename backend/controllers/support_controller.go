package controllers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type SupportController struct {
	supportEmail string
	smtpHost     string
	smtpPort     int
	smtpUser     string
	smtpPass     string
}

func NewSupportController() *SupportController {
	return &SupportController{
		supportEmail: "erni100105@gmail.com",
		smtpHost:     "smtp.gmail.com",
		smtpPort:     587,
		smtpUser:     "erni100105@gmail.com",
		smtpPass:     "qthe fpnb vivg zotd",
	}
}

func (sc *SupportController) HandleSupportMessage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	if c.Request.Method == "OPTIONS" {
		c.Status(http.StatusOK)
		return
	}

	subject := c.PostForm("subject")
	message := c.PostForm("message")

	if subject == "" || message == "" {
		log.Printf("Validation failed: subject or message empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subject and message are required"})
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sc.smtpUser)
	m.SetHeader("To", sc.supportEmail)
	m.SetHeader("Subject", "Support Request: "+subject)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(sc.smtpHost, sc.smtpPort, sc.smtpUser, sc.smtpPass)

	// Configure TLS properly
	d.TLSConfig = &tls.Config{
		ServerName:         "smtp.gmail.com",
		InsecureSkipVerify: false,
		MinVersion:         tls.VersionTLS12,
	}

	// Test SMTP connection first
	s, err := d.Dial()
	if err != nil {
		log.Printf("SMTP connection error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to connect to email server",
		})
		return
	}
	s.Close()

	// Send email
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Email sending error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Failed to send email: %v", err),
		})
		return
	}

	log.Printf("Email sent successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Support message sent successfully"})
}
