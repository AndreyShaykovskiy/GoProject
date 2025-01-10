package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// Создаем срез для хранения всех строк таблицы Message
	var messages []Message

	if err := DB.Find(&messages).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Возвращаем ошибку, если запрос не удался
		return
	}

	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "application/json")

	// Преобразуем срез сообщений в JSON
	response, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Возвращаем ошибку, если преобразование не удалось
		return
	}

	// Отправляем JSON-ответ клиенту
	w.WriteHeader(http.StatusOK) // Устанавливаем статус 200 OK
	w.Write(response)            // Отправляем ответ клиенту
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var message Message
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close() // Закрываем тело запроса после декодирования

		if err := decoder.Decode(&message); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем ошибку, если декодирование не удалось
			return
		}

		// Возвращаем созданное сообщение
		response, err := json.Marshal(message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated) // Устанавливаем статус 201 Created
		w.Write(response)                 // Отправляем ответ клиенту

		// создаем новую запись в базе данных
		result := DB.Create(&Message{Task: message.Task, IsDone: message.IsDone})
		if result.Error != nil {
			log.Fatal("Ошибка: ", result.Error) // обработка ошибки создания
		}

	} else {
		http.Error(w, "Недопустимый метод запроса", http.StatusBadRequest)
	}

}

func main() {
	// Вызываем метод InitDB() из файла db.go
	InitDB()

	// Автоматическая миграция модели Message
	DB.AutoMigrate(&Message{})

	// Создаем маршрутизатор
	router := mux.NewRouter()

	// Определяем маршруты для нашего API
	router.HandleFunc("/api/get", GetHandler).Methods("GET")
	router.HandleFunc("/api/post", PostHandler).Methods("POST")

	//Запуск сервера
	http.ListenAndServe(":8080", router)
}
