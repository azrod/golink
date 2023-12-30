package apiv1

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/azrod/golink/api/model"
	"github.com/azrod/golink/models"
)

// listLabels
// @Summary List all labels
// @Description List all labels available.
// @Tags labels
// @Produce  json
// @Success 200 {array} models.APIResponse{data=[]models.Label}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @Router /labels [get].
func (h *v1) listLabels(c echo.Context) error {
	labels, err := h.DB.ListLabels(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.list", models.APIResponseError500{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}))
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[[]models.Label](
		c,
		"label.list",
		labels,
	))
}

// getLabelByID
// @Summary Get a label by ID
// @Description get label by ID
// @Tags labels
// @Produce  json
// @Param id path string true "Label ID"
// @Success 200 {object} models.APIResponse{data=models.Label}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /label/{id} [get].
func (h *v1) getLabelByID(c echo.Context) error {
	labelID := c.Param("id")
	if labelID == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.get.id", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: "labelID is empty",
		}))
	}

	label, err := h.DB.GetLabelByID(c.Request().Context(), models.LabelID(labelID))
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "label.get.id", models.APIResponseError404{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.get.id", models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.get.id", models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Label](
		c,
		"label.get",
		label,
	))
}

// getLabelByName
// @Summary Get a label by name
// @Description get label by name
// @Tags labels
// @Produce  json
// @Param name path string true "Label name"
// @Success 200 {object} models.APIResponse{data=models.Label}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /label/{name} [get].
func (h *v1) getLabelByName(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.get.name", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: "name is empty",
		}))
	}

	label, err := h.DB.GetLabelByName(c.Request().Context(), name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "label.get.name", models.APIResponseError404{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.get.name", models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.get.name", models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Label](
		c,
		"label.get",
		label,
	))
}

// createLabel
// @Summary Create label
// @Description create label
// @Tags labels
// @Accept  json
// @Produce  json
// @Param label body models.LabelRequest true "Label"
// @Success 201 {object} models.APIResponse{data=models.Label}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 409 {object} models.APIResponseError{error=models.APIResponseError409}
// @Router /label [post].
func (h *v1) createLabel(c echo.Context) error {
	var label models.Label
	if err := c.Bind(&label); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.create", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	if err := label.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.create", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	// Check if label already exists by name. if it does, return 409.
	if _, err := h.DB.GetLabelByName(c.Request().Context(), label.Name); err == nil {
		return c.JSON(http.StatusConflict, model.NewAPIResponseError(c, "label.create", models.APIResponseError409{
			Code:    http.StatusConflict,
			Message: "label already exists",
		}))
	}

	if err := h.DB.CreateLabel(c.Request().Context(), label); err != nil {
		switch {
		case errors.Is(err, models.ErrAlreadyExists):
			return c.JSON(http.StatusConflict, model.NewAPIResponseError(c, "label.create", models.APIResponseError409{
				Code:    http.StatusConflict,
				Message: err.Error(),
			}))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.create", models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.create", models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		}
	}

	return c.JSON(http.StatusCreated, model.NewAPIResponse[models.Label](
		c,
		"label.create",
		label,
	))
}

// updateLabel
// @Summary Update label
// @Description update label
// @Tags labels
// @Accept json
// @Produce json
// @Param label body models.Label true "Label"
// @Success 200 {object} models.APIResponse{data=models.Label}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @Router /label [put].
func (h *v1) updateLabel(c echo.Context) error {
	var label models.Label
	if err := c.Bind(&label); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.update", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	if err := label.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.update", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}))
	}

	if err := h.DB.UpdateLabel(c.Request().Context(), label); err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "label.update", models.APIResponseError404{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.update", models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.update", models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		}
	}

	return c.JSON(http.StatusOK, label)
}

// deleteLabel
// @Summary Delete label
// @Description delete label
// @Tags labels
// @Produce  json
// @Param id path string true "Label ID"
// @Success 200 {object} models.APIResponse{}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @Router /label/{id} [delete].
func (h *v1) deleteLabel(c echo.Context) error {
	labelID := c.Param("id")
	if labelID == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.delete", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: "labelID is empty",
		}))
	}

	if err := h.DB.DeleteLabel(c.Request().Context(), models.LabelID(labelID)); err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "label.update", models.APIResponseError404{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.update", models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[struct{}](
		c,
		"label.delete",
		struct{}{},
	))
}

// deleteLabelByName
// @Summary Delete label by name
// @Description delete label by name
// @Tags labels
// @Produce  json
// @Param name path string true "Label name"
// @Success 200 {object} models.APIResponse{}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @Router /label/name/{name} [delete].
func (h *v1) deleteLabelByName(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "label.delete.name", models.APIResponseError400{
			Code:    http.StatusBadRequest,
			Message: "name is empty",
		}))
	}

	// GetLabelByName
	label, err := h.DB.GetLabelByName(c.Request().Context(), name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "label.delete.name", models.APIResponseError404{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.delete.name", models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		}
	}

	// DeleteLabel
	if err := h.DB.DeleteLabel(c.Request().Context(), label.ID); err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "label.delete.name", models.APIResponseError404{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "label.delete.name", models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[struct{}](
		c,
		"label.delete.name",
		struct{}{},
	))
}
