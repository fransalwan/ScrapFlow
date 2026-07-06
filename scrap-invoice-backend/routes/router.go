package routes

import (
	"scrap-invoice-backend/controllers"
	"scrap-invoice-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	// === PUBLIC ROUTES (Tanpa Login) ===
	api.POST("/login", controllers.Login)

	// === PROTECTED ROUTES (Wajib Login) ===
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware()) // <-- Semua route di bawah ini butuh login
	{
		// Invoice
		invoices := protected.Group("/invoices")
		{
			invoices.GET("", controllers.GetInvoices)
			invoices.GET("/:id", controllers.GetInvoiceByID)
			invoices.POST("", controllers.CreateInvoice)
			invoices.PUT("/:id", controllers.UpdateInvoice)
			invoices.DELETE("/:id", controllers.DeleteInvoice)

			invoices.GET("/:id/scales", controllers.GetScaleDetailsByInvoiceID)
			invoices.POST("/:id/scales", controllers.CreateScaleDetail)
			invoices.PUT("/:id/scales", controllers.UpdateScaleDetail)
			invoices.DELETE("/:id/scales", controllers.DeleteScaleDetail)

			invoices.GET("/:id/summary", controllers.GetScaleSummaryByInvoice)
			invoices.GET("/:id/print", controllers.GetInvoicePDF)
		}

		// Customers
		customers := protected.Group("/customers")
		{
			customers.GET("", controllers.GetCustomers)
			customers.POST("", controllers.CreateCustomer)
			customers.PUT("/:id", controllers.UpdateCustomer)
			customers.DELETE("/:id", controllers.DeleteCustomer)
		}

		// Categories
		categories := protected.Group("/categories")
		{
			categories.GET("", controllers.GetCategories)
			categories.POST("", controllers.CreateCategory)
			categories.PUT("/:id", controllers.UpdateCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
		}

		// Items
		items := protected.Group("/items")
		{
			items.GET("", controllers.GetItems)
			items.POST("", controllers.CreateItem)
			items.PUT("/:id", controllers.UpdateItem)
			items.DELETE("/:id", controllers.DeleteItem)
		}

		// Dashboard stats
		dashboard := protected.Group("/dashboard")
		{
			dashboard.GET("/invoices/count", controllers.GetInvoiceCount)
			dashboard.GET("/customers/count", controllers.GetCustomerCount)
		}
	}
}
