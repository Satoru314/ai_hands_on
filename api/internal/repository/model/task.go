package model

import (
	"time"

	"task-api/internal/domain/entity"
	"task-api/internal/domain/valueobject"
)

// TaskModel represents the database model for tasks with GORM tags
type TaskModel struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	Status    string     `gorm:"type:varchar(20);default:'pending'" json:"status"`
	DueDate   *time.Time `gorm:"type:timestamp;null" json:"due_date,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName specifies the table name for GORM
func (TaskModel) TableName() string {
	return "tasks"
}

// ToEntity converts TaskModel to entity.Task
func (t *TaskModel) ToEntity() *entity.Task {
	return &entity.Task{
		ID:        t.ID,
		Content:   t.Content,
		Status:    valueobject.TaskStatus(t.Status),
		DueDate:   t.DueDate,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

