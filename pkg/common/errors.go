package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	ErrBadRequest          = "bad request"
	ErrUserIdAlreadyExists = "user with given user-id already exists"
	ErrUnauthorized        = "Unauthorized"
	ErrUnprocessableEntity = "unprocessable entity"
)

var (
	StatusUnprocessableEntity = errors.New("status unprocessable entity")
	StatusInternalServerError = errors.New("status internal server error")
)

type HttpError interface {
	Status() int
	Error() string
}

type OpaqueHttpError struct {
	ErrStatus int		`json:"status"`
	ErrError  string	`json:"error"`
}

func (err *OpaqueHttpError) Status() int {
	return err.ErrStatus
}

func (err *OpaqueHttpError) Error() string {
	return fmt.Sprintf("status: %d, errors: %s", err.ErrStatus, err.ErrError)
}


func NewStatusUnprocessableEntityError() HttpError {
	return &OpaqueHttpError {
		ErrStatus: http.StatusUnprocessableEntity,
		ErrError: StatusUnprocessableEntity.Error(),
	}
}

func NewStatusInternalServerError() HttpError {
	return &OpaqueHttpError{
		ErrStatus: http.StatusInternalServerError,
		ErrError: StatusInternalServerError.Error(),
	}
}

func ParseErrors(err error) HttpError {
	switch {
	case strings.Contains(err.Error(), ErrBadRequest):
		return &OpaqueHttpError{http.StatusBadRequest, err.Error() }
	case strings.Contains(err.Error(), ErrUserIdAlreadyExists):
		return &OpaqueHttpError{http.StatusBadRequest, err.Error() }
	case strings.Contains(err.Error(), ErrUnauthorized):
		return &OpaqueHttpError{http.StatusUnauthorized, err.Error() }
	case strings.Contains(err.Error(), ErrUnprocessableEntity):
		return NewStatusUnprocessableEntityError()
	default:
		if httpError, ok := err.(HttpError); ok {
			return httpError
		}
		return NewStatusInternalServerError()
	}
}