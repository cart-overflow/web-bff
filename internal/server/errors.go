package server

import (
	"encoding/json"
	"net/http"

	"github.com/cart-overflow/web-bff/internal/core"
	"github.com/cart-overflow/web-bff/internal/dto"
)

func WriteDefaultInternal(w http.ResponseWriter, reason string) {
	WriteAppError(w, core.DefaultInternal(reason))
}

func WriteAppError(w http.ResponseWriter, err *core.AppError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HttpCode)
	dto := dto.AppError{Title: err.Title, Message: err.Message, Reason: err.Reason}
	encoder := json.NewEncoder(w)
	encoder.Encode(dto)
}

func WriteDefaultInternalWeb(w http.ResponseWriter) {
	WriteErrorWeb(w, core.InternalErrorMsg, http.StatusInternalServerError)
}

func WriteErrorWeb(w http.ResponseWriter, msg string, httpCode int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(httpCode)
	w.Write([]byte(msg))
}
