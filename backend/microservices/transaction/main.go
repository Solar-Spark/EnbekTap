package main

import (
	"fmt"
	"net/http"
	"time"
	"transactions/controllers"
	"transactions/infra"
	"transactions/router"
	"transactions/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func main() {
	// 	receipt := entities.Receipt{
	// 	CompanyName:   "EnbekTap",
	// 	TransactionID: 2,
	// 	OrderDate:     time.Now(),
	// 	CustomerName:  "John Doe",
	// 	PaymentMethod: "Credit Card (**** **** **** 1234)", // Example of masked card details
	// 	TotalAmount:   250.00,
	// 	Items: []entities.Item{
	// 		{Name: "Premium Subscription", UnitPrice: 50.00},
	// 	},
	// }

	// // Generate the PDF receipt
	// err := utils.GenerateReceiptPDF(receipt, "receipt.pdf")
	// if err != nil {
	// 	log.Fatal()
	// }
	db, err := infra.ConnectDB()
	if err != nil{
		fmt.Println("Failed to connect to database")
		return
	}

	transactionRepo := &infra.TransactionRepo{DB: db}
	transactionService := &services.TransactionService{Repo: transactionRepo}
	transactionController := &controllers.TransactionController{TransactionService: transactionService}

	r := gin.Default()

		r.Use(func(c *gin.Context) {
		start := time.Now()

		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Next()

		log.Info().
			Str("request_id", requestID).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("query", c.Request.URL.RawQuery).
			Str("client_ip", c.ClientIP()).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Int("body_size", c.Writer.Size()).
			Str("user_agent", c.Request.UserAgent()).
			Msg("Request processed")
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},                   // Allow all origins (restrict in production)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow essential HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "ngrok-skip-browser-warning", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"}, // Allow client-side access to these headers
		AllowCredentials: true,                                       // Allow cookies/auth credentials if needed
		AllowWildcard:    true,                                       // Support wildcard subdomains
		AllowFiles:       true,                                       // Allow file uploads
		MaxAge:           12 * time.Hour,                             // Cache preflight requests for 12 hours
	}))

	router.SetupRoutes(transactionController, r)

		go func() {
		log.Info().Msg("Server starting on port 8081")
		if err := http.ListenAndServe(":8081", r); err != nil {
			log.Fatal().
				Err(err).
				Msg("Server failed to start")
		}
	}()
	select {}
}