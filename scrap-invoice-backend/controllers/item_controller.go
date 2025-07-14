package controllers

import (
	"net/http"
	"time"

	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item

	// Preload Category biar relasi ikut diisi
	if err := config.DB.Preload("Category").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch items: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Items fetched successfully",
		"data":    items,
	})
}

func CreateItem(c *gin.Context) {
	var item models.Item

	// Bind JSON
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi
	if item.ItemName == "" || item.ItemCategoryID == 0 || item.PricePerKg == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, price per kg and category are required"})
		return
	}

	// Simpan item
	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item: " + err.Error()})
		return
	}

	// Fetch lagi dengan preload category
	var createdItem models.Item
	if err := config.DB.Preload("Category").First(&createdItem, item.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item with category: " + err.Error()})
		return
	}

	// Sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Item created successfully",
		"data":    createdItem,
	})
}

func UpdateItem(c *gin.Context) {
	var input models.Item
	id := c.Param("id")

	// Bind JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi
	if input.ItemName == "" || input.ItemCategoryID == 0 || input.PricePerKg == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name, price per kg and category are required"})
		return
	}

	// Fetch item lama dulu
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	// Update field yang diizinkan
	item.ItemName = input.ItemName
	item.ItemCategoryID = input.ItemCategoryID
	item.PricePerKg = input.PricePerKg
	item.UpdatedAt = time.Now()

	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item: " + err.Error()})
		return
	}

	// Fetch ulang dengan preload category
	var updatedItem models.Item
	if err := config.DB.Preload("Category").First(&updatedItem, item.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item updated successfully",
		"data":    updatedItem,
	})
}

func DeleteItem(c *gin.Context) {
	id := c.Param("id")

	// Hapus data di DB
	if err := config.DB.Where("item_id = ?", id).Delete(&models.Item{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Item deleted successfully",
	})
}
