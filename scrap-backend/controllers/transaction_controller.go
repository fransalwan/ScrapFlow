package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fransalwan/scrap-backend/config"
	"github.com/fransalwan/scrap-backend/models"
	"github.com/gin-gonic/gin"
)

// GET /transactions
func GetTransactions(c *gin.Context) {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "mock" {
		// dummy customer
		dummyCustomer := models.Customer{
			ID:      1,
			Name:    "Mock Customer",
			Phone:   "000",
			Address: "Mock City",
		}

		// dummy transaction
		dummyTransaction := models.Transaction{
			ID:         1,
			CustomerID: dummyCustomer.ID,
			Customer:   dummyCustomer,
			Material:   "Besi Tua",
			Weight:     120,
			PricePerKg: 3500,
			Total:      120 * 3500,
			CreatedAt:  time.Now(),
		}

		c.JSON(http.StatusOK, gin.H{"data": []models.Transaction{dummyTransaction}})
		return
	}

	// Real DB logic
	var transactions []models.Transaction
	config.DB.Preload("Customer").Find(&transactions)
	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// POST /transactions
func CreateTransaction(c *gin.Context) {
	appEnv := os.Getenv("APP_ENV")

	var input models.Transaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hitung total
	input.Total = input.Weight * input.PricePerKg

	if appEnv == "mock" {
		dummyCustomer := models.Customer{
			ID:      input.CustomerID,
			Name:    "Mock Customer " + fmt.Sprint(input.CustomerID),
			Phone:   "000",
			Address: "Mock City",
		}

		// Inject dummy customer ke dalam relasi transaction
		input.Customer = dummyCustomer

		c.JSON(http.StatusOK, gin.H{
			"data": input,
		})
		return
	}

	// Real logic DB
	config.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}
