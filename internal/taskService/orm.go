package taskService

import "gorm.io/gorm"

// Определяем структуру Message, которая будет представлять таблицу в базе данных
type Task struct {
	gorm.Model        // По сути, прописывая gorm.Model, мы просто добавляем поля ID, CreatedAt, UpdatedAt и DeletedAt
	Task       string `json:"task"`    // Наш сервер будет ожидать json c полем text
	IsDone     bool   `json:"is_done"` // В GO используем CamelCase, в Json - snake
}
