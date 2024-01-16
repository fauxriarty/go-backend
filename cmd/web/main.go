package main

import (
	"net/http"

	"github.com/fauxriarty/go-backend/pkg/handlers"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":8080", nil)
}
