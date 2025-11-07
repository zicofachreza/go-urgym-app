package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	env := os.Getenv("NODE_ENV")
	if env == "" {
		env = "development"
	}

	// ✅ Load .env file hanya jika environment = development
	if env == "development" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("⚠️ Warning: .env file not found, using system environment variables")
		}
	} else {
		log.Printf("🚀 Running in %s mode, skipping .env load", env)
	}

	// Ambil nilai dari environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// Format DSN
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbName, port, sslmode,
	)

	// Koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database (%s): %v", env, err)
	}

	log.Printf("✅ Database connected successfully (%s mode)", env)
	return db
}
