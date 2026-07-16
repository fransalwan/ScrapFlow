package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Try to load .env file (untuk development lokal)
	// Tidak error kalau tidak ada (untuk production di Railway)
	err := godotenv.Load()
	if err != nil {
		log.Println("ℹ️  No .env file found, using system env (production mode)")
	}

	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE") // ✅ FIX: SSL_MODE -> DB_SSLMODE

	// Validation & Logging
	if dbHost == "" {
		log.Println("️  WARNING: DB_HOST is empty, using default localhost")
		dbHost = "localhost"
		dbUser = "postgres"
		dbPassword = "postgres"
		dbName = "scrapflow_db"
		dbPort = "5432"
		dbSSLMode = "disable"
	} else {
		log.Printf("✅ Connecting to production database at: %s", dbHost)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
		dbSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ [error] failed to initialize database, got error %v", err)
	}

	DB = db

	if dbHost == "localhost" {
		log.Println("✅ Connected to LOCAL PostgreSQL database")
	} else {
		log.Println("✅ Connected to PRODUCTION PostgreSQL database (Neon.tech)")
	}
}
