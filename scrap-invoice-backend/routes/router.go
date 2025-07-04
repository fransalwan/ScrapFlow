package routes

import (
	"scrap-invoice-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// inisialisasi CRUD customer
	customer := r.Group("/api")
	{
		customer.GET("/customers", controllers.GetCustomers)
		customer.POST("/customer", controllers.CreateCustomer)
		customer.PUT("/customer/:id", controllers.UpdateCustomer)
		customer.DELETE("/customer/:id", controllers.DeleteCustomer)
	}

}
