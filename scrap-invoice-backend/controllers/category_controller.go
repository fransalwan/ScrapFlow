package controllers

import (
	"net/http"
	"time"

	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var itemCategory []models.ItemCategory

	// Ambil data dari database
	if err := config.DB.Find(&itemCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers: " + err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Categories fetched successfully",
		"data":    itemCategory,
	})
}

func CreateCategory(c *gin.Context) {
	var itemCategory models.ItemCategory

	// Bind JSON dari body ke struct
	if err := c.ShouldBindJSON(&itemCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi minimal name
	if itemCategory.ItemCategoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	// Set waktu otomatis (kalau belum di-handle di DB)
	now := time.Now()
	itemCategory.CreatedAt = now
	itemCategory.UpdatedAt = now

	// Simpan ke DB
	if err := config.DB.Create(&itemCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create itemCategory: " + err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Item Category created successfully",
		"data":    itemCategory,
	})
}

func UpdateCategory(c *gin.Context) {
	var itemCategory models.ItemCategory
	id := c.Param("id")

	// Bind JSON dari body ke struct
	if err := c.ShouldBindJSON(&itemCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi minimal name dan email
	if itemCategory.ItemCategoryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	// Update data di DB
	if err := config.DB.Model(&models.ItemCategory{}).Where("item_category_id = ?", id).Updates(itemCategory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item category: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item Category updated successfully",
	})
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	// Hapus data di DB
	if err := config.DB.Where("item_category_id = ?", id).Delete(&models.ItemCategory{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item category: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item Category deleted successfully",
	})
}
