package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Task struct {
	Task string `json:"task"`
}

var task Task

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("hello, %s", task.Task)
	fmt.Fprintln(w, response) // Отправляем ответ клиенту
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close() // Закрываем тело запроса после декодирования

		if err := decoder.Decode(&task); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Возвращаем ошибку, если декодирование не удалось
			return
		}

	} else {
		http.Error(w, "Недопустимый метод запроса", http.StatusBadRequest)
	}
}

func main() {
	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
