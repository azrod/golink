package model

import (
	"github.com/labstack/echo/v4"

	"github.com/azrod/golink/models"
	"github.com/azrod/golink/pkg/clients/clientmodel"
)

type (
	HandlerFunc echo.HandlerFunc
	Handlers    struct {
		DB         clientmodel.ClientDB
		EchoServer *echo.Echo
	}
	HandlerAPIVersion struct {
		DB        clientmodel.ClientDB
		EchoGroup *echo.Group
	}
)

func GenParamsList(c echo.Context) map[string]string {
	params := make(map[string]string)
	for _, param := range c.ParamNames() {
		params[param] = c.Param(param)
	}

	return params
}

func GetRequestID(c echo.Context) string {
	return c.Get("requestID").(string)
}

func NewAPIResponse[T any](c echo.Context, method string, data T) models.APIResponse[T] {
	return models.APIResponse[T]{
		ID:     GetRequestID(c),
		Method: method,
		Params: GenParamsList(c),
		Data:   data,
	}
}

func NewAPIResponseError[I any](c echo.Context, method string, err I) models.APIResponseError[I] {
	return models.APIResponseError[I]{
		ID:     GetRequestID(c),
		Method: method,
		Params: GenParamsList(c),
		Error:  err,
	}
}
