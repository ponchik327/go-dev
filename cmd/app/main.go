package main

import (
	"log"

	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
	"project/internal/userService"
	"project/internal/web/tasks"
	"project/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := database.NewDB()
	db.InitDB()

	// Инициализируем сервис задач
	tasksRepository := taskService.NewTaskRepository(db.Db)
	tasksService := taskService.NewTaskService(tasksRepository)

	// Инициализируем сервис пользователей
	usersRepository := userService.NewUserRepository(db.Db)
	usersService := userService.NewUserService(usersRepository)

	// Инициализируем echo
	server := echo.New()

	// используем Logger и Recover
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	baseURL := server.Group("/api")

	// Передаем и регистрируем хендлеры для сервиса задач в echo
	tasksHandlers := handlers.NewTaskHandlers(tasksService)
	taskStrictHandler := tasks.NewStrictHandler(tasksHandlers, nil)
	tasks.RegisterHandlers(baseURL, taskStrictHandler)

	// Передаем и регистрируем хендлеры для сервиса пользователей в echo
	usersHandlers := handlers.NewUserHandlers(usersService)
	userStrictHandler := users.NewStrictHandler(usersHandlers, nil)
	users.RegisterHandlers(baseURL, userStrictHandler)

	err := server.Start(":8080")
	if err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
