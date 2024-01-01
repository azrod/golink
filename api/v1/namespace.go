package apiv1

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/azrod/golink/api/model"
	"github.com/azrod/golink/models"
)

// listNamespaces
// @Summary List all Namespaces
// @Description get all Namespaces
// @Tags namespaces
// @Produce  json
// @Success 200 {array} models.APIResponse{data=[]models.Namespace}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @Router /namespaces [get].
func (h *v1) listNamespaces(c echo.Context) error {
	nss, err := h.DB.ListNamespaces(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
			c,
			"namespace.list",
			models.APIResponseError500{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			},
		))
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[[]models.Namespace](
		c,
		"namespace.list",
		nss,
	))
}

// getNamespace
// @Summary Get a Namespace by name
// @Description get Namespace by name
// @Tags namespaces
// @Produce  json
// @Param name path string true "Namespace"
// @Success 200 {object} models.APIResponse{data=models.Namespace}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name} [get].
func (h *v1) getNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.get.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "name is empty",
			},
		))
	}

	ns, err := h.DB.GetNamespace(c.Request().Context(), name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.get.name",
				models.APIResponseError404{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.get.name",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.get.name",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Namespace](
		c,
		"namespace.get.name",
		ns,
	))
}

// createNamespace
// @Summary Create Namespace
// @Description create Namespace
// @Tags namespaces
// @Accept  json
// @Produce  json
// @Param Namespace body models.NamespaceRequest true "Namespace"
// @Success 200 {object} models.APIResponse{data=models.Namespace}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 409 {object} models.APIResponseError{error=models.APIResponseError409}
// @Router /namespace [post].
func (h *v1) createNamespace(c echo.Context) error {
	var ns models.NamespaceRequest
	if err := c.Bind(&ns); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.create",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		))
	}

	if err := ns.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.create",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		))
	}

	nsCreated, err := h.DB.CreateNamespace(c.Request().Context(), ns)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.create",
				models.APIResponseError404{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.create",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrAlreadyExists):
			return c.JSON(http.StatusConflict, model.NewAPIResponseError(
				c,
				"namespace.create",
				models.APIResponseError409{
					Code:    http.StatusConflict,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.create",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Namespace](
		c,
		"namespace.create",
		nsCreated,
	))
}

// deleteNamespace
// @Summary Delete Namespace
// @Description delete Namespace
// @Tags namespaces
// @Produce json
// @Param name path string true "Namespace"
// @Success 200 {object} models.APIResponse{}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @Router /namespace/{name} [delete].
func (h *v1) deleteNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.delete",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	if err := h.DB.DeleteNamespace(c.Request().Context(), name); err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.delete",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.delete",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[struct{}](
		c,
		"namespace.delete",
		struct{}{},
	))
}

// listLinksByNamespace
// @Summary List all Links by Namespace
// @Description get all Links by Namespace
// @Tags namespaces
// @Produce  json
// @Param name path string true "Namespace"
// @Success 200 {array} models.APIResponse{data=[]models.Link}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/links [get].
func (h *v1) listLinksByNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.links.list",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	links, err := h.DB.ListLinks(c.Request().Context(), name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.links.list",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.links.list",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.links.list",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[[]models.Link](
		c,
		"namespace.links.list",
		links,
	))
}

// getLinkFromNamespace
// @Summary Get a Link by name from Namespace
// @Description get Link by name from Namespace
// @Tags namespaces
// @Produce  json
// @Param name path string true "Namespace"
// @Param linkID path string true "Link ID"
// @Success 200 {object} models.APIResponse{data=models.Link}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/link/{linkID} [get].
func (h *v1) getLinkFromNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.get.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	linkID := c.Param("linkID")
	if linkID == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.get.linkID",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "linkID is empty",
			},
		))
	}

	link, err := h.DB.GetLinkByID(c.Request().Context(), models.LinkID(linkID), name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
		c,
		"namespace.link.get",
		link,
	))
}

// getLinkByNameFromNamespace
// @Summary Get a Link by name from Namespace
// @Description get Link by name from Namespace
// @Tags namespaces
// @Produce  json
// @Param name path string true "Namespace"
// @Param linkName path string true "Link Name"
// @Success 200 {object} models.APIResponse{data=models.Link}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/link/name/{linkName} [get].
func (h *v1) getLinkByNameFromNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.get.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	linkName := c.Param("linkName")
	if linkName == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.get.linkName",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "linkName is empty",
			},
		))
	}

	link, err := h.DB.GetLinkByName(c.Request().Context(), linkName, name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
		c,
		"namespace.link.get",
		link,
	))
}

// getLinkByPathFromNamespace
// @Summary Get a Link by path from Namespace
// @Description get Link by path from Namespace
// @Tags namespaces
// @Produce  json
// @Param name path string true "Namespace"
// @Param path path string true "Link Path"
// @Success 200 {object} models.APIResponse{data=models.Link}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/link/path/{path} [get].
func (h *v1) getLinkByPathFromNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.get.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	path := c.Param("path")
	if path == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.get.path",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "path is empty",
			},
		))
	}

	link, err := h.DB.GetLinkByPath(c.Request().Context(), path, name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.get",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
		c,
		"namespace.link.get",
		link,
	))
}

// addLinkToNamespace
// @Summary Add a Link to Namespace
// @Description add a Link to Namespace
// @Tags namespaces
// @Accept  json
// @Produce  json
// @Param name path string true "Namespace"
// @Param Link body models.LinkRequest true "Link"
// @Success 200 {object} models.APIResponse{data=models.Link}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 409 {object} models.APIResponseError{error=models.APIResponseError409}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/link [post].
func (h *v1) addLinkToNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.add.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	var link models.LinkRequest
	if err := c.Bind(&link); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.add",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		))
	}

	// Check if namespace exists
	if _, err := h.DB.GetNamespace(c.Request().Context(), name); err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.add",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.add",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	// Check if sourcePath exists in namespace
	if _, err := h.DB.GetLinkByPath(c.Request().Context(), link.SourcePath, name); err == nil {
		return c.JSON(http.StatusConflict, model.NewAPIResponseError(
			c,
			"namespace.link.add",
			models.APIResponseError409{
				Code:    http.StatusConflict,
				Message: "sourcePath already exists in namespace",
			},
		))
	}

	l, err := h.DB.CreateLink(c.Request().Context(), link, name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.add",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.link.add",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrAlreadyExists):
			return c.JSON(http.StatusConflict, model.NewAPIResponseError(
				c,
				"namespace.link.add",
				models.APIResponseError409{
					Code:    http.StatusConflict,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.add",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
		c,
		"namespace.link.add",
		l,
	))
}

// updateLinkInNamespace
// @Summary Update a Link in Namespace
// @Description update a Link in Namespace
// @Tags namespaces
// @Accept  json
// @Produce  json
// @Param name path string true "Namespace"
// @Param linkID path string true "Link ID"
// @Param Link body models.LinkRequest true "Link"
// @Success 200 {object} models.APIResponse{data=models.Link}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 409 {object} models.APIResponseError{error=models.APIResponseError409}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/link/{linkID} [put].
func (h *v1) updateLinkInNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.update.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	linkID := c.Param("linkID")
	if linkID == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.update.linkID",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "linkID is empty",
			},
		))
	}

	var link models.LinkRequest
	if err := c.Bind(&link); err != nil {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.update",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			},
		))
	}

	l, err := h.DB.UpdateLink(c.Request().Context(), link, models.LinkID(linkID), name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.update",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.link.update",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrAlreadyExists):
			return c.JSON(http.StatusConflict, model.NewAPIResponseError(
				c,
				"namespace.link.update",
				models.APIResponseError409{
					Code:    http.StatusConflict,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.update",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
		c,
		"namespace.link.update",
		l,
	))
}

// deleteLinkFromNamespace
// @Summary Delete a Link from Namespace
// @Description delete a Link from Namespace
// @Tags namespaces
// @Produce json
// @Param name path string true "Namespace"
// @Param linkID path string true "Link ID"
// @Success 200 {object} models.APIResponse{}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/link/{linkID} [delete].
func (h *v1) deleteLinkFromNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.delete.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	linkID := c.Param("linkID")
	if linkID == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.delete.linkID",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "linkID is empty",
			},
		))
	}

	if err := h.DB.DeleteLink(c.Request().Context(), models.LinkID(linkID), name); err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.delete",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.delete",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[struct{}](
		c,
		"namespace.link.delete",
		struct{}{},
	))
}

// deleteLinkByNameFromNamespace
// @Summary Delete a Link by name from Namespace
// @Description delete a Link by name from Namespace
// @Tags namespaces
// @Produce json
// @Param name path string true "Namespace"
// @Param linkName path string true "Link Name"
// @Success 200 {object} models.APIResponse{}
// @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// @Router /namespace/{name}/link/name/{linkName} [delete].
func (h *v1) deleteLinkByNameFromNamespace(c echo.Context) error {
	name := c.Param("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.delete.name",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "namespace is empty",
			},
		))
	}

	linkName := c.Param("linkName")
	if linkName == "" {
		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
			c,
			"namespace.link.delete.linkName",
			models.APIResponseError400{
				Code:    http.StatusBadRequest,
				Message: "linkName is empty",
			},
		))
	}

	// GetLinkByName
	link, err := h.DB.GetLinkByName(c.Request().Context(), linkName, name)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.delete",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		case errors.Is(err, models.ErrInvalid):
			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(
				c,
				"namespace.link.delete",
				models.APIResponseError400{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.delete",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	if err := h.DB.DeleteLink(c.Request().Context(), link.ID, name); err != nil {
		switch {
		case errors.Is(err, models.ErrNotFound):
			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(
				c,
				"namespace.link.delete",
				models.APIResponseError404{
					Code:    http.StatusNotFound,
					Message: err.Error(),
				},
			))
		default:
			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(
				c,
				"namespace.link.delete",
				models.APIResponseError500{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			))
		}
	}

	return c.JSON(http.StatusOK, model.NewAPIResponse[struct{}](
		c,
		"namespace.link.delete",
		struct{}{},
	))
}
