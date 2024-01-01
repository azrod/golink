package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

var _ model = (*Namespace)(nil)

type (
	NamespaceRequest struct {
		// Name is the name of the namespace
		Name string `json:"name"`
	}
	// Namespace is the model for the namespace.
	Namespace struct {
		NamespaceRequest

		// Links is the list of links associated with the namespace
		Links []Link `json:"links"`

		// Enabled is the status of the namespace
		Enabled Enabled `json:"enabled"`
	}
)

// UnmarshalJSON unmarshals the JSON value into a NamespaceID.
func (g *Namespace) UnmarshalJSON(b []byte) error {
	var x struct {
		Name    string `json:"name"`
		Links   []Link `json:"links"`
		Enabled bool   `json:"enabled"`
	}

	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	g.Name = x.Name
	g.Links = x.Links
	g.Enabled = Enabled(x.Enabled)

	return nil
}

// Validate validates the namespace request.
func (g *NamespaceRequest) Validate() error {
	// Check if Name is empty
	if g.Name == "" {
		return fmt.Errorf("Validate: The namespace name %s is %w", g.Name, ErrIsEmpty)
	}

	// Set Name to lowercase
	g.Name = strings.ToLower(g.Name)

	return nil
}

// Validate validates the namespace.
func (g *Namespace) Validate() error {
	if err := g.NamespaceRequest.Validate(); err != nil {
		return err
	}

	g.Enabled = Enabled(true)

	return nil
}

// MarshalJSON marshal the JSON value.
func (g *Namespace) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name    string `json:"name"`
		Links   []Link `json:"links"`
		Enabled bool   `json:"enabled"`
	}{
		Name:    g.Name,
		Links:   g.Links,
		Enabled: bool(g.Enabled),
	})
}
