package kvmodel

import "errors"

var (
	ErrEmptyKey    = errors.New("key is empty")
	ErrNotFound    = errors.New("key not found")
	ErrEmptyPrefix = errors.New("prefix is empty")
)
