package models

type (
	model interface {
		Validate() error
		MarshalJSON() ([]byte, error)
		UnmarshalJSON([]byte) error
	}
)
