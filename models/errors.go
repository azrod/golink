package models

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalid       = errors.New("invalid")
	ErrIsEmpty       = errors.New("is empty")
	ErrAlreadyExists = errors.New("already exists")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
