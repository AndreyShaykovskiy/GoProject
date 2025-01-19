package handlers

import (
	"FirstTask/internal/userService"
	"FirstTask/internal/web/users"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	Service   *userService.UserService
	Validator *validator.Validate
}

// Нужна для создания структуры UserHandler на этапе инициализации приложения

func NewUserHandler(service *userService.UserService, validator *validator.Validate) *UserHandler {
	return &UserHandler{
		Service:   service,
		Validator: validator,
	}
}

func (h *UserHandler) PatchUserId(ctx context.Context, request users.PatchUserIdRequestObject) (users.PatchUserIdResponseObject, error) {
	// Извлекаем ID пользователя из запроса
	userID := request.Id

	// Создаем переменную для обновленной пользователя
	var updatedUser userService.User

	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body

	// Проверяем и обновляем email, если он предоставлен
	if userRequest.Email != nil {
		updatedUser.Email = *userRequest.Email

		// Валидация email
		if err := h.Validator.Var(updatedUser.Email, "required,email"); err != nil {
			return nil, fmt.Errorf("некорректный формат email: %v", err)
		}
	}

	// Проверяем и обновляем пароль, если он предоставлен
	if userRequest.Password != nil {
		if len(*userRequest.Password) < 8 {
			return nil, fmt.Errorf("пароль должен содержать не менее 8 символов")
		}
		updatedUser.Password = *userRequest.Password
	}

	// Обновляем пользователя в сервисе
	updatedUser, err := h.Service.UpdateUserBuID(userID, updatedUser)
	if err != nil {
		return nil, err
	}

	// Создаем ответ с обновленным пользователем
	response := users.PatchUserId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	// Возвращаем ответ
	return response, nil
}

func (h *UserHandler) DeleteUserId(ctx context.Context, request users.DeleteUserIdRequestObject) (users.DeleteUserIdResponseObject, error) {
	// Извлекаем ID юзера из запроса
	userID := request.Id

	// Вызываем метод сервиса для удаления юзера
	err := h.Service.DeleteUserByID(userID)
	if err != nil {
		// Если произошла ошибка, возвращаем ее
		return nil, err
	}

	// Возвращаем пустой ответ, так как статус 204 No Content
	response := users.DeleteUserId204Response{}
	return response, nil

}

func (h *UserHandler) GetUser(_ context.Context, _ users.GetUserRequestObject) (users.GetUserResponseObject, error) {
	// Получение всех юзеров из сервиса
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := users.GetUser200JSONResponse{}

	// Заполняем слайс response всеми пользователями из БД
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUser(_ context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и создаем пользователя
	userToCreate := userService.User{
		Email:    string(*userRequest.Email),
		Password: *userRequest.Password,
	}

	// Валидация пользователя
	if err := userToCreate.Validate(); err != nil {
		// Если валидация не прошла, возвращаем ошибку
		return nil, fmt.Errorf("ошибка валидации: %v", err)
	}

	// Обращаемся к сервису и создаем пользователя
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PostUser201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	// Просто возвращаем респонс!
	return response, nil
}
