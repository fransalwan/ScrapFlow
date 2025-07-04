package controllers

import (
	"fmt"
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

	// validasi unique email
	var existingCustomer models.Customer
	if err := config.DB.Where("email = ?", c.PostForm("email")).First(&existingCustomer).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Bind JSON dari body ke struct
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi minimal name dan email
	if customer.Name == "" || customer.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required"})
		return
	}

	// Set waktu otomatis (kalau belum di-handle di DB)
	now := time.Now()
	customer.CreatedAt = now
	customer.UpdatedAt = now

	// Simpan ke DB
	if err := config.DB.Create(&customer).Error; err != nil {
		// debugging isi customer
		fmt.Println("Customer data:", customer)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer: " + err.Error()})
		return
	}

	// Response sukses
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer created successfully",
		"data":    customer,
	})
}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	// Bind JSON dari body ke struct
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	// Validasi minimal name dan email
	if customer.Name == "" || customer.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name and Email are required"})
		return
	}

	// Update data di DB
	if err := config.DB.Model(&models.Customer{}).Where("id_customer = ?", id).Updates(customer).Error; err != nil {
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
	if err := config.DB.Where("id_customer = ?", id).Delete(&models.Customer{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Customer deleted successfully",
	})
}
