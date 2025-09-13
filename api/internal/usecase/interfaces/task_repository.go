package interfaces

import (
	"task-api/internal/domain/entity"
	"task-api/internal/repository/model"
)

// TaskRepository defines the interface for task data access
type TaskRepository interface {
	Create(taskModel *model.TaskModel) (*entity.Task, error)
	GetByID(id uint) (*entity.Task, error)
	GetAll() ([]*entity.Task, error)
	Update(taskModel *model.TaskModel) (*entity.Task, error)
	Delete(id uint) error
}