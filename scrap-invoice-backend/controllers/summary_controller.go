// File: controllers/summary_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"

	"github.com/gin-gonic/gin"
)

func GetScaleSummaryByInvoice(c *gin.Context) {
	invoiceId := c.Param("id")

	// ✅ TAMBAHKAN LOG INI
	fmt.Printf(" Getting summary for invoice ID: %s\n", invoiceId)

	var scaleDetails []models.ScaleDetail

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

	// Nested map[scale_type][item_id]*Summary
	summaryMap := make(map[string]map[int]*Summary)
	grandTotalMap := make(map[string]float64)

	for _, sd := range scaleDetails {
		scaleType := sd.ScaleType
		item := sd.Item

		if _, ok := summaryMap[scaleType]; !ok {
			summaryMap[scaleType] = make(map[int]*Summary)
		}

		if _, ok := summaryMap[scaleType][item.ID]; !ok {
			summaryMap[scaleType][item.ID] = &Summary{
				ItemID:     item.ID,
				ItemName:   item.ItemName,
				PricePerKg: item.PricePerKg,
			}
		}

		summary := summaryMap[scaleType][item.ID]
		summary.TotalWeight += sd.Weight
		summary.SubTotalPrice = summary.TotalWeight * summary.PricePerKg
		grandTotalMap[scaleType] += sd.Weight * summary.PricePerKg
	}

	// Convert to JSON serializable format
	finalData := make(map[string][]Summary)
	for scaleType, items := range summaryMap {
		for _, summary := range items {
			finalData[scaleType] = append(finalData[scaleType], *summary)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       finalData,
		"grandTotal": grandTotalMap,
	})
}
