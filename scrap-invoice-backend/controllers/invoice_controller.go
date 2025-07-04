package controllers

import (
	"fmt"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Optional: generate invoice number kalau belum ada
	if invoice.InvoiceNumber == "" {
		invoice.InvoiceNumber = fmt.Sprintf("INV-%d", time.Now().Unix())
	}

	// Validasi minimum: customer_id dan created_by harus ada
	if invoice.CustomerID == 0 || invoice.CreatedBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "customer_id and created_by are required"})
		return
	}

	// Simpan invoice dan summary (jika ada)
	if err := config.DB.Create(&invoice).Error; err != nil {

		// kasih liat isi invoice untuk debugging
		fmt.Printf("Failed to create invoice:>>>>>>>>>> %+v\n", invoice)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal insert invoice: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice created successfully",
		"data":    invoice,
	})
}
