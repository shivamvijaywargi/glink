package main

import (
	"log"
	"net/http"

	"github.com/shivamvijaywargi/glink/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.GetAllUrls)
	mux.HandleFunc("POST /", handlers.CreateShortUrl)
	mux.HandleFunc("PATCH /{id}", handlers.UpdateShortUrl)
	mux.HandleFunc("DELETE /{id}", handlers.DeleteShortUrl)

	log.Print("Listening on :8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal("ListenAndServe:", err)
}
