package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"

	"github.com/shivamvijaywargi/glink/internal/httperror"
	"github.com/shivamvijaywargi/glink/pkg/utils"
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

// TODO: Can be refactored as a utility function
func generateRandomShortUrl(length int) (string, error) {
	charSetLen := big.NewInt(int64(len(shortUrlCharSet)))
	result := make([]byte, length)

	for i := range result {
		randomNum, err := rand.Int(rand.Reader, charSetLen)
		if err != nil {
			return "", err
		}

		result[i] = shortUrlCharSet[randomNum.Int64()]
	}

	return string(result), nil
}

func GetAllUrls(w http.ResponseWriter, r *http.Request) {
	resp := utils.Response{
		Success: true,
		Message: "Fetched all the shortened URLs successfully",
		Data:    shortUrls,
	}

	utils.JsonResponse(w, resp, http.StatusOK)
}

func RedirectUsingShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.PathValue("shortUrl")

	for i := 0; i < len(shortUrls); i++ {
		if shortUrls[i].ShortUrl == shortUrl {
			w.WriteHeader(308)

			http.Redirect(w, r, shortUrls[i].OriginalUrl, http.StatusPermanentRedirect)
			return
		}
	}

	httperror.Writef(w, http.StatusNotFound, "Invalid ID or does not exist")
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("%+v", shortUrls)

	var urlObj UrlObj

	if err = json.Unmarshal(body, &urlObj); err != nil {
		httperror.Writef(w, http.StatusBadRequest, err.Error())
		return
	}

	if urlObj.ShortUrl != "" {
		for i := 0; i < len(shortUrls); i++ {
			if shortUrls[i].ShortUrl == urlObj.ShortUrl {
				httperror.Writef(w, http.StatusConflict, "Short URL that you provided is already taken, please try a different one.")
				return
			}
		}
	} else {
		// TODO: After generating a random URL need to verify if it is not duplicate as well
		urlObj.ShortUrl, err = generateRandomShortUrl(8)

		if err != nil {
			httperror.Writef(w, http.StatusInternalServerError, "Something went wrong, please try again later.")
			return
		}
	}

	urlObj.Id = len(shortUrls) + 1

	shortUrls = append(shortUrls, urlObj)

	resp := utils.Response{
		Success: true,
		Message: "Short URL created successfully",
		Data:    urlObj,
	}

	fmt.Printf("%+v \n", shortUrls)

	utils.JsonResponse(w, resp, 201)
}

func UpdateShortUrl(w http.ResponseWriter, r *http.Request) {
	urlId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || urlId < 1 {
		httperror.Writef(w, http.StatusBadRequest, err.Error())
		return
	}

	body, err := io.ReadAll(r.Body)

	var urlObj UrlObj

	if err = json.Unmarshal(body, &urlObj); err != nil {
		httperror.Writef(w, http.StatusBadRequest, err.Error())
		return
	}

	for i := 0; i < len(shortUrls); i++ {
		if shortUrls[i].Id == urlId {
			shortUrls[i].OriginalUrl = urlObj.OriginalUrl

			resp := utils.Response{
				Success: true,
				Message: fmt.Sprint("Updated the Short URL with ID: ", urlId),
				Data:    shortUrls[i],
			}

			utils.JsonResponse(w, resp, http.StatusCreated)
			return
		}
	}

	httperror.Writef(w, http.StatusNotFound, "ID is invalid or does not exist.")
}

func DeleteShortUrl(w http.ResponseWriter, r *http.Request) {
	urlId, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, urlObj := range shortUrls {
		if urlObj.Id == urlId {
			shortUrls = append(shortUrls[:i], shortUrls[i+1:]...)

			resp := utils.Response{
				Success: true,
				Message: "URL mapping deleted successfully",
			}

			utils.JsonResponse(w, resp, http.StatusOK)
			return
		}
	}

	httperror.Writef(w, http.StatusNotFound, "Invalid ID or does not exist")
}
