package models

type (
	Model interface {
		Validate() error
		MarshalJSON() ([]byte, error)
		UnmarshalJSON([]byte) error
	}
)
