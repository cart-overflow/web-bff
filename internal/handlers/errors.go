package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cart-overflow/web-bff/internal/dto"
)

const SERVER_ERROR_TITLE = "Oops! Something went wrong"

const SERVER_ERROR_MSG = "We hit a snag while processing your request. Don't worry - it's not you, it's us! Please try again in a moment"

func WriteInternalServerWebError(w http.ResponseWriter) {
	WriteWebError(w, SERVER_ERROR_MSG, http.StatusInternalServerError)
}

// TODO: make a pretty html error page
func WriteWebError(w http.ResponseWriter, msg string, httpCode int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(httpCode)
	w.Write([]byte(msg))
}

func WriteInternalServerError(w http.ResponseWriter) {
	WriteAppError(
		w,
		SERVER_ERROR_TITLE,
		SERVER_ERROR_MSG,
		"INTERNAL_SERVER_ERROR",
		http.StatusInternalServerError,
	)
}

func WriteAppError(w http.ResponseWriter, title string, msg string, errCode string, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	dto := dto.AppError{Title: title, Message: msg, Code: errCode}
	encoder := json.NewEncoder(w)
	encoder.Encode(dto)
}
