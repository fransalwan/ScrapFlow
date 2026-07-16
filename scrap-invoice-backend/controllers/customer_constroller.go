package controllers

import (
	"net/http"
	"time"

	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"

	"github.com/gin-gonic/gin"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customer

	// Ambil data dari database
	if err := config.DB.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers: " + err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Customers fetched successfully",
		"data":    customers,
	})
}

func CreateCustomer(c *gin.Context) {
	var customer models.Customer

	// Bind JSON dari body ke struct
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi field kosong
	if customer.Name == "" || customer.Email == "" || customer.Phone == "" || customer.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required: name, email, phone, address"})
		return
	}

	// Validasi tier
	validTiers := map[string]bool{
		"silver":   true,
		"gold":     true,
		"platinum": true,
	}
	if !validTiers[customer.Tier] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tier. Must be one of: silver, gold, platinum"})
		return
	}

	// Cek email unik
	var existing models.Customer
	if err := config.DB.Where("email = ?", customer.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Set waktu
	now := time.Now()
	customer.CreatedAt = now
	customer.UpdatedAt = now

	// Simpan ke DB
	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer created successfully",
		"data":    customer,
	})
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	// Bind JSON dari body
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi field
	if customer.Name == "" || customer.Email == "" || customer.Phone == "" || customer.Address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required: name, email, phone, address"})
		return
	}

	// Validasi tier
	validTiers := map[string]bool{
		"silver":   true,
		"gold":     true,
		"platinum": true,
	}
	if !validTiers[customer.Tier] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tier. Must be one of: silver, gold, platinum"})
		return
	}

	// Optional: validasi email unik hanya jika diubah
	var existing models.Customer
	if err := config.DB.Where("email = ? AND customer_id <> ?", customer.Email, id).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists for another customer"})
		return
	}

	// Update
	customer.UpdatedAt = time.Now()
	if err := config.DB.Model(&models.Customer{}).Where("customer_id = ?", id).Updates(customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer updated successfully",
	})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	// Hapus data di DB
	if err := config.DB.Where("customer_id = ?", id).Delete(&models.Customer{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted successfully",
	})
}
