package handlers

import (
	"context"
	"fmt"
	"project/internal/models"
	"project/internal/userService"
	"project/internal/web/users"
)

type userHandlers struct {
	service *userService.UserService
}

func NewUserHandlers(service *userService.UserService) *userHandlers {
	return &userHandlers{service: service}
}

// GetUsers implements users.StrictServerInterface.
func (u *userHandlers) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("не удалось найти всех пользователей: %w", err)
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:           &usr.ID,
			Email:        &usr.Email,
			PasswordHash: &usr.PasswordHash,
		}
		response = append(response, user)
	}

	return response, nil
}

// PostUsers implements users.StrictServerInterface.
func (u *userHandlers) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := models.User{
		Email:        *userRequest.Email,
		PasswordHash: *userRequest.PasswordHash,
	}

	createdUser, err := u.service.CreateUser(userToCreate)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать пользователя: %w", err)
	}

	response := users.PostUsers201JSONResponse{
		Id:           &createdUser.ID,
		Email:        &createdUser.Email,
		PasswordHash: &createdUser.PasswordHash,
	}

	return response, nil
}

// PatchUsersID implements users.StrictServerInterface.
func (u *userHandlers) PatchUsersID(ctx context.Context, request users.PatchUsersIDRequestObject) (users.PatchUsersIDResponseObject, error) {
	userRequest := request.Body

	userToUpdate := models.User{
		Email:        *userRequest.Email,
		PasswordHash: *userRequest.Email,
	}

	updatedUser, err := u.service.UpdateUserByID(request.ID, userToUpdate)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить данные пользователя: %w", err)
	}

	response := users.PatchUsersID200JSONResponse{
		Id:           &updatedUser.ID,
		Email:        &updatedUser.Email,
		PasswordHash: &updatedUser.PasswordHash,
	}

	return response, nil
}

// DeleteUsersID implements users.StrictServerInterface.
func (u *userHandlers) DeleteUsersID(ctx context.Context, request users.DeleteUsersIDRequestObject) (users.DeleteUsersIDResponseObject, error) {
	err := u.service.DeleteUserByID(request.ID)
	if err != nil {
		return nil, fmt.Errorf("не удалось удалить пользователя: %w", err)
	}

	response := users.DeleteUsersID204Response{}

	return response, nil
}

// GetUsersIDTasks implements users.StrictServerInterface.
func (u *userHandlers) GetUsersIDTasks(ctx context.Context, request users.GetUsersIDTasksRequestObject) (users.GetUsersIDTasksResponseObject, error) {
	userTasks, err := u.service.GetTasksByUserID(request.ID)
	if err != nil {
		return nil, fmt.Errorf("не удалось найти задачи пользователя: %w", err)
	}

	response := users.GetUsersIDTasks200JSONResponse{}

	for _, tsk := range userTasks {
		task := users.Task{
			Id:      &tsk.ID,
			Content: &tsk.Content,
			IsDone:  &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}
