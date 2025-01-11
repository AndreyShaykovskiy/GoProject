package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
	}

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
	return

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
	}

	var message Message
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close() // Закрываем тело запроса после декодирования

	if err := decoder.Decode(&message); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем ошибку, если декодирование не удалось
		return
	}

	// Создаем новую запись в базе данных
	result := DB.Create(&message) //&Message{Task: message.Task, IsDone: message.IsDone})
	if result.Error != nil {
		http.Error(w, "Ошибка создания: "+result.Error.Error(), http.StatusInternalServerError)
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
	return

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
	}

	// Извлекаем ID из URL-параметра
	vars := mux.Vars(r)
	id := vars["id"]

	// Удаляем строку в базе данных по id
	result := DB.Delete(&Message{}, id) // Передаем указатель на структуру Message и ID
	//result := DB.Model(&Message{}).Where("id = ?", id).Update("task", "")
	if result.Error != nil {
		http.Error(w, "Ошибка удаления: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Проверка на удаление строки
	if result.RowsAffected == 0 {
		http.Error(w, "Данная строка не найдена", http.StatusNotFound)
		return
	}

	// Успешное удаление
	w.WriteHeader(http.StatusNoContent)
	return

}

func PatchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Недопустимый метод запроса", http.StatusMethodNotAllowed)
	}

	// Извлекаем ID из URL-параметра
	vars := mux.Vars(r)
	id := vars["id"]

	// Создаем переменную для хранения обновленных данных
	var updateMessage Message

	// Декодируем JSON из тела запроса
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close() // Закрываем тело запроса после декодирования

	if err := decoder.Decode(&updateMessage); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем ошибку, если декодирование не удалось
		return
	}

	// Обновляем поле task по переданному id
	result := DB.Model(&Message{}).Where("id = ?", id).Select("task", "is_done").Updates(updateMessage)
	if result.Error != nil {
		http.Error(w, "Ошибка обновления: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Проверка на наличие обновленной записи
	if result.RowsAffected == 0 {
		http.Error(w, "Данная строка не найдена", http.StatusNotFound)
	}

	// Извлекаем обновлённую запись из базы данных
	var updatedMessage Message
	if err := DB.First(&updatedMessage, id).Error; err != nil {
		http.Error(w, "Ошибка получения обновлённой записи: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем созданное сообщение
	response, err := json.Marshal(updatedMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return

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
	router.HandleFunc("/api/delete/{id}", DeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/patch/{id}", PatchHandler).Methods("PATCH")

	//Запуск сервера
	http.ListenAndServe(":8080", router)
}
