package dto

import (
	"time"

	"task-api/internal/domain/entity"
)

// CreateTaskRequest represents the request for creating a task
type CreateTaskRequest struct {
	Content string     `json:"content" validate:"required"`
	Status  string     `json:"status,omitempty"`
	DueDate *time.Time `json:"due_date,omitempty"`
}

// UpdateTaskRequest represents the request for updating a task
type UpdateTaskRequest struct {
	Content string     `json:"content" validate:"required"`
	Status  string     `json:"status" validate:"required"`
	DueDate *time.Time `json:"due_date,omitempty"`
}

// TaskResponse represents the response for task operations
type TaskResponse struct {
	ID        uint       `json:"id"`
	Content   string     `json:"content"`
	Status    string     `json:"status"`
	DueDate   *time.Time `json:"due_date,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// TaskListResponse represents the response for getting all tasks
type TaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// ToResponse converts entity.Task to TaskResponse
func ToTaskResponse(task *entity.Task) TaskResponse {
	return TaskResponse{
		ID:        task.ID,
		Content:   task.Content,
		Status:    task.Status.String(),
		DueDate:   task.DueDate,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
}

// ToTaskListResponse converts slice of entity.Task to TaskListResponse
func ToTaskListResponse(tasks []*entity.Task) TaskListResponse {
	taskResponses := make([]TaskResponse, len(tasks))
	for i, task := range tasks {
		taskResponses[i] = ToTaskResponse(task)
	}
	return TaskListResponse{Tasks: taskResponses}
}