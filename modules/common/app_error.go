package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"statusCode"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
}

func NewAppError(statusCode int, root error, message string, log string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    message,
		Log:        log,
	}
}

func NewBadRequestError(root error, message string) *AppError {
	return NewAppError(http.StatusBadRequest, root, message, root.Error())
}

func NewNotFoundError(root error, message string) *AppError {
	return NewAppError(http.StatusNotFound, root, message, root.Error())
}

func NewServerError(root error, message string) *AppError {
	return NewAppError(http.StatusInternalServerError, root, message, root.Error())
}

func NewDbError(root error) *AppError {
	return NewAppError(http.StatusInternalServerError, root, "database error", root.Error())
}

func NewUnauthorizedError(root error, message string) *AppError {
	return NewAppError(http.StatusUnauthorized, root, message, root.Error())
}

func (e *AppError) RootError() error {
	var err *AppError
	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}
