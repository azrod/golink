package golink

import (
	"context"
	"errors"

	"github.com/go-resty/resty/v2"

	"github.com/azrod/golink/models"
)

var version = "v0.0.1"

type (
	// Client is the golink client.
	Client struct {
		c         *resty.Client
		namespace string
	}
)

// New - Create new golink Client.
func New(apiURL string, debug bool, namespace string) *Client {
	clientResty := resty.New().
		SetBaseURL(apiURL + "/api/v1").
		SetDebug(debug).
		SetHeaders(map[string]string{
			"User-Agent": "golinksdk/go:" + version,
		})

	return &Client{
		c:         clientResty,
		namespace: namespace,
	}
}

// SetNamespace - Set namespace for golink Client.
func (c *Client) SetNamespace(namespace string) {
	c.namespace = namespace
}

// GetVersion - Get the version of the application.
func (c Client) GetVersion(ctx context.Context) (version string, err error) {
	r, err := c.c.R().
		SetContext(ctx).
		SetResult(models.APIResponse[string]{}).
		SetError(models.APIResponseError[models.APIResponseError400]{}).
		Get("/version")
	if err != nil {
		return
	}

	if r.IsError() {
		return version, errors.New(r.Error().(*models.APIResponseError[models.APIResponseError400]).Error.Message)
	}

	return r.Result().(*models.APIResponse[string]).Data, nil
}
