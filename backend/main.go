package main

import (
	"enbektap/controllers"
	"enbektap/infra"
	"enbektap/middleware"
	"enbektap/router"
	"enbektap/services"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
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

func startNgrok() (string, error) {
	cmd := exec.Command("ngrok", "http", "8080", "--log", "stdout")

	if err := cmd.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start Ngrok")
		return "", err
	}

	time.Sleep(3 * time.Second)

	resp, err := http.Get("http://127.0.0.1:4040/api/tunnels")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	tunnels, ok := result["tunnels"].([]interface{})
	if !ok || len(tunnels) == 0 {
		return "", fmt.Errorf("no tunnels found")
	}

	firstTunnel := tunnels[0].(map[string]interface{})
	publicURL, ok := firstTunnel["public_url"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get public URL")
	}

	return publicURL, nil
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

	vacancyRepo := &infra.VacancyRepo{DB: db}
	vacancyService := &services.VacancyService{Repo: vacancyRepo}
	vacancyController := &controllers.VacancyController{Service: vacancyService}
	userRepo := &infra.UserRepo{DB: db}
	userService := &services.UserService{Repo: userRepo}
	userController := &controllers.UserController{Service: userService}

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
		AllowOrigins:     []string{"*"},                                       // Allow all origins (restrict in production)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allow essential HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "ngrok-skip-browser-warning", "Accept", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"}, // Allow client-side access to these headers
		AllowCredentials: true,                                       // Allow cookies/auth credentials if needed
		AllowWildcard:    true,                                       // Support wildcard subdomains
		AllowFiles:       true,                                       // Allow file uploads
		MaxAge:           12 * time.Hour,                             // Cache preflight requests for 12 hours
	}))

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, ngrok-skip-browser-warning, Accept, X-Requested-With")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Status(http.StatusOK)
	})

	router.SetupRoutes(vacancyController, userController, r)

	go func() {
		log.Info().Msg("Server starting on port 8080")
		if err := http.ListenAndServe(":8080", r); err != nil {
			log.Fatal().
				Err(err).
				Msg("Server failed to start")
		}
	}()

	ngrokURL, err := startNgrok()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start Ngrok")
	}

	log.Info().Msgf("Ngrok tunnel established: %s", ngrokURL)

	select {}
}
