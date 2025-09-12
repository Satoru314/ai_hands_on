package entity

import (
	"time"

	"task-api/internal/domain/valueobject"
)

// Task represents a task entity
type Task struct {
	ID        uint
	Content   string
	Status    valueobject.TaskStatus
	DueDate   *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}