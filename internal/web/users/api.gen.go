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

// User defines model for User.
type User struct {
	Email    *string `json:"email,omitempty"`
	Id       *uint   `json:"id,omitempty"`
	Password *string `json:"password,omitempty"`
}

// PostUserJSONRequestBody defines body for PostUser for application/json ContentType.
type PostUserJSONRequestBody = User

// PatchUserIdJSONRequestBody defines body for PatchUserId for application/json ContentType.
type PatchUserIdJSONRequestBody = User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /user)
	GetUser(ctx echo.Context) error
	// Create a new user
	// (POST /user)
	PostUser(ctx echo.Context) error
	// Delete user by id
	// (DELETE /user/{id})
	DeleteUserId(ctx echo.Context, id uint) error
	// Patch user by id
	// (PATCH /user/{id})
	PatchUserId(ctx echo.Context, id uint) error
	// Get all tasks by user ID
	// (GET /user/{userId}/tasks)
	GetTasksByUserID(ctx echo.Context, userId uint) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUser(ctx)
	return err
}

// PostUser converts echo context to params.
func (w *ServerInterfaceWrapper) PostUser(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUser(ctx)
	return err
}

// DeleteUserId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUserId(ctx, id)
	return err
}

// PatchUserId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUserId(ctx, id)
	return err
}

// GetTasksByUserID converts echo context to params.
func (w *ServerInterfaceWrapper) GetTasksByUserID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId uint

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetTasksByUserID(ctx, userId)
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

	router.GET(baseURL+"/user", wrapper.GetUser)
	router.POST(baseURL+"/user", wrapper.PostUser)
	router.DELETE(baseURL+"/user/:id", wrapper.DeleteUserId)
	router.PATCH(baseURL+"/user/:id", wrapper.PatchUserId)
	router.GET(baseURL+"/user/:userId/tasks", wrapper.GetTasksByUserID)

}

type GetUserRequestObject struct {
}

type GetUserResponseObject interface {
	VisitGetUserResponse(w http.ResponseWriter) error
}

type GetUser200JSONResponse []User

func (response GetUser200JSONResponse) VisitGetUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUserRequestObject struct {
	Body *PostUserJSONRequestBody
}

type PostUserResponseObject interface {
	VisitPostUserResponse(w http.ResponseWriter) error
}

type PostUser201JSONResponse User

func (response PostUser201JSONResponse) VisitPostUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUserIdRequestObject struct {
	Id uint `json:"id"`
}

type DeleteUserIdResponseObject interface {
	VisitDeleteUserIdResponse(w http.ResponseWriter) error
}

type DeleteUserId204Response struct {
}

func (response DeleteUserId204Response) VisitDeleteUserIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type PatchUserIdRequestObject struct {
	Id   uint `json:"id"`
	Body *PatchUserIdJSONRequestBody
}

type PatchUserIdResponseObject interface {
	VisitPatchUserIdResponse(w http.ResponseWriter) error
}

type PatchUserId200JSONResponse User

func (response PatchUserId200JSONResponse) VisitPatchUserIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTasksByUserIDRequestObject struct {
	UserId uint `json:"userId"`
}

type GetTasksByUserIDResponseObject interface {
	VisitGetTasksByUserIDResponse(w http.ResponseWriter) error
}

type GetTasksByUserID200JSONResponse []struct {
	Id     uint   `json:"id"`
	IsDone bool   `json:"is_done"`
	Task   string `json:"task"`
}

func (response GetTasksByUserID200JSONResponse) VisitGetTasksByUserIDResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /user)
	GetUser(ctx context.Context, request GetUserRequestObject) (GetUserResponseObject, error)
	// Create a new user
	// (POST /user)
	PostUser(ctx context.Context, request PostUserRequestObject) (PostUserResponseObject, error)
	// Delete user by id
	// (DELETE /user/{id})
	DeleteUserId(ctx context.Context, request DeleteUserIdRequestObject) (DeleteUserIdResponseObject, error)
	// Patch user by id
	// (PATCH /user/{id})
	PatchUserId(ctx context.Context, request PatchUserIdRequestObject) (PatchUserIdResponseObject, error)
	// Get all tasks by user ID
	// (GET /user/{userId}/tasks)
	GetTasksByUserID(ctx context.Context, request GetTasksByUserIDRequestObject) (GetTasksByUserIDResponseObject, error)
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

// GetUser operation middleware
func (sh *strictHandler) GetUser(ctx echo.Context) error {
	var request GetUserRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUser(ctx.Request().Context(), request.(GetUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUserResponseObject); ok {
		return validResponse.VisitGetUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUser operation middleware
func (sh *strictHandler) PostUser(ctx echo.Context) error {
	var request PostUserRequestObject

	var body PostUserJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUser(ctx.Request().Context(), request.(PostUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUserResponseObject); ok {
		return validResponse.VisitPostUserResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUserId operation middleware
func (sh *strictHandler) DeleteUserId(ctx echo.Context, id uint) error {
	var request DeleteUserIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUserId(ctx.Request().Context(), request.(DeleteUserIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUserId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUserIdResponseObject); ok {
		return validResponse.VisitDeleteUserIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUserId operation middleware
func (sh *strictHandler) PatchUserId(ctx echo.Context, id uint) error {
	var request PatchUserIdRequestObject

	request.Id = id

	var body PatchUserIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUserId(ctx.Request().Context(), request.(PatchUserIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUserId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUserIdResponseObject); ok {
		return validResponse.VisitPatchUserIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetTasksByUserID operation middleware
func (sh *strictHandler) GetTasksByUserID(ctx echo.Context, userId uint) error {
	var request GetTasksByUserIDRequestObject

	request.UserId = userId

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetTasksByUserID(ctx.Request().Context(), request.(GetTasksByUserIDRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTasksByUserID")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetTasksByUserIDResponseObject); ok {
		return validResponse.VisitGetTasksByUserIDResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
