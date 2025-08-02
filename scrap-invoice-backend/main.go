package main

import (
	"fmt"
	"log"
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

	r := gin.New() // Ganti dari gin.Default() biar middleware gak ketiban
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // ⛔ Jangan pakai "*" di production
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Set true kalau pakai cookie / token auth
		AllowOriginFunc: func(origin string) bool {
			// Tambahin logic whitelist kalau mau custom (misalnya banyak domain frontend)
			// return origin == "https://frontend.example.com"
			return true // ⚠️ sementara buat dev, true
		},
		MaxAge: 12 * time.Hour,
	}))

	// Logging buat debugging route
	r.Use(func(c *gin.Context) {
		log.Printf("REQ: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		log.Println("RES HEADERS:", c.Writer.Header())
	})

	// Handler preflight OPTIONS
	r.OPTIONS("/*path", func(c *gin.Context) {
		c.Status(204)
	})

	// Register route
	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
