package router

import (
	"enbektap/controllers"
	"net/http"
)

func SetupRoutes(controller *controllers.VacancyController) {
	http.HandleFunc("/createvacancy", controller.CreateVacancy)
	http.HandleFunc("/vacancies", controller.GetAllVacancies)
	http.HandleFunc("/vacancy", controller.GetVacancy)
	http.HandleFunc("/deletevacancy", controller.DeleteVacancy)
	http.HandleFunc("/updatevacancy", controller.UpdateVacancy)
}
