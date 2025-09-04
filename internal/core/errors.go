package core

import (
	"net/http"
)

const InternalErrorMsg = "We hit a snag while processing your request. Don't worry - it's not you, it's us! Please try again in a moment"
const InternalErrorTitle = "Oops! Something went wrong"

const JsonEncodingReason = "JSON_ENCODING_FAILED"
const JsonDecodingReason = "INVALID_JSON"
const UnknownReason = "UNKNOWN"

type AppError struct {
	Title    string
	Message  string
	Reason   string
	HttpCode int
}

func NewError(title string, message string, reason string, httpCode int) *AppError {
	return &AppError{
		Title:    title,
		Message:  message,
		Reason:   reason,
		HttpCode: httpCode,
	}
}

func DefaultInternal(reason string) *AppError {
	return &AppError{
		Title:    InternalErrorTitle,
		Message:  InternalErrorMsg,
		Reason:   reason,
		HttpCode: http.StatusInternalServerError,
	}
}
