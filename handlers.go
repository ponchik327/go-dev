package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var msgs []Message
	result := DB.Find(&msgs)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	for i, msg := range msgs {
		fmt.Fprintf(w, "задача номер: %d, Task=%s, Is_done=%v\n", i, msg.Task, msg.IsDone)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	msg, err := loadTaskFromJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := DB.Create(&msg)
	if result.Error != nil {
		http.Error(w, "Не удалось добавить новую задачу: "+result.Error.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Добавлена задача: Task=%s, Is_done=%v\n", msg.Task, msg.IsDone)
}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	msg, err := findTaskByID(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMsg, err := loadTaskFromJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg.Task = newMsg.Task
	msg.IsDone = newMsg.IsDone

	if err := DB.Save(&msg).Error; err != nil {
		http.Error(w, "Не удалось обновить задачу", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Задача с id=%s обновлена: Task=%s, Is_done=%v\n", idStr, msg.Task, msg.IsDone)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	msg, err := findTaskByID(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := DB.Delete(&msg).Error; err != nil {
		http.Error(w, "Не удалось удалить задачу: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Задача с id=%s успешна удалена", idStr)
}
