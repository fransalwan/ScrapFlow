package routes

import (
	"scrap-invoice-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")

	// // Auth
	// api.POST("/login", controllers.Login)
	// api.GET("/protected", middleware.AuthMiddleware(), func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Lo berhasil akses route yang diproteksi!"})
	// })

	// Invoice
	invoices := api.Group("/invoices")
	{
		invoices.GET("", controllers.GetInvoices)          // GET /api/invoices
		invoices.GET("/:id", controllers.GetInvoiceByID)   // GET /api/invoices/:id
		invoices.POST("", controllers.CreateInvoice)       // POST /api/invoices
		invoices.PUT("/:id", controllers.UpdateInvoice)    // PUT /api/invoices/:id
		invoices.DELETE("/:id", controllers.DeleteInvoice) // DELETE /api/invoices/:id

		invoices.GET("/:id/scales", controllers.GetScaleDetailsByInvoiceID)
		invoices.POST("/:id/scales", controllers.CreateScaleDetail)
		invoices.PUT("/:id/scales", controllers.UpdateScaleDetail)
		invoices.DELETE("/:id/scales", controllers.DeleteScaleDetail)

		invoices.GET("/:id/summary", controllers.GetScaleSummaryByInvoice)
		invoices.GET("/:id/print", controllers.GetInvoicePDF)
	}

	customers := api.Group("/customers")
	{
		customers.GET("", controllers.GetCustomers)
		customers.POST("", controllers.CreateCustomer)
		customers.PUT("/:id", controllers.UpdateCustomer)
		customers.DELETE("/:id", controllers.DeleteCustomer)
	}

	// Categories
	categories := api.Group("/categories")
	{
		categories.GET("", controllers.GetCategories) // /api/categories
		categories.POST("", controllers.CreateCategory)
		categories.PUT("/:id", controllers.UpdateCategory)
		categories.DELETE("/:id", controllers.DeleteCategory)
	}

	// Items
	items := api.Group("/items")
	{
		items.GET("", controllers.GetItems) // /api/items
		items.POST("", controllers.CreateItem)
		items.PUT("/:id", controllers.UpdateItem)
		items.DELETE("/:id", controllers.DeleteItem)
	}

	// Dashboard stats
	dashboard := api.Group("/dashboard")
	{
		dashboard.GET("/invoices/count", controllers.GetInvoiceCount)
		dashboard.GET("/customers/count", controllers.GetCustomerCount)
	}

}
