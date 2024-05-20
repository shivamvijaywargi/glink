package httperror

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPError struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	httpError := HTTPError{
		Success: false,
		Code:    code,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(httpError)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding error"))
		return
	}
}

func Writef(w http.ResponseWriter, code int, message string, args ...interface{}) {
	WriteError(w, code, fmt.Sprintf(message, args...))
}
