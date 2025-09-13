package repository

import (
	"gorm.io/gorm"

	"task-api/internal/domain/entity"
	"task-api/internal/repository/model"
	"task-api/internal/usecase/interfaces"
)

// taskRepository implements the TaskRepository interface
type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository creates a new task repository instance
func NewTaskRepository(db *gorm.DB) interfaces.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// Create creates a new task in the database
func (r *taskRepository) Create(taskModel *model.TaskModel) (*entity.Task, error) {
	if err := r.db.Create(taskModel).Error; err != nil {
		return nil, err
	}

	return taskModel.ToEntity(), nil
}

// GetByID retrieves a task by its ID
func (r *taskRepository) GetByID(id uint) (*entity.Task, error) {
	var taskModel model.TaskModel
	
	if err := r.db.First(&taskModel, id).Error; err != nil {
		return nil, err
	}

	return taskModel.ToEntity(), nil
}

// GetAll retrieves all tasks
func (r *taskRepository) GetAll() ([]*entity.Task, error) {
	var taskModels []model.TaskModel
	
	if err := r.db.Find(&taskModels).Error; err != nil {
		return nil, err
	}

	tasks := make([]*entity.Task, len(taskModels))
	for i, taskModel := range taskModels {
		tasks[i] = taskModel.ToEntity()
	}

	return tasks, nil
}

// Update updates an existing task
func (r *taskRepository) Update(taskModel *model.TaskModel) (*entity.Task, error) {
	if err := r.db.Save(taskModel).Error; err != nil {
		return nil, err
	}

	return taskModel.ToEntity(), nil
}

// Delete deletes a task by its ID
func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&model.TaskModel{}, id).Error
}