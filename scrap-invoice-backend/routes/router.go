package routes

import (
	"scrap-invoice-backend/controllers"
	"scrap-invoice-backend/middleware.go"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	login := r.Group("/api")
	{
		login.POST("/login", controllers.Login)

		// Protected route (contoh)
		login.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Lo berhasil akses route yang diproteksi!"})
		})
	}

	// inisialisasi CRUD invoice
	invoice := r.Group("/api")
	{
		invoice.GET("/invoices", controllers.GetInvoices)
		invoice.GET("/invoice/:id", controllers.GetInvoiceByID)
		invoice.POST("/invoice", controllers.CreateInvoice)
		invoice.PUT("/invoice/:id", controllers.UpdateInvoice)
		invoice.DELETE("/invoice/:id", controllers.DeleteInvoice)

		invoice.GET("/invoice/:id/scales", controllers.GetScaleDetailsByInvoiceID)
		invoice.POST("/invoice/:id/scales", controllers.CreateScaleDetail)
		invoice.PUT("/invoice/:id/scales", controllers.UpdateScaleDetail)
		invoice.DELETE("/invoice/:id/scales", controllers.DeleteScaleDetail)

		invoice.GET("/invoice/:id/summary", controllers.GetScaleSummaryByInvoice)
		invoice.GET("/invoice/:id/print", controllers.GetInvoicePDF)
	}

	dashboard := r.Group("/api")
	{
		dashboard.GET("/invoices/count", controllers.GetInvoiceCount)
		dashboard.GET("/customers/count", controllers.GetCustomerCount)
	}

	// inisialisasi CRUD customer
	customer := r.Group("/api")
	{
		customer.GET("/customers", controllers.GetCustomers)
		customer.POST("/customer", controllers.CreateCustomer)
		customer.PUT("/customer/:id", controllers.UpdateCustomer)
		customer.DELETE("/customer/:id", controllers.DeleteCustomer)
	}

	// inisialisasi CRUD Category Item
	category := r.Group("/api")
	{
		category.GET("/categories", controllers.GetCategories)
		category.POST("/category", controllers.CreateCategory)
		category.PUT("/category/:id", controllers.UpdateCategory)
		category.DELETE("/category/:id", controllers.DeleteCategory)
	}

	item := r.Group("/api")
	{
		item.GET("/items", controllers.GetItems)
		item.POST("/item", controllers.CreateItem)
		item.PUT("/item/:id", controllers.UpdateItem)
		item.DELETE("/item/:id", controllers.DeleteItem)
	}

}
