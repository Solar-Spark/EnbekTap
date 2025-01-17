package controllers

import (
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func ContactSupportHandler(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form data"})
		return
	}

	subject := c.Request.FormValue("subject")
	message := c.Request.FormValue("message")

	if subject == "" || message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Subject and message are required"})
		return
	}

	attachments := []string{}
	formFiles := c.Request.MultipartForm.File["attachments"]
	for _, fileHeader := range formFiles {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
			return
		}
		defer file.Close()

		tempDir := "uploads"
		if _, err := os.Stat(tempDir); os.IsNotExist(err) {
			os.Mkdir(tempDir, os.ModePerm)
		}

		filePath := filepath.Join(tempDir, fileHeader.Filename)
		out, err := os.Create(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error copying file"})
			return
		}

		attachments = append(attachments, filePath)
	}

	if err := sendEmail(subject, message, attachments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	for _, file := range attachments {
		os.Remove(file)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Message sent successfully"})
}

func sendEmail(subject, message string, attachments []string) error {
	from := "erni100105@gmail.com"
	password := "lwgv hdie kzjq jqim"
	to := "erni100105@gmail.com"
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	if _, err := mail.ParseAddress(from); err != nil {
		return fmt.Errorf("invalid sender email address: %w", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	if err := attachFiles(m, attachments); err != nil {
		return fmt.Errorf("failed to attach files: %w", err)
	}

	d := gomail.NewDialer(smtpHost, smtpPort, from, password)

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
