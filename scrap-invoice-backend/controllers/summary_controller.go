// File: controllers/summary_controller.go
package controllers

import (
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
	"scrap-invoice-backend/services"

	"github.com/gin-gonic/gin"
)

// GetScaleSummaryByInvoice: perhitungan dipindah ke services.SummarizeScaleDetails
// supaya UI dan AI agent memakai rumus yang sama. Bentuk JSON tidak berubah
// (data + grandTotal); hanya ada tambahan total_alas dan net_weight per item.
func GetScaleSummaryByInvoice(c *gin.Context) {
	invoiceID := c.Param("id")
	var scaleDetails []models.ScaleDetail

	if err := config.DB.
		Preload("Item").
		Where("invoice_id = ?", invoiceID).
		Find(&scaleDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get scale details"})
		return
	}

	c.JSON(http.StatusOK, services.SummarizeScaleDetails(scaleDetails))
}
