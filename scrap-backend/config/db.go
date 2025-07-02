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
	// coba load dulu dari file .env.mock atau sejenis
	_ = godotenv.Load(".env.mock") // default fallback
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "dev"
	}

	// coba load ulang .env yang sesuai mode
	envFile := ".env." + appEnv
	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("⚠️  Gagal load %s, fallback ke .env", envFile)
		_ = godotenv.Load(".env")
	}

	log.Println("ENV:", appEnv)

	if appEnv == "mock" {
		log.Println("✅ Running in MOCK mode: skipping DB connection")
		return
	}

	// connect DB real
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	DB = db
	log.Println("✅ Connected to DB in", appEnv, "mode")
}
