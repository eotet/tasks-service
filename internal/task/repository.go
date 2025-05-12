package task

import (
	errs "github.com/eotet/tasks-service/internal/errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateTask(task Task) (Task, error)
	GetTaskByID(id uint32) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint32, task UpdateTaskRequest) (Task, error)
	DeleteTaskByID(id uint32) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateTask(task Task) (Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *repository) GetTaskByID(id uint32) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, errs.ErrTaskNotFound
	}
	return task, nil
}

func (r *repository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *repository) UpdateTaskByID(id uint32, task UpdateTaskRequest) (Task, error) {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Task{}, errs.ErrTaskNotFound
	}

	if task.Text != nil {
		existingTask.Text = *task.Text
	}
	if task.IsDone != nil {
		existingTask.IsDone = *task.IsDone
	}

	if err := r.db.Save(&existingTask).Error; err != nil {
		return Task{}, err
	}

	return existingTask, nil
}

func (r *repository) DeleteTaskByID(id uint32) error {
	var existingTask Task
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return errs.ErrTaskNotFound
	}

	if err := r.db.Delete(&existingTask, id).Error; err != nil {
		return err
	}

	return nil
}
