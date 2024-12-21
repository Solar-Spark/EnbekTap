package main

import (
	"enbektap/router"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/vacancies", router.EnableCORS(http.HandlerFunc(router.Handlers)))
	fmt.Println("Сервер запущен на порту 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
