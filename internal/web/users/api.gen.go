// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// Task defines model for Task.
type Task struct {
	Content *string `json:"content,omitempty"`
	Id      *uint   `json:"id,omitempty"`
	IsDone  *bool   `json:"is_done,omitempty"`
	UserId  *uint   `json:"user_id,omitempty"`
}

// User defines model for User.
type User struct {
	Email        *string `json:"email,omitempty"`
	Id           *uint   `json:"id,omitempty"`
	PasswordHash *string `json:"password_hash,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = User

// PatchUsersIDJSONRequestBody defines body for PatchUsersID for application/json ContentType.
type PatchUsersIDJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Create a new user
	// (POST /users)
	PostUsers(ctx echo.Context) error
	// Delete user by ID
	// (DELETE /users/{ID})
	DeleteUsersID(ctx echo.Context, iD uint) error
	// Update user by ID
	// (PATCH /users/{ID})
	PatchUsersID(ctx echo.Context, iD uint) error
	// Get all tasks some user by userID
	// (GET /users/{ID}/tasks)
	GetUsersIDTasks(ctx echo.Context, iD uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// DeleteUsersID converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ID" -------------
	var iD uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "ID", runtime.ParamLocationPath, ctx.Param("ID"), &iD)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsersID(ctx, iD)
	return err
}

// PatchUsersID converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUsersID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ID" -------------
	var iD uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "ID", runtime.ParamLocationPath, ctx.Param("ID"), &iD)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUsersID(ctx, iD)
	return err
}

// GetUsersIDTasks converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersIDTasks(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "ID" -------------
	var iD uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "ID", runtime.ParamLocationPath, ctx.Param("ID"), &iD)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter ID: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsersIDTasks(ctx, iD)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.DELETE(baseURL+"/users/:ID", wrapper.DeleteUsersID)
	router.PATCH(baseURL+"/users/:ID", wrapper.PatchUsersID)
	router.GET(baseURL+"/users/:ID/tasks", wrapper.GetUsersIDTasks)

}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []User

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUsersRequestObject struct {
	Body *PostUsersJSONRequestBody
}

type PostUsersResponseObject interface {
	VisitPostUsersResponse(w http.ResponseWriter) error
}

type PostUsers201JSONResponse User

func (response PostUsers201JSONResponse) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUsersIDRequestObject struct {
	ID uint `json:"ID"`
}

type DeleteUsersIDResponseObject interface {
	VisitDeleteUsersIDResponse(w http.ResponseWriter) error
}

type DeleteUsersID204Response struct {
}

func (response DeleteUsersID204Response) VisitDeleteUsersIDResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type PatchUsersIDRequestObject struct {
	ID   uint `json:"ID"`
	Body *PatchUsersIDJSONRequestBody
}

type PatchUsersIDResponseObject interface {
	VisitPatchUsersIDResponse(w http.ResponseWriter) error
}

type PatchUsersID200JSONResponse User

func (response PatchUsersID200JSONResponse) VisitPatchUsersIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUsersIDTasksRequestObject struct {
	ID uint `json:"ID"`
}

type GetUsersIDTasksResponseObject interface {
	VisitGetUsersIDTasksResponse(w http.ResponseWriter) error
}

type GetUsersIDTasks200JSONResponse []Task

func (response GetUsersIDTasks200JSONResponse) VisitGetUsersIDTasksResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Create a new user
	// (POST /users)
	PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error)
	// Delete user by ID
	// (DELETE /users/{ID})
	DeleteUsersID(ctx context.Context, request DeleteUsersIDRequestObject) (DeleteUsersIDResponseObject, error)
	// Update user by ID
	// (PATCH /users/{ID})
	PatchUsersID(ctx context.Context, request PatchUsersIDRequestObject) (PatchUsersIDResponseObject, error)
	// Get all tasks some user by userID
	// (GET /users/{ID}/tasks)
	GetUsersIDTasks(ctx context.Context, request GetUsersIDTasksRequestObject) (GetUsersIDTasksResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUsers operation middleware
func (sh *strictHandler) PostUsers(ctx echo.Context) error {
	var request PostUsersRequestObject

	var body PostUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsers(ctx.Request().Context(), request.(PostUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersResponseObject); ok {
		return validResponse.VisitPostUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUsersID operation middleware
func (sh *strictHandler) DeleteUsersID(ctx echo.Context, iD uint) error {
	var request DeleteUsersIDRequestObject

	request.ID = iD

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUsersID(ctx.Request().Context(), request.(DeleteUsersIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUsersID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUsersIDResponseObject); ok {
		return validResponse.VisitDeleteUsersIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUsersID operation middleware
func (sh *strictHandler) PatchUsersID(ctx echo.Context, iD uint) error {
	var request PatchUsersIDRequestObject

	request.ID = iD

	var body PatchUsersIDJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUsersID(ctx.Request().Context(), request.(PatchUsersIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUsersID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUsersIDResponseObject); ok {
		return validResponse.VisitPatchUsersIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetUsersIDTasks operation middleware
func (sh *strictHandler) GetUsersIDTasks(ctx echo.Context, iD uint) error {
	var request GetUsersIDTasksRequestObject

	request.ID = iD

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsersIDTasks(ctx.Request().Context(), request.(GetUsersIDTasksRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsersIDTasks")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersIDTasksResponseObject); ok {
		return validResponse.VisitGetUsersIDTasksResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
