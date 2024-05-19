package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello World"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)

	log.Print("Listening on :8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}
