package main

import (
	"enbektap/controllers"
	"enbektap/infra"
	"enbektap/router"
	"enbektap/services"
	"log"
	"net/http"
)

func main() {
	// Database connection
	db, error := infra.ConnectDB()

	if error != nil {
		return
	}

	// Dependency injection
	repo := &infra.VacancyRepo{DB: db}
	service := &services.VacancyService{Repo: repo}
	controller := &controllers.VacancyController{Service: service}

	// Setup routes
	router.SetupRoutes(controller)

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
