package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global variable to hold db connection instance and pointer to gorm.DB
var DB *gorm.DB

func InitDB() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	// get all required variable from env
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	//connection establish to db with default config
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect db: %v", err)
	}

	DB = db
}
