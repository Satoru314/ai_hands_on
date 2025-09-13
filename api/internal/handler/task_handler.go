package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"task-api/internal/domain/valueobject"
	"task-api/internal/handler/dto"
	"task-api/internal/usecase"
)

// TaskHandler handles HTTP requests for task operations
type TaskHandler struct {
	taskUsecase *usecase.TaskUsecase
}

// NewTaskHandler creates a new task handler instance
func NewTaskHandler(taskUsecase *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		taskUsecase: taskUsecase,
	}
}

// CreateTask handles POST /tasks
func (h *TaskHandler) CreateTask(c echo.Context) error {
	var req dto.CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request format",
		})
	}

	// Validation
	if req.Content == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: "Content is required",
		})
	}

	// Create task
	task, err := h.taskUsecase.CreateTask(req.Content, req.DueDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "creation_failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, dto.ToTaskResponse(task))
}

// GetTask handles GET /tasks/:id
func (h *TaskHandler) GetTask(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid task ID",
		})
	}

	task, err := h.taskUsecase.GetTask(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Error:   "task_not_found",
			Message: "Task not found",
		})
	}

	return c.JSON(http.StatusOK, dto.ToTaskResponse(task))
}

// GetAllTasks handles GET /tasks
func (h *TaskHandler) GetAllTasks(c echo.Context) error {
	tasks, err := h.taskUsecase.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "fetch_failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ToTaskListResponse(tasks))
}

// UpdateTask handles PUT /tasks/:id
func (h *TaskHandler) UpdateTask(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid task ID",
		})
	}

	var req dto.UpdateTaskRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_request",
			Message: "Invalid request format",
		})
	}

	// Validation
	if req.Content == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: "Content is required",
		})
	}

	// Validate status
	status := valueobject.TaskStatus(req.Status)
	if !status.IsValid() {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "validation_error",
			Message: "Invalid status",
		})
	}

	// Update task
	task, err := h.taskUsecase.UpdateTask(uint(id), req.Content, status, req.DueDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "update_failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ToTaskResponse(task))
}

// DeleteTask handles DELETE /tasks/:id
func (h *TaskHandler) DeleteTask(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error:   "invalid_id",
			Message: "Invalid task ID",
		})
	}

	err = h.taskUsecase.DeleteTask(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error:   "deletion_failed",
			Message: err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}