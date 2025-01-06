package taskService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content string `json:"content"` // Наш сервер будет ожидать json c полем text
	IsDone  bool   `json:"is_done"` // В GO используем CamelCase, в Json - snake
}
