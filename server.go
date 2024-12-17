package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура для данных JSON
type RequestData struct {
	Message *string `json:"message"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/json", handleJSON)
	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe("localhost:8080", nil)
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	// Установка заголовков для ответа
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		var requestData RequestData
		// Чтение JSON из тела запроса
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil || requestData.Message == nil {
			// Некорректный JSON
			jsonResponse(w, http.StatusBadRequest, "fail", "Invalid JSON message")
			return
		}

		// Логирование сообщения в консоль
		fmt.Println("Получено сообщение:", *requestData.Message)

		// Успешный ответ
		jsonResponse(w, http.StatusOK, "success", "Data successfully received")
		return
	}

	if r.Method == http.MethodGet {
		// Для GET запросов просто возвращаем JSON ответ
		jsonResponse(w, http.StatusOK, "success", "GET request received")
		return
	}

	// Если метод не поддерживается
	jsonResponse(w, http.StatusMethodNotAllowed, "fail", "Invalid HTTP method")
}

// Утилита для отправки JSON ответа
func jsonResponse(w http.ResponseWriter, statusCode int, status, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ResponseData{
		Status:  status,
		Message: message,
	})
}
