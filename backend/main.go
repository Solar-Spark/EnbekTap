package main

import (
	"enbektap/router"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/vacancies", router.Handlers)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
