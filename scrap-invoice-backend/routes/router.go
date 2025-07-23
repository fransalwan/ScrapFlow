package routes

import (
	"scrap-invoice-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

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
