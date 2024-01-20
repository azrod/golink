package apiv1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/azrod/golink/api/model"
	"github.com/azrod/golink/models"
)

// getVersion godoc
// @Summary get the version of the application
// @Description get the version of the application
// @Tags system
// @Produce  json
// @Success 200 {object} string
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @Router /version [get].
func (h *v1) getVersion(c echo.Context) error {
	version, ok := c.Get("version").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.list", models.APIResponseError500{
			Code:    http.StatusInternalServerError,
			Message: "unable to get version",
		}))
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse(c, "version", version))
}
