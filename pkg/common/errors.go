package common

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	StatusUnprocessableEntity = errors.New("status unprocessable entity")
)

type HttpError interface {
	Error() string
}

type OpaqueHttpError struct {
	ErrStatus int		`json:"status"`
	ErrError  string	`json:"error"`
}

func (err OpaqueHttpError) Error() string {
	return fmt.Sprintf("status: %d, errors: %s", err.ErrStatus, err.ErrError)
}


func NewStatusUnprocessableEntityError() HttpError {
	return OpaqueHttpError {
		ErrStatus: http.StatusUnprocessableEntity,
		ErrError: StatusUnprocessableEntity.Error(),
	}
}