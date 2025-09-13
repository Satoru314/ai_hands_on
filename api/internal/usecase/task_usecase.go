package usecase

import (
	"errors"
	"time"

	"task-api/internal/domain/entity"
	"task-api/internal/domain/valueobject"
	"task-api/internal/repository/model"
	"task-api/internal/usecase/interfaces"
)

// TaskUsecase handles business logic for tasks
type TaskUsecase struct {
	taskRepo interfaces.TaskRepository
}

// NewTaskUsecase creates a new task usecase instance
func NewTaskUsecase(taskRepo interfaces.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		taskRepo: taskRepo,
	}
}

// CreateTask creates a new task with validation
func (u *TaskUsecase) CreateTask(content string, dueDate *time.Time) (*entity.Task, error) {
	// Validation
	if content == "" {
		return nil, errors.New("task content cannot be empty")
	}

	// Create model in usecase layer with timestamps
	now := time.Now()
	taskModel := &model.TaskModel{
		Content:   content,
		Status:    valueobject.StatusPending.String(),
		DueDate:   dueDate,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Save to repository
	return u.taskRepo.Create(taskModel)
}

// GetTask retrieves a task by ID
func (u *TaskUsecase) GetTask(id uint) (*entity.Task, error) {
	if id == 0 {
		return nil, errors.New("invalid task ID")
	}

	return u.taskRepo.GetByID(id)
}

// GetAllTasks retrieves all tasks
func (u *TaskUsecase) GetAllTasks() ([]*entity.Task, error) {
	return u.taskRepo.GetAll()
}

// UpdateTask updates an existing task
func (u *TaskUsecase) UpdateTask(id uint, content string, status valueobject.TaskStatus, dueDate *time.Time) (*entity.Task, error) {
	// Validation
	if id == 0 {
		return nil, errors.New("invalid task ID")
	}
	if content == "" {
		return nil, errors.New("task content cannot be empty")
	}
	if !status.IsValid() {
		return nil, errors.New("invalid task status")
	}

	// Check if task exists
	_, err := u.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create model in usecase layer with updated timestamp
	taskModel := &model.TaskModel{
		ID:        id,
		Content:   content,
		Status:    status.String(),
		DueDate:   dueDate,
		UpdatedAt: time.Now(),
	}

	// Update via repository
	return u.taskRepo.Update(taskModel)
}

// UpdateTaskStatus updates only the status of a task
func (u *TaskUsecase) UpdateTaskStatus(id uint, status valueobject.TaskStatus) (*entity.Task, error) {
	// Validation
	if id == 0 {
		return nil, errors.New("invalid task ID")
	}
	if !status.IsValid() {
		return nil, errors.New("invalid task status")
	}

	// Get existing task
	task, err := u.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Create model with updated status and timestamp
	taskModel := &model.TaskModel{
		ID:        id,
		Content:   task.Content,
		Status:    status.String(),
		DueDate:   task.DueDate,
		UpdatedAt: time.Now(),
	}

	// Update via repository
	return u.taskRepo.Update(taskModel)
}

// DeleteTask deletes a task by ID
func (u *TaskUsecase) DeleteTask(id uint) error {
	if id == 0 {
		return errors.New("invalid task ID")
	}

	// Check if task exists
	_, err := u.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Delete task
	return u.taskRepo.Delete(id)
}