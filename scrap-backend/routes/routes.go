package routes

import (
	"github.com/fransalwan/scrap-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/suppliers", controllers.GetSuppliers)
		api.POST("/suppliers", controllers.CreateSupplier)

		api.GET("/transactions", controllers.GetTransactions)
		api.POST("/transactions", controllers.CreateTransaction)
	}
}
