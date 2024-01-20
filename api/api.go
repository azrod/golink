package api

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/azrod/golink/api/model"
	apiv1 "github.com/azrod/golink/api/v1"
	"github.com/azrod/golink/pkg/sb"
)

// @title Golink API
// @version 1.0
// @contact.url http://github.com/azrod/golink
// @host go
// @BasePath /api/v1
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @description This is a API for Golink Server.
func NewHandlers(db sb.Client, e *echo.Echo) *model.Handlers {
	h := &model.Handlers{
		DB:         db,
		EchoServer: e,
	}
	api := h.EchoServer.Group("/api")
	api.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Create a new request ID
			id := uuid.New().String()
			// Add it to the context of the request
			c.Set("requestID", id)
			// Add it to the response header
			c.Response().Header().Set("X-Request-ID", id)
			c.Logger().SetHeader(id)
			return next(c)
		}
	}))
	apiv1.New(db, api)

	return h
}
