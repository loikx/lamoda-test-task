package errors

import (
	"encoding/json"
)

type Error struct {
	Err error
}

func NewError(err error) *Error {
	return &Error{Err: err}
}

func (e *Error) MarshalJSON() ([]byte, error) {
	dto := struct {
		Error string `json:"error"`
	}{
		e.Err.Error(),
	}

	return json.Marshal(dto)
}
