package main

import (
	"enbektap/controllers"
	"enbektap/infra"
	"enbektap/router"
	"enbektap/services"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db, error := infra.ConnectDB()

	if error != nil {
		return
	}

	repo := &infra.VacancyRepo{DB: db}
	service := &services.VacancyService{Repo: repo}
	controller := &controllers.VacancyController{Service: service}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // You can set specific origins here
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.SetupRoutes(controller, r)

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
