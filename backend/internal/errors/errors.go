package errors

import (
	"net/http"
)

type ErrorType string

const (
	ErrorTypeFail  ErrorType = "fail"
	ErrorTypeError ErrorType = "error"
)

// Client-facing error messages
const (
	MsgBadRequest    = "Invalid request. Please check your input and try again."
	MsgUnauthorized  = "Authentication required. Please log in and try again."
	MsgNotFound      = "The requested resource was not found."
	MsgInternalError = "Something went wrong. Please try again or contact support if the issue persists."
)

// APIError represents API errors with HTTP status codes
type APIError struct {
	Code          int       `json:"code"`
	internalMsg   string    // Internal error details for logging only
	ClientMessage string    `json:"message,omitempty"` // Safe message for client response
	Type          ErrorType `json:"type"`
	Data          any       `json:"data,omitempty"`
}

func (e *APIError) Error() string {
	return e.internalMsg
}

func BadRequest(message string, data map[string]any) *APIError {
	return &APIError{
		Code:          http.StatusBadRequest,
		internalMsg:   message,
		ClientMessage: MsgBadRequest,
		Type:          ErrorTypeFail,
		Data:          data,
	}
}

func Unauthorized(message string, data any) *APIError {
	return &APIError{
		Code:          http.StatusUnauthorized,
		internalMsg:   message,
		ClientMessage: MsgUnauthorized,
		Type:          ErrorTypeFail,
		Data:          data,
	}
}

func NotFound(message string, data any) *APIError {
	return &APIError{
		Code:          http.StatusNotFound,
		internalMsg:   message,
		ClientMessage: MsgNotFound,
		Type:          ErrorTypeFail,
		Data:          data,
	}
}

func InternalError(message string) *APIError {
	return &APIError{
		Code:          http.StatusInternalServerError,
		internalMsg:   message,
		ClientMessage: MsgInternalError,
		Type:          ErrorTypeError,
	}
}

func GetAPIError(err error) *APIError {
	if err == nil {
		return nil
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr
	}
	return InternalError(err.Error())
}
