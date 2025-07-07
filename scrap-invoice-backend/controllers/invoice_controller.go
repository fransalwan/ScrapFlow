package controllers

import (
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

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
	var invoice models.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Cek apakah invoice_number sudah dipakai
	var existing models.Invoice
	if err := config.DB.Where("invoice_number = ?", invoice.InvoiceNumber).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invoice number already exists"})
		return
	}

	// Set default values
	if invoice.Status == "" {
		invoice.Status = "draft"
	}

	invoice.CreatedAt = time.Now()

	// Simpan ke DB
	if err := config.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice created successfully",
		"data":    invoice,
	})
}

func UpdateInvoice(c *gin.Context) {
	id := c.Param("id")
	var invoice models.Invoice

	// Cari invoice berdasarkan ID
	if err := config.DB.First(&invoice, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	// Bind data baru ke struct invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Update ke DB
	if err := config.DB.Save(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice: " + err.Error()})
		return
	}

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
