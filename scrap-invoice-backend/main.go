package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"scrap-invoice-backend/config"
	"scrap-invoice-backend/models"
	"scrap-invoice-backend/routes"
	"scrap-invoice-backend/seeders"
)

func main() {
	// 1. Set Gin Mode ke Release jika di production (mengurangi log spam & meningkatkan performa)
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	config.ConnectDB()

	// 2. Flag --reset (Aman untuk local, diblokir di production)
	resetDB := flag.Bool("reset", false, "Reset database (drop all tables)")
	flag.Parse()

	if *resetDB {
		if os.Getenv("APP_ENV") == "production" {
			log.Fatal("🚫 ERROR: Resetting database is strictly disabled in production!")
		}
		log.Println("🔄 Resetting database...")
		// Drop child tables dulu (yang punya foreign key)
		config.DB.Migrator().DropTable("scale_details", "summaries", "invoices", "items", "item_categories", "customers", "users")
		log.Println("✅ All tables dropped")
	}

	// 3. Migrate
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

	// 4. Seed Data (HANYA jalankan di development untuk mencegah data production tertimpa)
	if os.Getenv("APP_ENV") == "development" {
		seeders.SeedData()
	}

	r := gin.Default()

	// 5. Dynamic CORS (Mendukung multiple domain dipisah koma)
	corsOrigin := os.Getenv("CORS_ORIGIN")
	if corsOrigin == "" {
		corsOrigin = "*" // Fallback aman untuk local development
	}
	allowedOrigins := strings.Split(corsOrigin, ",")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 6. Request Logger Middleware
	r.Use(func(c *gin.Context) {
		log.Printf("REQ: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
		// log.Println("RES HEADERS:", c.Writer.Header()) // Opsional: bisa di-comment di prod agar log lebih bersih
	})

	routes.RegisterRoutes(r)

	// 7. Dynamic Port untuk PaaS (Render, Railway, dll)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback untuk local
	}

	// 8. Setup HTTP Server untuk Graceful Shutdown
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Jalankan server di goroutine terpisah
	go func() {
		log.Printf("🚀 Server is running on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// 9. Graceful Shutdown (Menunggu sinyal dari PaaS saat redeploy)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("✅ Server exiting cleanly")
}
