// controllers/scale_detail_controller.go
package controllers

import (
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
	"scrap-invoice-backend/models/response"
	"time"

	"github.com/gin-gonic/gin"
)

func GetScaleDetail(c *gin.Context) {
	id := c.Param("id")
	var detail models.ScaleDetail

	if err := config.DB.Preload("Item.Category").Preload("Invoice").First(&detail, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Scale detail not found"})
		return
	}

	responseData := response.ScaleDetailResponse{
		ID:         detail.ID,
		Weight:     detail.Weight,
		AlasWeight: detail.AlasWeight,
		Photo:      detail.Photo,
		ScaleType:  detail.ScaleType,
		CreatedAt:  detail.CreatedAt,
		Invoice: response.InvoiceInfo{
			ID:            detail.Invoice.ID,
			InvoiceNumber: detail.Invoice.InvoiceNumber,
		},
		Item: response.ItemInfo{
			ID:       detail.Item.ID,
			Name:     detail.Item.ItemName,
			Category: detail.Item.Category.ItemCategoryName,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData})
}

func CreateScaleDetail(c *gin.Context) {
	var input models.ScaleDetail

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal simpan scale detail"})
		return
	}

	// Load relasi
	var detail models.ScaleDetail
	if err := config.DB.
		Preload("Item.Category").
		Preload("Invoice").
		First(&detail, input.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal load relasi"})
		return
	}

	// Mapping ke DTO
	response := response.ScaleDetailResponse{
		ID:         detail.ID,
		Weight:     detail.Weight,
		AlasWeight: detail.AlasWeight,
		Photo:      detail.Photo,
		ScaleType:  detail.ScaleType,
		CreatedAt:  detail.CreatedAt,
		Invoice: response.InvoiceInfo{
			ID:            detail.Invoice.ID,
			InvoiceNumber: detail.Invoice.InvoiceNumber,
		},
		Item: response.ItemInfo{
			ID:       detail.Item.ID,
			Name:     detail.Item.ItemName,
			Category: detail.Item.Category.ItemCategoryName,
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
