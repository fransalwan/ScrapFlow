package controllers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/dto"
	"scrap-invoice-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetInvoices(c *gin.Context) {
	var invoices []models.Invoice
	fmt.Println("Hit GetInvoices!!!!!!!!!!!!")

	// Ambil semua invoice dengan preload relasi yang diperlukan
	if err := config.DB.
		Preload("Customer").
		Preload("Summaries").
		Preload("Summaries.Item").
		Find(&invoices).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invoices"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Invoices fetched successfully",
		"data":    invoices,
	})
}

func GetInvoiceByID(c *gin.Context) {
	id := c.Param("id")

	var invoice models.Invoice
	if err := config.DB.
		Preload("Customer").
		Preload("Summaries").
		Preload("Summaries.Item").
		First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice found",
		"data":    invoice,
	})
}

func CreateInvoice(c *gin.Context) {
	fmt.Println("Hit CreateInvoice!!!!!!!!!!!!")
	// DEBUG: Log raw request body
	bodyBytes, _ := io.ReadAll(c.Request.Body)
	log.Printf("📦 Raw request body: %s", string(bodyBytes))

	// Reset body buat binding
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// 1. Bind input dari frontend
	var input dto.InvoiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("❌ Bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":         "Invalid JSON: " + err.Error(),
			"received_body": string(bodyBytes),
		})
		return
	}

	log.Printf("✅ Parsed input: %+v", input)

	// 2. Validasi & parse invoice_date (WAJIB untuk create)
	if input.InvoiceDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invoice_date is required for creating invoice",
		})
		return
	}

	parsedInvoiceDate, err := time.Parse("2006-01-02", input.InvoiceDate)
	if err != nil {
		log.Printf("❌ Date parse error: %v (received: %s)", err, input.InvoiceDate)
		c.JSON(http.StatusBadRequest, gin.H{
			"error":           "Invalid invoice_date format. Use YYYY-MM-DD",
			"received":        input.InvoiceDate,
			"expected_format": "YYYY-MM-DD",
		})
		return
	}

	// 3. Gunakan waktu sekarang sebagai created_at
	now := time.Now()

	// 4. Hitung jumlah invoice di tanggal invoice_date (bukan created_at)
	var count int64
	config.DB.Model(&models.Invoice{}).
		Where("DATE(invoice_date) = ?", parsedInvoiceDate.Format("2006-01-02")).
		Count(&count)

	// 5. Generate nomor invoice
	newNumber := fmt.Sprintf("INV-%s-%04d", parsedInvoiceDate.Format("20060102"), count+1)

	// 6. Bangun model invoice dari input DTO
	invoice := input.ToModel(now, parsedInvoiceDate, newNumber)

	// 7. Simpan ke database
	if err := config.DB.Create(&invoice).Error; err != nil {
		log.Printf("❌ Create error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice: " + err.Error()})
		return
	}

	// 8. Ambil ulang data lengkap dengan relasi
	var fullInvoice models.Invoice
	if err := config.DB.Preload("Customer").First(&fullInvoice, invoice.CustomerID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load customer data: " + err.Error()})
		return
	}

	// 9. Kirim response
	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice created successfully",
		"data":    fullInvoice,
	})
}

func DeleteInvoice(c *gin.Context) {
	id := c.Param("id")

	var invoice models.Invoice
	if err := config.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	if err := config.DB.Delete(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invoice"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
}
