package model

import (
	"errors"
	"fmt"
)

var ErrorIDAlreadyExist = errors.New("the id already exists")
var ErrorIDNotFound = errors.New("the id not exists")

type Error struct {
	err        error
	statusHTTP int
	data       interface{}
	apiMessage string
}

// NewError returns a new pointer Error
func NewError() *Error {
	return &Error{}
}

// Error implements the interface error
func (e *Error) Error() string {
	return fmt.Sprintf("Status: %d | Err: %v",
		e.statusHTTP, e.err)
}

// SetError sets the error field
func (e *Error) SetError(err error) { e.err = err }

// APIMessage gets the api message field
func (e *Error) APIMessage() string { return e.apiMessage }

// SetAPIMessage sets the api message field
func (e *Error) SetAPIMessage(message string) { e.apiMessage = message }

// StatusHTTP gets the status http field
func (e *Error) StatusHTTP() int { return e.statusHTTP }

// SetStatusHTTP sets the status http field
func (e *Error) SetStatusHTTP(status int) { e.statusHTTP = status }

// HasStatusHTTP returns true if the struct has the status http field
func (e *Error) HasStatusHTTP() bool { return e.statusHTTP != 0 }

// Data gets the data field
func (e *Error) Data() interface{} { return e.data }

// SetData sets the data field
func (e *Error) SetData(data interface{}) { e.data = data }

// HasData returns true if the struct has the data field
func (e *Error) HasData() bool { return e.data != nil }
