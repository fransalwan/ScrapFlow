package main

import (
	"log"

	"scrap-invoice-backend/config"
	"scrap-invoice-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()
	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
