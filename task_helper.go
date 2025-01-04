package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

func findTaskByID(idStr string) (Message, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return Message{}, fmt.Errorf("неверный id задачи: %w", err)
	}

	var msg Message
	result := DB.First(&msg, id)
	if result.Error != nil {
		return msg, fmt.Errorf("не удалось найти задачу: %w", err)
	}

	return msg, nil
}

func loadTaskFromJSON(payload io.Reader) (Message, error) {
	var msg Message
	err := json.NewDecoder(payload).Decode(&msg)
	if err != nil {
		return msg, fmt.Errorf("не получилось считать задачу из JSON %w", err)
	}
	return msg, nil
}
