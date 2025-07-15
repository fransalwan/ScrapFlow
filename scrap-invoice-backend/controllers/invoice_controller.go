package controllers

import (
	"fmt"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/dto"
	"scrap-invoice-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetInvoices(c *gin.Context) {
	var invoices []models.Invoice

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
	// 1. Bind input dari frontend
	var input dto.InvoiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// 2. Parse invoice_date dari FE
	parsedInvoiceDate, err := time.Parse("2006-01-02", input.InvoiceDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice_date format. Use YYYY-MM-DD"})
		return
	}

	// 3. Gunakan waktu sekarang sebagai created_at
	now := time.Now()

	// 4. Hitung jumlah invoice di tanggal yang sama (created_at)
	today := now.Format("2006-01-02")
	var count int64
	config.DB.Model(&models.Invoice{}).
		Where("DATE(created_at) = ?", today).
		Count(&count)

	// 5. Generate nomor invoice dari invoice_date (boleh juga dari now)
	newNumber := fmt.Sprintf("INV-%s-%04d", parsedInvoiceDate.Format("20060102"), count+1)

	// 6. Bangun model invoice dari input DTO
	invoice := input.ToModel(now, parsedInvoiceDate, newNumber)

	// 7. Simpan ke database
	if err := config.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice: " + err.Error()})
		return
	}

	// 8. Ambil ulang data lengkap (dengan relasi)
	var fullInvoice models.Invoice
	if err := config.DB.Preload("Customer").First(&fullInvoice, invoice.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load customer data: " + err.Error()})
		return
	}

	// 9. Kirim response
	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice created successfully",
		"data":    fullInvoice,
	})
}

func UpdateInvoice(c *gin.Context) {
	// 1. Ambil ID dari parameter
	id := c.Param("id")

	// 2. Ambil data invoice dari database
	var invoice models.Invoice
	if err := config.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	// 3. Bind input dari frontend ke DTO
	var input dto.InvoiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// 4. Siapkan map untuk fields yang akan diupdate
	updates := map[string]interface{}{
		"customer_id": input.CustomerID,
	}

	// 5. Handle optional fields
	if input.Status != nil {
		updates["status"] = *input.Status
	}
	if input.PaymentMethod != nil {
		updates["payment_method"] = *input.PaymentMethod
	}
	if input.Note != nil {
		updates["note"] = *input.Note
	}
	if input.InvoiceDate != "" {
		parsedInvoiceDate, err := time.Parse("2006-01-02", input.InvoiceDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice_date format. Use YYYY-MM-DD"})
			return
		}
		updates["invoice_date"] = parsedInvoiceDate
	}

	// 6. Jalankan update ke database
	if err := config.DB.Model(&invoice).Updates(updates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice: " + err.Error()})
		return
	}

	// 7. Ambil ulang invoice + preload relasi
	if err := config.DB.Preload("Customer").First(&invoice, invoice.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load updated invoice"})
		return
	}

	// 8. Kirim response
	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice updated successfully",
		"data":    invoice,
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
