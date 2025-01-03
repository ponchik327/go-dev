package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var msgs []Message
	result := DB.Find(&msgs)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
	}

	for i, msg := range msgs {
		fmt.Fprintf(w, "задача номер: %d, Task=%s, Is_done=%v\n", i, msg.Task, msg.IsDone)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Ошибка при декодировании JSON", http.StatusBadRequest)
		return
	}

	result := DB.Create(&msg)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "В базу данных добавлена задача: %s, со статусом %v\n", msg.Task, msg.IsDone)
}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()

	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/tasks", GetHandler).Methods("GET")
	router.HandleFunc("/api/tasks", PostHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}
