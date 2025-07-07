// File: controllers/summary_controller.go
package controllers

import (
	"fmt"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
	"sort"

	"github.com/gin-gonic/gin"
)

type SummaryResponse struct {
	InvoiceID   int     `json:"invoice_id"`
	ItemID      int     `json:"item_id"`
	ItemName    string  `json:"item_name"`
	TotalWeight float64 `json:"total_weight"`
	PricePerKg  float64 `json:"price_per_kg"`
	TotalPrice  float64 `json:"total_price"`
}

func GenerateSummary(c *gin.Context) {
	invoiceID := c.Param("invoice_id")

	// 1. Validasi invoice
	var invoice models.Invoice
	if err := config.DB.First(&invoice, invoiceID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	// 2. Ambil semua scale_detail yang belum masuk summary
	var scaleDetails []models.ScaleDetail
	if err := config.DB.
		Preload("Item").
		Where("invoice_id = ? AND summary_id IS NULL", invoiceID).
		Find(&scaleDetails).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch scale details"})
		return
	}

	if len(scaleDetails) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No new scale details to summarize"})
		return
	}

	// 3. Kelompokkan berdasarkan item_id
	summaryMap := make(map[int]*models.Summary)
	itemNameMap := make(map[int]string) // buat sorting by item_name
	for _, detail := range scaleDetails {
		if summaryMap[detail.ItemID] == nil {
			summaryMap[detail.ItemID] = &models.Summary{
				InvoiceID:     detail.InvoiceID,
				ItemID:        detail.ItemID,
				TotalWeight:   0,
				SubTotalPrice: 0,
			}
			itemNameMap[detail.ItemID] = detail.Item.ItemName
		}
		summaryMap[detail.ItemID].TotalWeight += detail.Weight
		summaryMap[detail.ItemID].SubTotalPrice += detail.Weight * detail.Item.PricePerKg
	}

	// 4. Simpan ke DB dan update scale_detail
	var responseList []SummaryResponse
	var totalWeight float64
	var totalPrice float64
	for itemID, summary := range summaryMap {
		if err := config.DB.Create(summary).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save summary"})
			return
		}

		// Update semua scale_detail yang belum tersummary
		config.DB.Model(&models.ScaleDetail{}).
			Where("invoice_id = ? AND item_id = ? AND summary_id IS NULL", summary.InvoiceID, summary.ItemID).
			Update("summary_id", summary.ID)

		// Akumulasi invoice
		totalWeight += summary.TotalWeight
		totalPrice += summary.SubTotalPrice

		responseList = append(responseList, SummaryResponse{
			InvoiceID:   summary.InvoiceID,
			ItemID:      summary.ItemID,
			ItemName:    itemNameMap[itemID],
			TotalWeight: summary.TotalWeight,
			PricePerKg:  summary.SubTotalPrice / summary.TotalWeight,
			TotalPrice:  summary.SubTotalPrice,
		})
	}

	// 5. Update total invoice
	result := config.DB.Model(&models.Invoice{}).
		Where("invoice_id = ?", invoice.ID).
		Updates(map[string]interface{}{
			"total_weight": totalWeight,
			"total_price":  totalPrice,
		})

	fmt.Println("Rows affected:", result.RowsAffected)
	fmt.Println("Error:", result.Error)

	// 6. Urutkan berdasarkan nama item
	sort.Slice(responseList, func(i, j int) bool {
		return responseList[i].ItemName < responseList[j].ItemName
	})

	// 7. Response
	c.JSON(http.StatusOK, gin.H{
		"message": "Summaries created successfully",
		"data":    responseList,
	})
}

// func GenerateSummary(c *gin.Context) {
// 	invoiceID := c.Param("invoice_id")
// 	var scaleDetails []models.ScaleDetail

// 	// 1. Ambil semua scale_detail yang belum masuk summary untuk invoice tertentu
// 	if err := config.DB.
// 		Preload("Item").
// 		Where("invoice_id = ? AND summary_id IS NULL", invoiceID).
// 		Find(&scaleDetails).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch scale details"})
// 		return
// 	}

// 	if len(scaleDetails) == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No new scale details to summarize"})
// 		return
// 	}

// 	// 2. Kelompokkan berdasarkan item_id
// 	summaryMap := make(map[int]*models.Summary)
// 	for _, detail := range scaleDetails {
// 		if summaryMap[detail.ItemID] == nil {
// 			summaryMap[detail.ItemID] = &models.Summary{
// 				InvoiceID:     detail.InvoiceID,
// 				ItemID:        detail.ItemID,
// 				TotalWeight:   0,
// 				SubTotalPrice: 0,
// 			}
// 		}
// 		summaryMap[detail.ItemID].TotalWeight += detail.Weight
// 		summaryMap[detail.ItemID].SubTotalPrice += detail.Weight * detail.Item.PricePerKg
// 	}

// 	// 3. Simpan ke DB dan update scale_detail
// 	var responseList []SummaryResponse
// 	for _, summary := range summaryMap {
// 		if err := config.DB.Create(summary).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save summary"})
// 			return
// 		}
// 		// Update semua scale_detail yang belum tersummary
// 		config.DB.Model(&models.ScaleDetail{}).
// 			Where("invoice_id = ? AND item_id = ? AND summary_id IS NULL", summary.InvoiceID, summary.ItemID).
// 			Update("summary_id", summary.ID)

// 		responseList = append(responseList, SummaryResponse{
// 			InvoiceID:   summary.InvoiceID,
// 			ItemID:      summary.ItemID,
// 			TotalWeight: summary.TotalWeight,
// 			PricePerKg:  scaleDetails[0].Item.PricePerKg, // ambil salah satu aja
// 			TotalPrice:  summary.SubTotalPrice,
// 		})
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Summaries created successfully",
// 		"data":    responseList,
// 	})
// }
