package main

import (
	"os"

	"github.com/fransalwan/scrap-backend/config"
	"github.com/fransalwan/scrap-backend/models"
	"github.com/fransalwan/scrap-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDB()

	// skip AutoMigrate kalau mode mock
	if os.Getenv("APP_ENV") != "mock" {
		config.DB.AutoMigrate(&models.Customer{}, &models.Transaction{})
	}

	routes.SetupRoutes(r)
	r.Run(":8080")
}
