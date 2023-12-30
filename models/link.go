package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

var _ model = (*Link)(nil)

type (
	// LinkRequest is the model for the link request.
	LinkRequest struct {
		// Name is the name of the link
		Name string `json:"name" example:"example" validate:"required"`

		// SourcePath is the path of the link
		SourcePath string `json:"sourcePath" example:"/example" validate:"required"`

		// TargetURL is the original URL
		TargetURL string `json:"targetUrl" example:"https://example.com" validate:"required"`

		// Labels is the list of labels associated with the link
		Labels []LabelID `json:"labels,omitempty" example:"a54abdd8-776c-4982-bba7-34caa58596c4,fcb0726f-9701-492c-afe7-aa7121afc9cf"`

		// Enabled is the status of the link
		Enabled Enabled `json:"enabled" example:"true" validate:"required"`
	}

	// Link is the model for the shortened link.
	Link struct {
		// ID is the unique identifier of the link in the database
		ID LinkID `json:"id" example:"84214e93-437e-434d-96c2-22c3a63b3c67"`

		// CreatedAt is the time the link was created - RFC3339
		CreatedAt string `json:"createdAt" example:"2021-07-04T15:04:05.999999999Z"`

		// UpdatedAt is the time the link was last modified - RFC3339
		UpdatedAt string `json:"lastModified" example:"2021-07-04T15:04:05.999999999Z"`

		// Namespace is the group the link belongs to
		NameSpace string `json:"namespace" example:"prod"`

		LinkRequest
	}
	LinkID string
)

// String returns the string representation of the link ID.
func (l LinkID) String() string {
	return string(l)
}

// Validate validates the link.
func (l *Link) Validate() error {
	// check if ID is empty
	if l.ID == "" {
		// if ID is empty generate a new UUID
		id, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		l.ID = LinkID(id.String())
	}

	// Check if ID is UUIDv4 format
	if !isValidUUID(l.ID.String()) {
		return fmt.Errorf("Validate: The link ID %s is %w", l.ID, ErrInvalid)
	}

	// Check if SourcePath is empty
	if l.SourcePath == "" {
		return fmt.Errorf("Validate: The link source path %s is %w", l.SourcePath, ErrIsEmpty)
	}

	// Check if SourcePath have slash at the beginning
	if l.SourcePath[0] != '/' {
		return fmt.Errorf("Validate: The link source path %s is %w, must start with a slash", l.SourcePath, ErrInvalid)
	}

	// Check if TargetURL is empty
	if l.TargetURL == "" {
		return fmt.Errorf("Validate: The link target URL %s is %w", l.TargetURL, ErrIsEmpty)
	}

	// Check if TargetURL is valid
	if _, err := url.ParseRequestURI(l.TargetURL); err != nil {
		return fmt.Errorf("Validate: The link target URL %s have %w url format - %w", l.TargetURL, ErrInvalid, err)
	}

	// Check if Name is empty
	if l.Name == "" {
		return fmt.Errorf("Validate: The link name %s is %w", l.Name, ErrIsEmpty)
	}

	// Set Name to lowercase
	l.Name = strings.ToLower(l.Name)

	switch l.NameSpace == "" {
	case true:
		// Set default namespace if empty
		l.NameSpace = "default"
	case false:
		// Set NameSpace to lowercase
		l.NameSpace = strings.ToLower(l.NameSpace)
	}

	// If CreatedAt is empty, set current time
	if l.CreatedAt == "" {
		l.CreatedAt = time.Now().UTC().Format(time.RFC3339)
	}

	return nil
}

// UnmarshalJSON unmarshals the JSON value into a LinkID.
func (l *Link) UnmarshalJSON(b []byte) error {
	var x struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		SourcePath string   `json:"sourcePath"`
		TargetURL  string   `json:"targetUrl"`
		Enabled    bool     `json:"enabled"`
		CreatedAt  string   `json:"createdAt,omitempty"`
		UpdatedAt  string   `json:"lastModified,omitempty"`
		Labels     []string `json:"labels,omitempty"`
		NameSpace  string   `json:"namespace,omitempty"`
	}

	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	labels := make([]LabelID, len(x.Labels))
	for i, label := range x.Labels {
		labels[i] = LabelID(label)
	}

	l.ID = LinkID(x.ID)
	l.Name = x.Name
	l.SourcePath = x.SourcePath
	l.TargetURL = x.TargetURL
	l.Enabled = Enabled(x.Enabled)
	l.CreatedAt = x.CreatedAt
	l.UpdatedAt = x.UpdatedAt
	l.Labels = labels
	l.NameSpace = x.NameSpace

	return nil
}

// MarshalJSON marshal the JSON value.
func (l *Link) MarshalJSON() ([]byte, error) {
	labels := make([]string, len(l.Labels))
	for i, label := range l.Labels {
		labels[i] = label.String()
	}

	if err := l.Validate(); err != nil {
		return nil, err
	}

	return json.Marshal(struct {
		ID         string   `json:"id"`
		Name       string   `json:"name,omitempty"`
		SourcePath string   `json:"sourcePath"`
		TargetURL  string   `json:"targetUrl"`
		Enabled    bool     `json:"enabled"`
		CreatedAt  string   `json:"createdAt"`
		UpdatedAt  string   `json:"lastModified"`
		Labels     []string `json:"labels,omitempty"`
		NameSpace  string   `json:"namespace,omitempty"`
	}{
		ID:         l.ID.String(),
		Name:       strings.ToLower(l.Name),
		SourcePath: l.SourcePath,
		TargetURL:  l.TargetURL,
		Enabled:    bool(l.Enabled),
		CreatedAt:  l.CreatedAt,
		UpdatedAt:  l.UpdatedAt,
		Labels:     labels,
		NameSpace:  l.NameSpace,
	})
}
