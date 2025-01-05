package handlers

import (
	"context"

	"project/internal/taskService"
	"project/internal/web/tasks"
)

type Handler struct {
	service *taskService.TaskService
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{service: service}
}

// GetTasks implements tasks.StrictServerInterface.
func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Text:   &tsk.Text,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

// PostTasks implements tasks.StrictServerInterface.
func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskService.Task{
		Text:   *taskRequest.Text,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Text:   &createdTask.Text,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

// PatchTasksID implements tasks.StrictServerInterface.
func (h *Handler) PatchTasksID(ctx context.Context, request tasks.PatchTasksIDRequestObject) (tasks.PatchTasksIDResponseObject, error) {
	taskRequest := request.Body

	taskToUpdate := taskService.Task{
		Text:   *taskRequest.Text,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.service.UpdateTaskByID(request.ID, taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksID200JSONResponse{
		Id:     &updatedTask.ID,
		Text:   &updatedTask.Text,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

// DeleteTasksID implements tasks.StrictServerInterface.
func (h *Handler) DeleteTasksID(ctx context.Context, request tasks.DeleteTasksIDRequestObject) (tasks.DeleteTasksIDResponseObject, error) {
	err := h.service.DeleteTaskByID(request.ID)
	if err != nil {
		return nil, err
	}

	response := tasks.DeleteTasksID204Response{}

	return response, nil
}
