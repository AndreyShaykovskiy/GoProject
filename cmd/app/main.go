package main

import (
	"FirstTask/internal/database"
	"FirstTask/internal/handlers"
	"FirstTask/internal/tasksService"
	"FirstTask/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&tasksService.Task{})
	if err != nil {
		log.Fatalf("Error during AutoMigrate: %v", err)
	}

	repo := tasksService.NewTaskRepository(database.DB)
	service := tasksService.NewService(repo)

	handler := handlers.NewHandler(service)

	// Инициализируем echo
	e := echo.New()

	// используем Logger и Recover
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
	strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
	tasks.RegisterHandlers(e, strictHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
