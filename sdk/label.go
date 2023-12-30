package golink

import (
	"context"
	"errors"

	"github.com/azrod/golink/models"
)

// GetLabels gets a list of labels from the database.
func (c Client) GetLabels(ctx context.Context) (labels []models.Label, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[[]models.Label]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Get("/labels")
	if err != nil {
		return []models.Label{}, err
	}

	if r.IsError() {
		return []models.Label{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[[]models.Label]).Data, nil
}

// GetLabelByID gets a label from the database.
func (c Client) GetLabelByID(ctx context.Context, labelID models.LabelID) (label models.Label, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[models.Label]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetPathParam("id", labelID.String()).
		Get("/label/{id}")
	if err != nil {
		return models.Label{}, err
	}

	if r.IsError() {
		return models.Label{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Label]).Data, nil
}

// GetLabelByName gets a label from the database.
func (c Client) GetLabelByName(ctx context.Context, name string) (label models.Label, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[models.Label]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetPathParam("name", name).
		Get("/label/name/{name}")
	if err != nil {
		return models.Label{}, err
	}

	if r.IsError() {
		return models.Label{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Label]).Data, nil
}

// DeleteLabelByID deletes a label from the database.
func (c Client) DeleteLabelByID(ctx context.Context, labelID models.LabelID) error {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[models.Label]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetPathParam("id", labelID.String()).
		Delete("/label/{id}")
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return nil
}

// DeleteLabelByName deletes a label from the database.
func (c Client) DeleteLabelByName(ctx context.Context, name string) error {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[models.Label]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetPathParam("name", name).
		Delete("/label/name/{name}")
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return nil
}

// AddLabel adds a label to the database.
func (c Client) AddLabel(ctx context.Context, label models.LabelRequest) (models.Label, error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetBody(label).
		SetResult(models.APIResponse[models.Label]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Post("/label")
	if err != nil {
		return models.Label{}, err
	}

	if r.IsError() {
		return models.Label{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Label]).Data, nil
}
