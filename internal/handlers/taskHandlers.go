package handlers

import (
	"FirstTask/internal/tasksService" // Импортируем наш сервис
	"FirstTask/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *tasksService.TaskService
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	// Извлекаем ID задачи из запроса
	taskID := request.Id

	// Создаем переменную для обновленной задачи
	var updatedTask tasksService.Task

	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body

	// Обращаемся к сервису и создаем задачу
	if taskRequest.Task != nil {
		updatedTask.Task = *taskRequest.Task
	}
	if taskRequest.IsDone != nil {
		updatedTask.IsDone = *taskRequest.IsDone
	}

	// Обновляем задачу в сервисе
	updatedTask, err := h.Service.UpdateTaskByID(taskID, updatedTask)
	if err != nil {
		return nil, err
	}

	// Создаем ответ с обновленной задачей
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	// Возвращаем ответ
	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	// Извлекаем ID задачи из запроса
	taskID := request.Id

	// Вызываем метод сервиса для удаления задачи
	err := h.Service.DeleteTaskByID(taskID)
	if err != nil {
		// Если произошла ошибка, возвращаем ее
		return nil, err
	}

	// Возвращаем пустой ответ, так как статус 204 No Content
	response := tasks.DeleteTasksId204Response{}
	return response, nil

}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	// Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := tasks.GetTasks200JSONResponse{}

	// Заполняем слайс response всеми задачами из БД
	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	taskRequest := request.Body
	// Обращаемся к сервису и создаем задачу
	taskToCreate := tasksService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}
	// Просто возвращаем респонс!
	return response, nil
}

// Нужна для создания структуры Handler на этапе инициализации приложения

func NewHandler(service *tasksService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
