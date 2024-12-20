package main

import (
	"enbektap/router"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/vacancies", router.Handlers)
	fmt.Println("Сервер запущен на порту 5500...")
	http.ListenAndServe("localhost:5500", nil)
}
