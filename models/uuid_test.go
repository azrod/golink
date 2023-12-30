package models

import (
	"testing"
)

func TestIsValidUUID(t *testing.T) {
	testCases := []struct {
		uuid     string
		expected bool
	}{
		{
			uuid:     "123e4567-e89b-12d3-a456-426614174000",
			expected: true,
		},
		{
			uuid:     "123e4567-e89b-12d3-a456-42661417400",
			expected: false,
		},
		{
			uuid:     "123e4567-e89b-12d3-a456-4266141740000",
			expected: false,
		},
		{
			uuid:     "123e4567-e89b-12dfd3-a456-426614174000",
			expected: false,
		},
	}

	for _, tc := range testCases {
		actual := isValidUUID(tc.uuid)
		if actual != tc.expected {
			t.Errorf("Expected IsValidUUID(%s) to return %v, but got %v", tc.uuid, tc.expected, actual)
		}
	}
}
