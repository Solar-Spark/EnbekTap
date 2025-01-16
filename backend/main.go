package main

import (
	"enbektap/controllers"
	"enbektap/infra"
	"enbektap/middleware"
	"enbektap/router"
	"enbektap/services"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func setupLogger() {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatal().Err(err).Msg("Failed to create log directory")
	}

	logFile := filepath.Join(logDir, time.Now().Format("2006-01-02")+".json")
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create log file")
	}

	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	log.Logger = zerolog.New(multi).
		With().
		Timestamp().
		Caller().
		Logger()
}

func main() {
	setupLogger()

	db, err := infra.ConnectDB()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to connect to database")
		return
	}
	log.Info().Msg("Database connection established")

	repo := &infra.VacancyRepo{DB: db}
	service := &services.VacancyService{Repo: repo}
	controller := &controllers.VacancyController{Service: service}

	r := gin.Default()

	r.Use(middleware.RateLimiterMiddleware())

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
		AllowOrigins:     []string{"*"}, // You can set specific origins here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.SetupRoutes(controller, r)

	log.Info().Msg("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal().
			Err(err).
			Msg("Server failed to start")
	}
}
