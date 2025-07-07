package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"scrap-invoice-backend/config"
	"scrap-invoice-backend/routes"
)

func main() {
	fmt.Println("Starting Scrap Invoice Backend...")
	log.Println("Setting up CORS middleware...")
	config.ConnectDB()

	r := gin.Default()

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			log.Println("CORS origin check:", origin) // <-- log di sini
			// Allow semua localhost:xxxx (dev tools, vite, dkk)
			return strings.HasPrefix(origin, "http://localhost:")
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
