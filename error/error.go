package appError

import (
	"fmt"
	"net/http"
)

type AppError interface {
	Error() string
	Is(target error) bool
	HttpStatus() int
}

type ErrorInvalidArguments struct {
	Key string
	Err error
}

func (err ErrorInvalidArguments) Error() string {
	return fmt.Sprintf("[%v] invalid arguments: %v", err.Key, err.Err.Error())
}
func (err ErrorInvalidArguments) Is(target error) bool {
	_, ok := target.(ErrorInvalidArguments)
	return ok
}
func (err ErrorInvalidArguments) HttpStatus() int {
	return http.StatusBadRequest
}

type ErrorAlreadyExists struct {
	Key string
	Err error
}

func (err ErrorAlreadyExists) Error() string {
	return fmt.Sprintf("[%v] already exists: %v", err.Key, err.Err.Error())
}
func (err ErrorAlreadyExists) Is(target error) bool {
	_, ok := target.(ErrorAlreadyExists)
	return ok
}
func (err ErrorAlreadyExists) HttpStatus() int {
	return http.StatusConflict
}

type ErrorNotFound struct {
	Key string
	Err error
}

func (err ErrorNotFound) Error() string {
	return fmt.Sprintf("[%v] not found: %v", err.Key, err.Err.Error())
}
func (err ErrorNotFound) Is(target error) bool {
	_, ok := target.(ErrorNotFound)
	return ok
}
func (err ErrorNotFound) HttpStatus() int {
	return http.StatusNotFound
}
