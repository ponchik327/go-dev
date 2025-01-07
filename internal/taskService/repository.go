package taskService

import (
	"project/internal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task models.Task) (models.Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]models.Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task models.Task) (models.Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task models.Task) (models.Task, error) {
	err := r.db.Create(&task).Error

	return task, err
}

func (r *taskRepository) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error

	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, newTask models.Task) (models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return models.Task{}, err
	}

	task.Content = newTask.Content
	task.IsDone = newTask.IsDone

	err = r.db.Save(&task).Error
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var task models.Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return err
	}

	return r.db.Unscoped().Delete(task).Error
}
