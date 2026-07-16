package main

import (
	"flag"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
	"scrap-invoice-backend/routes"
	"scrap-invoice-backend/seeders"
)

func main() {
	config.ConnectDB()

	// Buat flag --reset
	resetDB := flag.Bool("reset", false, "Reset database (drop all tables)")
	flag.Parse()

	// Kalau ada flag --reset, drop semua tabel dengan urutan yang bener
	if *resetDB {
		log.Println("🔄 Resetting database...")

		// Drop child tables dulu (yang punya foreign key)
		config.DB.Migrator().DropTable("scale_details")
		config.DB.Migrator().DropTable("summaries")
		config.DB.Migrator().DropTable("invoices")
		config.DB.Migrator().DropTable("items")
		config.DB.Migrator().DropTable("item_categories")
		config.DB.Migrator().DropTable("customers")
		config.DB.Migrator().DropTable("users")

		log.Println("✅ All tables dropped")
	}

	// Migrate
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Customer{},
		&models.ItemCategory{},
		&models.Item{},
		&models.Invoice{},
		&models.Summary{},
		&models.ScaleDetail{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	log.Println("✅ Database migrated successfully")

	// Seed
	seeders.SeedData()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(func(c *gin.Context) {
		log.Printf("REQ: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		log.Println("RES HEADERS:", c.Writer.Header())
	})

	routes.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
