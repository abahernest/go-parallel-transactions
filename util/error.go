package util

import (
	"net/http"
)


func RespondWithError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Write([]byte(message))
}