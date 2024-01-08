package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var _ Model = (*Label)(nil)

type (
	LabelRequest struct {
		// Name is the name of the label
		Name string `json:"name" example:"example" validate:"required"`

		// Color is the color of the label
		// If color is empty, set default color
		Color ColorName `json:"color,omitempty" example:"red" default:"gray" validate:"optional"`
	}

	// Label is the model for the label.
	Label struct {
		// ID is the unique identifier of the label
		ID LabelID `json:"id,omitempty" example:"f7c5c4a0-6a1f-4d4b-8a9f-5b5d6d9b8c4d"`
		LabelRequest
	}
	LabelID string
)

// String returns the string representation of the label ID.
func (l LabelID) String() string {
	return string(l)
}

// UnmarshalJSON unmarshals the JSON value into a LabelID.
func (l *Label) UnmarshalJSON(b []byte) error {
	var x struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}

	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	l.ID = LabelID(x.ID)
	l.Name = x.Name
	l.Color = ColorName(x.Color)

	return nil
}

// Validate validates the label.
// If ID is invalid the ErrInvalid is returned.
// If Name is empty the ErrIsEmpty is returned.
// If Color is not found in table the ErrNotFound is returned.
func (l *Label) Validate() error {
	// check if ID is empty
	if l.ID == "" {
		// if ID is empty generate a new UUID
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		l.ID = LabelID(id.String())
	}

	// check if ID is valid UUID
	if !isValidUUID(l.ID.String()) {
		return fmt.Errorf("Validate: The label ID %s is %w", l.ID, ErrInvalid)
	}

	// check if Name is empty
	if l.Name == "" {
		return fmt.Errorf("Validate: The label name %s is %w", l.Name, ErrIsEmpty)
	}

	// force lowercase name
	l.Name = strings.ToLower(l.Name)

	// check if Color is empty. If empty, set default color
	if l.Color == "" {
		l.Color = ColorDefault
	}

	// check if Color is valid
	if ok, err := IsValidColor(l.Color); !ok {
		return err
	}

	return nil
}

// MarshalJSON marshal the JSON value.
func (l Label) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Color string `json:"color"`
	}{
		ID:    l.ID.String(),
		Name:  strings.ToLower(l.Name),
		Color: l.Color.String(),
	})
}
