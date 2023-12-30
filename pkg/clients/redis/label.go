package redis

import (
	"context"
	"fmt"
	"strings"

	"github.com/azrod/golink/models"
)

// GetLabelByID gets a label from the database.
func (c Client) GetLabelByID(ctx context.Context, labelID models.LabelID) (models.Label, error) {
	if !strings.Contains(labelID.String(), labelKey) {
		labelID = models.LabelID(labelKey + labelID.String())
	}

	v, err := c.c.Get(ctx, labelID.String()).Result()
	if err != nil {
		return models.Label{}, err
	}

	// v is JSON string
	var label models.Label
	if err := label.UnmarshalJSON([]byte(v)); err != nil {
		return models.Label{}, err
	}

	return label, nil
}

// GetLabelByName gets a label from the database.
func (c Client) GetLabelByName(ctx context.Context, name string) (models.Label, error) {
	iter := c.c.Scan(ctx, 0, labelKey+"*", 0).Iterator()
	for iter.Next(ctx) {
		label, err := c.GetLabelByID(ctx, models.LabelID(iter.Val()))
		if err != nil {
			continue
		}

		if label.Name == name {
			return label, nil
		}
	}

	return models.Label{}, fmt.Errorf("label name %s is %w", name, models.ErrNotFound)
}

// UpdateLabel updates a label in the database.
func (c Client) UpdateLabel(ctx context.Context, label models.Label) error {
	b, err := label.MarshalJSON()
	if err != nil {
		return err
	}

	return c.c.Set(ctx, labelKey+label.ID.String(), string(b), 0).Err()
}

// DeleteLabel deletes a label from the database.
func (c Client) DeleteLabel(ctx context.Context, labelID models.LabelID) error {
	return c.c.Del(ctx, labelKey+labelID.String()).Err()
}

// CreateLabel creates a label in the database.
func (c Client) CreateLabel(ctx context.Context, label models.Label) error {
	b, err := label.MarshalJSON()
	if err != nil {
		return err
	}

	return c.c.Set(ctx, labelKey+label.ID.String(), string(b), 0).Err()
}

// ListLabels lists all labels in the database.
func (c Client) ListLabels(ctx context.Context) ([]models.Label, error) {
	labels := make([]models.Label, 0)
	iter := c.c.Scan(ctx, 0, labelKey+"*", 0).Iterator()
	for iter.Next(ctx) {
		label, err := c.GetLabelByID(ctx, models.LabelID(iter.Val()))
		if err != nil {
			continue
		}

		labels = append(labels, label)
	}

	return labels, nil
}

// ListLinksByLabel lists all links with a label in the database.
func (c Client) ListLinksByLabel(ctx context.Context, labelID models.LabelID) ([]models.Link, error) {
	var links []models.Link
	iter := c.c.Scan(ctx, 0, linkKey+"*", 0).Iterator()
	for iter.Next(ctx) {
		// itel.Val() is JSON string
		var link models.Link
		if err := link.UnmarshalJSON([]byte(iter.Val())); err != nil {
			continue
		}

		for _, l := range link.Labels {
			if l == labelID {
				links = append(links, link)
			}
		}
	}

	return links, nil
}
