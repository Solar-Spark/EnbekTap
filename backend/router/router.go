package router

import (
	"enbektap/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(vacancyController *controllers.VacancyController, UserController *controllers.UserController, r *gin.Engine) {
	r.POST("/createvacancy", vacancyController.CreateVacancy)
	r.GET("/vacancies", vacancyController.GetAllVacancies)
	r.GET("/vacancy", vacancyController.GetVacancy)
	r.DELETE("/deletevacancy", vacancyController.DeleteVacancy)
	r.PUT("/updatevacancy", vacancyController.UpdateVacancy)
	r.POST("/support/contact", controllers.ContactSupportHandler)
	r.POST("/signup", UserController.Signup)
	r.POST("/login", UserController.Login)
	//admin routes
	r.POST("/admin/createuser", UserController.CreateUser)
	r.GET("/admin/users", UserController.GetAllUsers)
	r.GET("/admin/user", UserController.GetUser)
	r.DELETE("/admin/deleteuser", UserController.DeleteUser)
	r.PUT("/admin/updateuser", UserController.UpdateUser)
}
