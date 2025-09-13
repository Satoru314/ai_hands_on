package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"task-api/internal/usecase"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(e *echo.Echo, taskUsecase *usecase.TaskUsecase) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Create handlers
	taskHandler := NewTaskHandler(taskUsecase)

	// API routes
	api := e.Group("/api")

	// Task routes
	api.POST("/tasks", taskHandler.CreateTask)
	api.GET("/tasks/:id", taskHandler.GetTask)
	api.GET("/tasks", taskHandler.GetAllTasks)
	api.PUT("/tasks/:id", taskHandler.UpdateTask)
	api.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
