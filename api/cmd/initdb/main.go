package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"task-api/internal/repository/model"
)

func main() {
	// Load environment variables
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get database URL
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	fmt.Printf("Connecting to database: %s\n", dbURL)

	// Connect to database
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Database connection successful!")

	// Auto migrate tables
	err = db.AutoMigrate(&model.TaskModel{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	fmt.Println("✅ Database tables created successfully!")

	// Verify table creation
	var count int64
	db.Model(&model.TaskModel{}).Count(&count)
	fmt.Printf("✅ Current task count: %d\n", count)

	fmt.Println("🎉 Database initialization completed successfully!")
}