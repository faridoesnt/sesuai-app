package ahttp

import (
	"Sesuai/internal/api/constants"
	"Sesuai/internal/api/models"
	"Sesuai/pkg/alog"
	"net/http"
)

// Standard Errors for Mobile
func ErrInvalid(message string) Error {
	return Error{
		Code:    http.StatusOK,
		Status:  "invalid",
		Message: message,
		Data:    "",
	}
}

func ErrDenied(message string) Error {
	return Error{
		Code:    http.StatusOK,
		Status:  "denied",
		Message: message,
		Data:    "",
	}
}

func ErrFailure(message string) Error {
	return Error{
		Code:    http.StatusOK,
		Status:  "fail",
		Message: message,
		Data:    "",
	}
}

func ErrNotFound(message string) Error {
	return Error{
		Code:    http.StatusOK,
		Status:  "not_found",
		Message: message,
		Data:    "",
	}
}

// Standard Errors
var ErrBadRequest = Error{
	Code:    http.StatusBadRequest,
	Status:  "400",
	Message: "Bad Request",
	Data:    "",
}
var ErrUnauthorized = Error{
	Code:    http.StatusUnauthorized,
	Status:  "401",
	Message: "Unauthorized",
	Data:    "",
}
var ErrForbidden = Error{
	Code:    http.StatusForbidden,
	Status:  "403",
	Message: "Forbidden",
	Data:    "",
}
var ErrMethodNotAllowed = Error{
	Code:    http.StatusMethodNotAllowed,
	Status:  "405",
	Message: "Method Not Allowed",
	Data:    "",
}
var ErrConflict = Error{
	Code:    http.StatusConflict,
	Status:  "409",
	Message: "Conflict",
	Data:    "",
}
var ErrUnprocessableEntity = Error{
	Code:    http.StatusUnprocessableEntity,
	Status:  "422",
	Message: "Unprocessable Entity",
	Data:    "",
}
var ErrInternalServer = Error{
	Code:    http.StatusInternalServerError,
	Status:  "500",
	Message: "Internal Error",
	Data:    "",
}

// Error represent standard API error that contains HTTP Status (Status) and API-scoped Error Code (Code).
type Error struct {
	Code    int         `yaml:"code" json:"-"`
	Status  string      `yaml:"status" json:"status"`
	Message string      `yaml:"message" json:"message"`
	Data    interface{} `yaml:"data" json:"data"`
}

// Error is an implementation of built-in error type interface
func (e Error) Error() string {
	return e.Message
}

// CastError cast error interface as an Error
func CastError(err error, headers *models.Headers) Error {
	apiErr, ok := err.(Error)
	if !ok {
		alog.Logger.Errorf("internal error occurred: %s", err)
		// If assert type fail, create new internal error
		apiErr = ErrInternalServer

		if headers.OS == constants.Android {
			apiErr.Data = nil
		}
	}

	if headers.OS == constants.Android {
		apiErr.Data = nil
	}

	return apiErr
}
