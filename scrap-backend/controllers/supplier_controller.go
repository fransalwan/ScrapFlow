package controllers

import (
	"net/http"
	"os"

	"github.com/fransalwan/scrap-backend/config"
	"github.com/fransalwan/scrap-backend/models"
	"github.com/gin-gonic/gin"
)

func GetSuppliers(c *gin.Context) {
	if os.Getenv("APP_ENV") == "mock" {
		// return dummy data
		c.JSON(http.StatusOK, gin.H{"data": []models.Supplier{
			{Name: "Mock Supplier", Phone: "000", Address: "Mock City"},
		}})
		return
	}

	// real logic
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)
	c.JSON(http.StatusOK, gin.H{"data": suppliers})
}

func CreateSupplier(c *gin.Context) {
	var input models.Supplier
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if os.Getenv("APP_ENV") == "mock" {
		// pretend it's saved
		input.ID = 1
		c.JSON(http.StatusOK, gin.H{"data": input})
		return
	}

	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}
