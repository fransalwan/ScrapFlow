// controllers/scale_detail_controller.go
package controllers

import (
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/dto"
	"scrap-invoice-backend/models"
	"scrap-invoice-backend/models/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetScaleDetailsByInvoiceID(c *gin.Context) {
	invoiceID := c.Param("id") // pakai nama yang sama kayak di route
	var details []models.ScaleDetail

	if err := config.DB.
		Preload("Item.Category").
		Preload("Invoice").
		Where("invoice_id = ?", invoiceID).
		Find(&details).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scale details not found"})
		return
	}

	var responseData []response.ScaleDetailResponse
	for _, d := range details {
		responseData = append(responseData, response.ScaleDetailResponse{
			ID:         d.ID,
			Weight:     d.Weight,
			AlasWeight: d.AlasWeight,
			Photo:      d.Photo,
			ScaleType:  d.ScaleType,
			CreatedAt:  d.CreatedAt,
			Invoice: response.InvoiceInfo{
				ID:            d.Invoice.ID,
				InvoiceNumber: d.Invoice.InvoiceNumber,
			},
			Item: response.ItemInfo{
				ID:       d.Item.ID,
				Name:     d.Item.ItemName,
				Category: d.Item.Category.ItemCategoryName,
			},
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

func CreateScaleDetail(c *gin.Context) {
	// Convert :id
	invoiceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	// Bind ke DTO
	var input dto.CreateScaleDetailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map ke model
	scaleDetail := models.ScaleDetail{
		InvoiceID:  invoiceID,
		ItemID:     input.ItemID,
		Weight:     input.Weight,
		AlasWeight: input.AlasWeight,
		Photo:      input.Photo,
		ScaleType:  input.ScaleType,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Simpan
	if err := config.DB.Create(&scaleDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan scale detail"})
		return
	}

	// Load relasi
	if err := config.DB.
		Preload("Item.Category").
		Preload("Invoice").
		First(&scaleDetail, scaleDetail.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal load relasi"})
		return
	}

	// Mapping ke response DTO
	response := dto.ScaleDetailResponse{
		ID:         scaleDetail.ID,
		Weight:     scaleDetail.Weight,
		AlasWeight: scaleDetail.AlasWeight,
		Photo:      scaleDetail.Photo,
		ScaleType:  scaleDetail.ScaleType,
		CreatedAt:  scaleDetail.CreatedAt,
		Invoice: dto.InvoiceInfo{
			ID:            scaleDetail.Invoice.ID,
			InvoiceNumber: scaleDetail.Invoice.InvoiceNumber,
		},
		Item: dto.ItemInfo{
			ID:       scaleDetail.Item.ID,
			Name:     scaleDetail.Item.ItemName,
			Category: scaleDetail.Item.Category.ItemCategoryName,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Scale detail created",
		"data":    response,
	})
}

func UpdateScaleDetail(c *gin.Context) {
	id := c.Param("id")
	var existing models.ScaleDetail

	// Cari scale_detail berdasarkan ID
	if err := config.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scale detail not found"})
		return
	}

	// Bind JSON
	var updated models.ScaleDetail
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Update field yang boleh diubah
	existing.Weight = updated.Weight
	existing.AlasWeight = updated.AlasWeight
	existing.Photo = updated.Photo
	existing.ScaleType = updated.ScaleType
	existing.ItemID = updated.ItemID
	existing.UpdatedAt = time.Now()

	if err := config.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update scale detail"})
		return
	}

	// 🔁 Preload ulang untuk keperluan response
	var fullDetail models.ScaleDetail
	if err := config.DB.Preload("Item.Category").Preload("Invoice").First(&fullDetail, existing.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated detail"})
		return
	}

	// 📦 Map ke response DTO
	response := response.ScaleDetailResponse{
		ID:         fullDetail.ID,
		Weight:     fullDetail.Weight,
		AlasWeight: fullDetail.AlasWeight,
		Photo:      fullDetail.Photo,
		ScaleType:  fullDetail.ScaleType,
		CreatedAt:  fullDetail.CreatedAt,
		Invoice: response.InvoiceInfo{
			ID:            fullDetail.Invoice.ID,
			InvoiceNumber: fullDetail.Invoice.InvoiceNumber,
		},
		Item: response.ItemInfo{
			ID:       fullDetail.Item.ID,
			Name:     fullDetail.Item.ItemName,
			Category: fullDetail.Item.Category.ItemCategoryName,
		},
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scale detail updated successfully", "data": response})
}

func DeleteScaleDetail(c *gin.Context) {
	id := c.Param("id")
	var detail models.ScaleDetail

	if err := config.DB.First(&detail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scale detail not found"})
		return
	}

	if err := config.DB.Delete(&detail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete scale detail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scale detail deleted successfully"})
}
