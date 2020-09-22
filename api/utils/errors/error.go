package errors

import "net/http"

// APIError is the error model for the services
type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	APIStatus  int    `json:"status"`
	APIMessage string `json:"message"`
	APIError   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.APIStatus
}

func (e *apiError) Message() string {
	return e.APIMessage
}

func (e *apiError) Error() string {
	return e.APIError
}

// CreateError creates a error
func CreateError(statusCode int, message string) APIError {
	return &apiError{
		APIStatus:  statusCode,
		APIMessage: message,
	}
}

// CreateNotFoundError creates a not found error
func CreateNotFoundError(message string) APIError {
	return &apiError{
		APIStatus:  http.StatusNotFound,
		APIMessage: message,
	}
}

// CreateBadRequestError creates a bad request error
func CreateBadRequestError(message string) APIError {
	return &apiError{
		APIStatus:  http.StatusBadRequest,
		APIMessage: message,
	}
}

// CreateInternalServerError creates a internal server error
func CreateInternalServerError(message string) APIError {
	return &apiError{
		APIStatus:  http.StatusInternalServerError,
		APIMessage: message,
	}
}
