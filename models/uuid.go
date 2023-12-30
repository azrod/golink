package models

import "regexp"

// isValidUUID checks if the given string is a valid UUID.
func isValidUUID(uuid string) bool {
	return regexp.MustCompile(`(?m)^\w{8}-\w{4}-\w{4}-\w{4}-\w{12}$`).MatchString(uuid)
}

// IsValidUUID checks if the given string is a valid UUID.
func IsValidUUID(uuid string) bool {
	return isValidUUID(uuid)
}
