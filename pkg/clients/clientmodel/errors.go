package clientmodel

import "errors"

var (
	ErrKeyNotFound        = errors.New("key not found")
	ErrSourcePathNotFound = errors.New("source path not found")
	// ErrLabelNameNotFound  = errors.New("label name not found").
	ErrGroupNameNotFound = errors.New("group name not found")
	ErrLinkNameNotFound  = errors.New("link name not found")
)
