package golink

import "github.com/go-resty/resty/v2"

var version = "v0.0.1"

type (
	// Client is the golink client.
	Client struct {
		c         *resty.Client
		namespace string
	}
)

// New - Create new golink Client.
func New(apiURL string, debug *bool, namespace *string) *Client {
	clientResty := resty.New().
		SetBaseURL(apiURL + "/api/v1").
		SetDebug(*debug).
		SetHeaders(map[string]string{
			"User-Agent": "golinksdk/go:" + version,
		})

	return &Client{
		c:         clientResty,
		namespace: *namespace,
	}
}

// SetNamespace - Set namespace for golink Client.
func (c *Client) SetNamespace(namespace string) {
	c.namespace = namespace
}
