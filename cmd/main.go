package main

import (
	"log"
	"net/http"

	"github.com/shivamvijaywargi/glink/internal/handlers"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello World"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("POST /", handlers.CreateShortUrl)
	mux.HandleFunc("PATCH /{id}", handlers.UpdateShortUrl)
	mux.HandleFunc("DELETE /{id}", handlers.DeleteShortUrl)

	log.Print("Listening on :8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal("ListenAndServe:", err)
}
