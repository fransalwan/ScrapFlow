package controllers

import (
	"log"
	"net/http"
	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"

	"github.com/gin-gonic/gin"
)

func GetInvoiceCount(c *gin.Context) {
	log.Println("Hit GetInvoiceCount")

	if config.DB == nil {
		log.Println("DB is nil!!")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB not initialized"})
		return
	}

	var count int64
	if err := config.DB.Model(&models.Invoice{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal ambil data invoice"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}

func GetCustomerCount(c *gin.Context) {
	var count int64

	if err := config.DB.Model(&models.Customer{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal ambil jumlah customer",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Jumlah customer berhasil diambil",
		"count":   count,
	})
}
