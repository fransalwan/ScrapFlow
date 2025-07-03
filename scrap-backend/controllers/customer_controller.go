package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fransalwan/scrap-backend/config"
	"github.com/fransalwan/scrap-backend/models"
	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	if os.Getenv("APP_ENV") == "mock" {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>Running in mock mode")
		// return dummy data
		c.JSON(http.StatusOK, gin.H{"data": []models.Customer{
			{Name: "Mock Customer", Phone: "000", Address: "Mock City"},
		}})
		return
	}

	// real logic
	var customers []models.Customer
	config.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func CreateCustomer(c *gin.Context) {
	var input models.Customer
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
