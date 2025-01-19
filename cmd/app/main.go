package main

import (
	"FirstTask/internal/database"
	"FirstTask/internal/handlers"
	"FirstTask/internal/tasksService"
	"FirstTask/internal/userService"
	"FirstTask/internal/web/tasks"
	"FirstTask/internal/web/users"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	// Инициализация валидатора
	validate := validator.New()
	database.InitDB()
	err := database.DB.AutoMigrate(&tasksService.Task{})
	if err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
	}

	err = database.DB.AutoMigrate(&userService.User{})
	if err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
	}

	tasksRepo := tasksService.NewTaskRepository(database.DB)
	userRepo := userService.NewUserRepository(database.DB)
	TasksService := tasksService.NewTaskService(tasksRepo)
	UserService := userService.NewUserService(userRepo)

	tasksHandler := handlers.NewTaskHandler(TasksService)
	userHandler := handlers.NewUserHandler(UserService, validate)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictTasksHandler := tasks.NewStrictHandler(tasksHandler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictTasksHandler)
	strictUserHandler := users.NewStrictHandler(userHandler, nil) // тут будет ошибка
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
