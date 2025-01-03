package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/tasks", GetHandler).Methods("GET")
	router.HandleFunc("/api/tasks", PostHandler).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", PatchHandler).Methods("PATCH")
	router.HandleFunc("/api/tasks/{id}", DeleteHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
