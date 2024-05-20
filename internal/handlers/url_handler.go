package handlers

import (
	"fmt"
	"net/http"
)

type UrlObj struct {
	Id          int    `json:"id"`
	OriginalUrl string `json:"originalUrl"`
	ShortUrl    string `json:"shortUrl"`
}

var shortUrls []UrlObj

var shortUrlCharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type PostUrlHandlerParams struct {
	originalUrl string
	shortUrl    string
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Print(r)
}

func UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	urlId := r.PathValue("id")

	msg := fmt.Sprint("Update a specific snippet by ID: ", urlId)

	w.Write([]byte(msg))
}

func DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r)
}
