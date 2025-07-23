// File: controllers/summary_controller.go
package controllers

import (
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"

	"github.com/gin-gonic/gin"
)

func GetScaleSummaryByInvoice(c *gin.Context) {
	invoiceId := c.Param("id")
	var scaleDetails []models.ScaleDetail
	var grandTotal float64 // FIXED: ubah dari int ke float64

	// Step 1: Ambil semua scale_detail dengan relasi item
	if err := config.DB.
		Preload("Item").
		Where("invoice_id = ?", invoiceId).
		Find(&scaleDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get scale details"})
		return
	}

	type Summary struct {
		ItemID        int     `json:"item_id"`
		ItemName      string  `json:"item_name"`
		PricePerKg    float64 `json:"price_per_kg"`
		TotalWeight   float64 `json:"total_weight"`
		SubTotalPrice float64 `json:"sub_total_price"`
	}

	summaryMap := make(map[int]*Summary)

	for _, sd := range scaleDetails {
		item := sd.Item

		if _, ok := summaryMap[item.ID]; !ok {
			summaryMap[item.ID] = &Summary{
				ItemID:     item.ID,
				ItemName:   item.ItemName,
				PricePerKg: item.PricePerKg,
			}
		}

		summary := summaryMap[item.ID]
		summary.TotalWeight += sd.Weight
		summary.SubTotalPrice = summary.TotalWeight * summary.PricePerKg
	}

	// Convert map ke slice
	summaries := make([]Summary, 0, len(summaryMap))
	for _, s := range summaryMap {
		summaries = append(summaries, *s)
		grandTotal += s.SubTotalPrice // FIXED: Sekarang ga error
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       summaries,
		"grandTotal": grandTotal, // bisa lo aktifin sekarang
	})
}
