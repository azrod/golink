package sb

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/azrod/golink/models"
)

// GetLabelByID gets a label from the database.
func (c Client) GetLabelByID(ctx context.Context, labelID models.LabelID) (label models.Label, err error) {
	key := labelID.String()

	if !strings.Contains(labelID.String(), labelKey) {
		key = fmt.Sprintf("%s%s", labelKey, labelID.String())
	}

	if err := c.c.Get(ctx, key, &label); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.Label{}, fmt.Errorf("label %s %w", labelID, models.ErrNotFound)
		}
		return models.Label{}, err
	}

	return label, nil
}

// GetLabelByName gets a label from the database.
func (c Client) GetLabelByName(ctx context.Context, name string) (models.Label, error) {
	labels, err := c.ListLabels(ctx)
	if err != nil {
		return models.Label{}, err
	}

	for _, label := range labels {
		if label.Name == name {
			return label, nil
		}
	}

	return models.Label{}, fmt.Errorf("label name %s is %w", name, models.ErrNotFound)
}

// UpdateLabel updates a label in the database.
func (c Client) UpdateLabel(ctx context.Context, label models.Label) error {
	l, err := c.GetLabelByID(ctx, label.ID)
	if err != nil {
		return err
	}

	if err := label.Validate(); err != nil {
		return err
	}

	return c.c.Set(ctx, labelKey+l.ID.String(), &label)
}

// DeleteLabel deletes a label from the database.
func (c Client) DeleteLabel(ctx context.Context, labelID models.LabelID) error {
	return c.c.Delete(ctx, labelKey+labelID.String())
}

// CreateLabel creates a label in the database.
func (c Client) CreateLabel(ctx context.Context, label models.Label) error {
	if err := label.Validate(); err != nil {
		return err
	}

	return c.c.Set(ctx, labelKey+label.ID.String(), &label)
}

// ListLabels lists all labels in the database.
func (c Client) ListLabels(ctx context.Context) ([]models.Label, error) {
	labels := make([]models.Label, 0)
	keys, err := c.c.List(ctx, labelKey)
	if err != nil {
		return []models.Label{}, err
	}

	for _, key := range keys {
		label, err := c.GetLabelByID(ctx, models.LabelID(key))
		if err != nil {
			continue
		}

		labels = append(labels, label)
	}

	return labels, nil
}

// ListLinksByLabel lists all links with a label in the database.
func (c Client) ListLinksByLabel(_ context.Context, _ models.LabelID) ([]models.Link, error) {
	// TODO: implement
	return []models.Link{}, nil
}
