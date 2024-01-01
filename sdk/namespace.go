package golink

import (
	"context"
	"errors"

	"github.com/azrod/golink/models"
)

// GetNamespaces gets a list of namespaces from the database.
func (c Client) GetNamespaces(ctx context.Context) (namespaces []models.Namespace, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[[]models.Namespace]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Get("/namespaces")
	if err != nil {
		return []models.Namespace{}, err
	}

	if r.IsError() {
		return []models.Namespace{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[[]models.Namespace]).Data, nil
}

// GetNamespace gets a namespace from the database.
func (c Client) GetNamespace(ctx context.Context, namespace string) (ns models.Namespace, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[models.Namespace]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetPathParam("namespace", namespace).
		Get("/namespace/{namespace}")
	if err != nil {
		return models.Namespace{}, err
	}

	if r.IsError() {
		return models.Namespace{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Namespace]).Data, nil
}

// DeleteNamespace deletes a namespace from the database.
func (c Client) DeleteNamespace(ctx context.Context, namespace string) (err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetPathParam("namespace", namespace).
		Delete("/namespace/{namespace}")
	if err != nil {
		return err
	}

	if r.IsError() {
		return errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return nil
}

// CreateNamespace creates a namespace in the database.
func (c Client) CreateNamespace(ctx context.Context, namespace string) (ns models.Namespace, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[models.Namespace]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		SetBody(models.NamespaceRequest{Name: namespace}).
		Post("/namespace")
	if err != nil {
		return models.Namespace{}, err
	}

	if r.IsError() {
		return models.Namespace{}, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[models.Namespace]).Data, nil
}
