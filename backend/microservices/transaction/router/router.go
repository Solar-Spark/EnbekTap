package router

import (
	"transactions/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(controller *controllers.TransactionController, r *gin.Engine){
	r.POST("/createtransaction", controller.CreateTransaction)
	r.GET("/transaction", controller.ReadTransaction)
	r.GET("/transactions", controller.ReadTransactions)
	r.DELETE("/deletetransaction", controller.DeleteTransaction)
}
