package model

import "net/http"

type ControllerError struct {
	StatusCode int    `json:"statusCode,omitempty"`
	Status     string `json:"status,omitempty"`
	Error      string `json:"error,omitempty"`
}

func NewError() *ControllerError {
	return &ControllerError{}
}

func (e *ControllerError) BadRequest(err error) *ControllerError {
	e.StatusCode = http.StatusBadRequest
	e.Status = http.StatusText(http.StatusBadRequest)
	e.Error = err.Error()

	return e
}

func (e *ControllerError) InternalServer(err error) *ControllerError {
	e.StatusCode = http.StatusInternalServerError
	e.Status = http.StatusText(http.StatusInternalServerError)
	e.Error = err.Error()

	return e
}
