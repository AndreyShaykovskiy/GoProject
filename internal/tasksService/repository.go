package tasksService

import (
	"fmt"
	"gorm.io/gorm"
)

type TaskRepository interface {
	// CreateTask - Передаем в функцию task типа Task из orm.go
	// возвращаем созданный Task и ошибку
	CreateTask(task Task) (Task, error)
	// GetAllTasks - Возвращаем массив из всех задач в БД и ошибку
	GetAllTasks() ([]Task, error)
	// UpdateTaskByID - Передаем id и Task, возвращаем обновленный Task
	// и ошибку
	UpdateTaskByID(id uint, task Task) (Task, error)
	// DeleteTaskByID - Передаем id для удаления, возвращаем только ошибку
	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	result := r.db.Model(&Task{}).Where("id = ?", id).Select("task", "is_done").Updates(task)
	if result.Error != nil {
		return Task{}, result.Error
	}

	// Проверяем, было ли обновлено хотя бы одно поле
	if result.RowsAffected == 0 {
		return Task{}, fmt.Errorf("задача с ID %d не найдена", id)
	}

	// Получаем обновленную задачу из базы данных
	var updatedTask Task
	if err := r.db.First(&updatedTask, id).Error; err != nil {
		return Task{}, err // Если не удалось найти обновленную задачу
	}

	return updatedTask, nil // Возвращаем обновленную задачу
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	result := r.db.Delete(&Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
