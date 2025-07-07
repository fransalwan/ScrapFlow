package routes

import (
	"log"
	"scrap-invoice-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	log.Println("Setting up CORS middleware...")

	// inisialisasi CRUD customer
	customer := r.Group("/api")
	{
		customer.GET("/customers", controllers.GetCustomers)
		customer.POST("/customer", controllers.CreateCustomer)
		customer.PUT("/customer/:id", controllers.UpdateCustomer)
		customer.DELETE("/customer/:id", controllers.DeleteCustomer)
	}

	// inisialisasi CRUD invoice
	invoice := r.Group("/api")
	{
		invoice.GET("/invoices", controllers.GetInvoices)
		invoice.GET("/invoice/:id", controllers.GetInvoiceByID)
		invoice.POST("/invoice", controllers.CreateInvoice)
		invoice.PUT("/invoice/:id", controllers.UpdateInvoice)
		invoice.DELETE("/invoice/:id", controllers.DeleteInvoice)
	}

	//inisialisasi CRUD scale detail
	scaleDetail := r.Group("/api")
	{
		scaleDetail.GET("/scale_detail/:id", controllers.GetScaleDetail)
		scaleDetail.POST("/scale_detail", controllers.CreateScaleDetail)
		scaleDetail.PUT("/scale_detail/:id", controllers.UpdateScaleDetail)
		scaleDetail.DELETE("/scale_detail/:id", controllers.DeleteScaleDetail)
	}

	// inisialisasi CRUD summary
	summary := r.Group("/api")
	{
		// summary.GET("/summary/:id", controllers.GetSummary)
		summary.POST("/summary/:invoice_id", controllers.GenerateSummary)
		// summary.PUT("/summary/:id", controllers.UpdateSummary)
		// summary.DELETE("/summary/:id", controllers.DeleteSummary)
	}

	// inisialisasi print pdf
	print := r.Group("/api")
	{
		print.GET("/invoices/:id/pdf", controllers.GetInvoicePDF)
	}
}
