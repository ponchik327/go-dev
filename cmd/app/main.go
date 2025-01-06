package main

import (
	"log"

	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
	"project/internal/web/tasks"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := database.NewDB()
	db.InitDB()

	repository := taskService.NewTaskRepository(db.Db)
	service := taskService.NewTaskService(repository)

	handler := handlers.NewHandler(service)

	// Инициализируем echo
	server := echo.New()

	// используем Logger и Recover
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	baseURL := server.Group("/api")

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(baseURL, strictHandler)

	err := server.Start(":8080")
	if err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
