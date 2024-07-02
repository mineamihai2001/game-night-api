package services

import (
	"fmt"
	"strconv"
)

type ServiceError struct {
	Message   string
	ErrorCode int
}

func (e *ServiceError) Error() string {
	return e.Message
}

func (e *ServiceError) HttpStatus() int {
	hex := fmt.Sprintf("%X", e.ErrorCode)[0:3]

	value, err := strconv.ParseInt(hex, 16, 16)
	if err != nil {
		return 500
	}

	return int(value)
}

func (e *ServiceError) StringErrorCode() string {
	return fmt.Sprintf("%X", e.ErrorCode)
}

func NewServiceError(errorCode int, format string, v ...any) *ServiceError {
	return &ServiceError{
		Message:   fmt.Sprintf(format, v...),
		ErrorCode: errorCode,
	}
}

const (
	DocumentNotFound    = 0x190001
	EndpointNotFound    = 0x190002
	InternalServerError = 0x1F4001
	RepositoryError     = 0x1F4002
)
