package router

import (
	"enbektap/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(controller *controllers.VacancyController, r *gin.Engine) {
	r.POST("/createvacancy", controller.CreateVacancy)
	r.GET("/vacancies", controller.GetAllVacancies)
	r.GET("/vacancy", controller.GetVacancy)
	r.DELETE("/deletevacancy", controller.DeleteVacancy)
	r.PUT("/updatevacancy", controller.UpdateVacancy)
}
