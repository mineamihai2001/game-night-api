package api_error

import (
	"net/http"

	"github.com/mineamihai2001/game-night/internal/infrastructure/services"
)

type ApiError struct {
	StatusCode        int    `json:"statusCode"`
	Error             string `json:"error"`
	Message           string `json:"message"`
	InternalErrorCode string `json:"internalErrorCode"`
}

func New(statusCode int, err error) ApiError {
	var internalErrorCode string

	switch e := err.(type) {
	default:
		internalErrorCode = "0x000000"
	case *services.ServiceError:
		internalErrorCode = e.StringErrorCode()
	}

	return ApiError{
		StatusCode:        statusCode,
		Error:             http.StatusText(statusCode),
		Message:           err.Error(),
		InternalErrorCode: internalErrorCode,
	}
}
