package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура для данных JSON
type RequestData struct {
	Message string `json:"message"`
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
		if err != nil || requestData.Message == "" {
			// Некорректный JSON
			jsonResponse(w, http.StatusBadRequest, "fail", "Invalid JSON message")
			return
		}

		// Логирование сообщения в консоль
		fmt.Println("Получено сообщение:", requestData.Message)

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

// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", HelloHandler)
// 	http.HandleFunc("/hello", MetGet)
// 	http.HandleFunc("/post", MetPost)
// 	fmt.Println("Now running")
// 	http.ListenAndServe("localhost:8080", nil)

// }

// var count int

// func HelloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "hello world")
// 	// fmt.Println("hellop")
// 	count++
// 	fmt.Println(count)

// }

// func MetGet(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		fmt.Fprintln(w, "Это функция GET")
// 	} else {
// 		fmt.Fprintln(w, "Что-то другое(гет)")
// 	}
// }
// func MetPost(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {
// 		fmt.Fprintln(w, "Это функция Пост")
// 	} else {
// 		fmt.Fprintln(w, "Что-то другое (пост)")
// 	}
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type Response struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// }

// func handlePost(w http.ResponseWriter, r *http.Request) {
// 	// Устанавливаем заголовки ответа
// 	w.Header().Set("Content-Type", "application/json")

// 	// Проверяем, что это POST-запрос
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	// Декодируем JSON из тела запроса
// 	var data map[string]interface{}
// 	err := json.NewDecoder(r.Body).Decode(&data)
// 	if err != nil {
// 		// Если ошибка в парсинге, возвращаем ошибку
// 		response := Response{
// 			Status:  "fail",
// 			Message: "Invalid JSON message",
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	// Проверяем наличие поля "message"
// 	if message, ok := data["message"].(string); ok {
// 		// Если все в порядке, возвращаем успешный ответ
// 		fmt.Println("Received message:", message) // Выводим полученное сообщение в консоль
// 		response := Response{
// 			Status:  "success",
// 			Message: "Data successfully received",
// 		}
// 		json.NewEncoder(w).Encode(response)
// 	} else {
// 		// Если "message" нет или его тип неправильный
// 		response := Response{
// 			Status:  "fail",
// 			Message: "Invalid JSON message",
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(response)
// 	}
// }

// func handleGet(w http.ResponseWriter, r *http.Request) {
// 	// Просто возвращаем статус с сообщением
// 	fmt.Println("GET request received") // Сообщение о получении GET-запроса
// 	response := Response{
// 		Status:  "success",
// 		Message: "GET request successful",
// 	}
// 	json.NewEncoder(w).Encode(response)
// }

// func main() {
// 	// Обработчик POST-запросов на /post
// 	http.HandleFunc("/post", handlePost)

// 	// Обработчик GET-запросов на /get
// 	http.HandleFunc("/get", handleGet)

// 	// Слушаем порт 8080
// 	fmt.Println("Server started at http://localhost:8080") // Сообщение о запуске сервера
// 	http.ListenAndServe(":8080", nil)
// }
