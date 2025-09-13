package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"task-api/internal/handler"
	"task-api/internal/infrastructure"
	"task-api/internal/repository"
	"task-api/internal/usecase"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize database
	db, err := infrastructure.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repository
	taskRepo := repository.NewTaskRepository(db)

	// Initialize usecase
	taskUsecase := usecase.NewTaskUsecase(taskRepo)

	// Initialize Echo
	e := echo.New()

	// Setup routes
	handler.SetupRoutes(e, taskUsecase)

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Fatal(e.Start(":" + port))
}