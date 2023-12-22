package infrastructure

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewApiError(message string, code int) *ApiError {
	return &ApiError{Message: message, Code: code}
}

func (e *ApiError) Error() string {
	return e.Message
}

func SendResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
