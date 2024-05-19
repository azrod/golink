package golink

import (
	"context"
	"errors"

	"github.com/azrod/golink/models"
)

// GetLinkByID gets a link from the database.
func (c Client) GetLinkByID(ctx context.Context, linkID models.LinkID) (link models.Link, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"id":        linkID.String(),
			"namespace": c.namespace,
		}).
		SetResult(models.APIResponse[models.Link]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Get("/namespace/{namespace}/link/{id}")
	if err != nil {
		return models.Link{}, err
	}

	if r.IsError() {
		return models.Link{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Link]).Data, nil
}

// GetLinkByNames gets a link from the database.
func (c Client) GetLinkByName(ctx context.Context, name string) (link models.Link, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"name":      name,
			"namespace": c.namespace,
		}).
		SetResult(models.APIResponse[models.Link]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Get("/namespace/{namespace}/link/name/{name}")
	if err != nil {
		return models.Link{}, err
	}

	if r.IsError() {
		return models.Link{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Link]).Data, nil
}

// GetLinks gets a list of links from the database.
func (c Client) GetLinks(ctx context.Context) (links []models.Link, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetPathParam("namespace", c.namespace).
		SetResult(models.APIResponse[[]models.Link]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Get("/namespace/{namespace}/links")
	if err != nil {
		return []models.Link{}, err
	}

	if r.IsError() {
		return []models.Link{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[[]models.Link]).Data, nil
}

// GetLinksAllNamespace gets a list of links from all namespaces.
func (c Client) GetLinksAllNamespace(ctx context.Context) (links []models.Link, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[[]models.Link]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Get("/namespaces/links")
	if err != nil {
		return []models.Link{}, err
	}

	if r.IsError() {
		return []models.Link{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[[]models.Link]).Data, nil
}

// UpdateLink updates a link in the database.
func (c Client) UpdateLink(ctx context.Context, link models.LinkRequest, linkID string) error {
	r, err := c.c.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"id":        linkID,
			"namespace": c.namespace,
		}).
		SetResult(models.APIResponse[models.Link]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetBody(link).
		Put("/namespace/{namespace}/link/{linkID}")
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return nil
}

// DeleteLink deletes a link from the database. Is a alias for DeleteLinkByID.
func (c Client) DeleteLink(ctx context.Context, linkID models.LinkID) error {
	return c.DeleteLinkByID(ctx, linkID)
}

// DeleteLinkByID deletes a link from the database.
func (c Client) DeleteLinkByID(ctx context.Context, linkID models.LinkID) error {
	r, err := c.c.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"id":        linkID.String(),
			"namespace": c.namespace,
		}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Delete("/namespace/{namespace}/link/{id}")
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return nil
}

// DeleteLinkByName deletes a link from the database.
func (c Client) DeleteLinkByName(ctx context.Context, name string) error {
	r, err := c.c.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"name":      name,
			"namespace": c.namespace,
		}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Delete("/namespace/{namespace}/link/name/{name}")
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return nil
}

// CreateLink creates a link in the database.
func (c Client) CreateLink(ctx context.Context, link models.LinkRequest) (models.Link, error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetBody(link).
		SetPathParam("namespace", c.namespace).
		SetResult(models.APIResponse[models.Link]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Post("/namespace/{namespace}/link")
	if err != nil {
		return models.Link{}, err
	}

	if r.IsError() {
		return models.Link{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Link]).Data, nil
}

// GetLinksAssociatedToLabel gets a list of links associated to a label from the database.
func (c Client) GetLinksAssociatedToLabel(ctx context.Context, labelID models.LabelID) (links []models.Link, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[[]models.Link]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetPathParam("id", labelID.String()).
		Get("/label/{id}/links")
	if err != nil {
		return []models.Link{}, err
	}

	if r.IsError() {
		return []models.Link{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[[]models.Link]).Data, nil
}
