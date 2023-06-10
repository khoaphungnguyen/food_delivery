package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewAuthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "Something went wrong", err.Error(), "db_error")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "invalid_request")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotList%s", entity))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot Delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotDelete%s", entity))
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot Update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotUpdate%s", entity))
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot Get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotGet%s", entity))
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s already existed", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyexisted", entity))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, fmt.Sprintln("No Permission"),
		fmt.Sprintln("ErrNoPermissions"))
}

var RecordNotFound = errors.New("record not found")
