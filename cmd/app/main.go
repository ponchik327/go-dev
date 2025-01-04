package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"project/internal/database"
	"project/internal/handlers"
	"project/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repository := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repository)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", handler.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/tasks", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
