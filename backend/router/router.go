package router

import (
	"enbektap/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(vacancyController *controllers.VacancyController, UserController *controllers.UserController, MicroController *controllers.MicroserviceController, r *gin.Engine) {
	auth := r.Group("/auth")
	api := r.Group("/api")
	admin := r.Group("/admin")
	//Authorized routes
	auth.POST("/createvacancy", vacancyController.CreateVacancy)
	auth.GET("/vacancies", vacancyController.GetAllVacancies)
	auth.GET("/vacancy", vacancyController.GetVacancy)
	auth.POST("/support/contact", controllers.ContactSupportHandler)
	auth.POST("/logout", UserController.Logout)
	auth.GET("/profile", UserController.Profile)
	auth.POST("/createtransaction", MicroController.CreateTransaction)
	auth.GET("/transaction")
	auth.DELETE("/deletetransaction")
	auth.GET("/transactions", MicroController.GetAllTransactions)
	//Routes
	api.POST("/signup", UserController.Signup)
	api.POST("/login", UserController.Login)
	api.POST("/send-code", UserController.SendCode)
	//admin routes
	admin.DELETE("/deletevacancy", vacancyController.DeleteVacancy)
	admin.PUT("/updatevacancy", vacancyController.UpdateVacancy)
	admin.POST("/createuser", UserController.CreateUser)
	admin.GET("/users", UserController.GetAllUsers)
	admin.GET("/user", UserController.GetUser)
	admin.DELETE("/deleteuser", UserController.DeleteUser)
	admin.PUT("/updateuser", UserController.UpdateUser)
}
