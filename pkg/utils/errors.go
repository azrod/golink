package utils

import "strings"

// IsNotFoundError checks if the given error is a not found error.
func IsNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "not found")
}
