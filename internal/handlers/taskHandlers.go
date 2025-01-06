package handlers

import (
	"context"
	"fmt"

	"project/internal/taskService"
	"project/internal/web/tasks"
)

type taskHandlers struct {
	service *taskService.TaskService
}

func NewTaskHandlers(service *taskService.TaskService) *taskHandlers {
	return &taskHandlers{service: service}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *taskHandlers) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, fmt.Errorf("не удалось найти все задачи: %w", err)
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:      &tsk.ID,
			Content: &tsk.Content,
			IsDone:  &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *taskHandlers) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Content: *taskRequest.Content,
		IsDone:  *taskRequest.IsDone,
	}

	createdTask, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать задачу: %w", err)
	}

	response := tasks.PostTasks201JSONResponse{
		Id:      &createdTask.ID,
		Content: &createdTask.Content,
		IsDone:  &createdTask.IsDone,
	}

	return response, nil
}

// PatchTasksID implements tasks.StrictServerInterface.
func (h *taskHandlers) PatchTasksID(ctx context.Context, request tasks.PatchTasksIDRequestObject) (tasks.PatchTasksIDResponseObject, error) {
	taskRequest := request.Body

	taskToUpdate := taskService.Task{
		Content: *taskRequest.Content,
		IsDone:  *taskRequest.IsDone,
	}

	updatedTask, err := h.service.UpdateTaskByID(request.ID, taskToUpdate)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить задачу: %w", err)
	}

	response := tasks.PatchTasksID200JSONResponse{
		Id:      &updatedTask.ID,
		Content: &updatedTask.Content,
		IsDone:  &updatedTask.IsDone,
	}

	return response, nil
}

// DeleteTasksID implements tasks.StrictServerInterface.
func (h *taskHandlers) DeleteTasksID(ctx context.Context, request tasks.DeleteTasksIDRequestObject) (tasks.DeleteTasksIDResponseObject, error) {
	err := h.service.DeleteTaskByID(request.ID)
	if err != nil {
		return nil, fmt.Errorf("не удалось удалить задачу: %w", err)
	}

	response := tasks.DeleteTasksID204Response{}

	return response, nil
}
