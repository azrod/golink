package apiv1

// // listLinks
// // @Summary List all links
// // @Description get all links
// // @Tags links
// // @Produce  json
// // @Success 200 {array} models.APIResponse{data=[]models.Link}
// // @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// // @Router /links [get].
// func (h *v1) listLinks(c echo.Context) error {
// 	links, err := h.DB.ListLinks(c.Request().Context())
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "links.list", models.APIResponseError500{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		}))
// 	}

// 	return c.JSON(http.StatusOK, model.NewAPIResponse[[]models.Link](
// 		c,
// 		"links.list",
// 		links,
// 	))
// }

// // getLinkByID
// // @Summary Get link by ID
// // @Description get link by ID
// // @Tags links
// // @Produce  json
// // @Param id path string true "Link ID"
// // @Success 200 {object} models.APIResponse{data=models.Link}
// // @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// // @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// // @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// // @Router /link/{id} [get].
// func (h *v1) getLinkByID(c echo.Context) error {
// 	linkID := c.Param("id")
// 	if linkID == "" {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.get.id", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: "linkID is empty",
// 		}))
// 	}

// 	link, err := h.DB.GetLinkByID(c.Request().Context(), models.LinkID(linkID))
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, models.ErrNotFound):
// 			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "link.get.id", models.APIResponseError404{
// 				Code:    http.StatusNotFound,
// 				Message: err.Error(),
// 			}))
// 		case errors.Is(err, models.ErrInvalid):
// 			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.get.id", models.APIResponseError400{
// 				Code:    http.StatusBadRequest,
// 				Message: err.Error(),
// 			}))
// 		default:
// 			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "link.get.id", models.APIResponseError500{
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			}))
// 		}
// 	}

// 	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
// 		c,
// 		"label.get",
// 		link,
// 	))
// }

// // getLinkByName
// // @Summary Get link by name
// // @Description get link by name
// // @Tags links
// // @Produce  json
// // @Param name path string true "Link name"
// // @Success 200 {object} models.APIResponse{data=models.Link}
// // @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// // @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// // @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// // @Router /link/name/{name} [get].
// func (h *v1) getLinkByName(c echo.Context) error {
// 	linkName := c.Param("name")
// 	if linkName == "" {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.get.name", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: "linkName is empty",
// 		}))
// 	}

// 	link, err := h.DB.GetLinkByName(c.Request().Context(), linkName)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, models.ErrNotFound):
// 			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "link.get.name", models.APIResponseError404{
// 				Code:    http.StatusNotFound,
// 				Message: err.Error(),
// 			}))
// 		case errors.Is(err, models.ErrInvalid):
// 			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.get.name", models.APIResponseError400{
// 				Code:    http.StatusBadRequest,
// 				Message: err.Error(),
// 			}))
// 		default:
// 			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "link.get.name", models.APIResponseError500{
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			}))
// 		}
// 	}

// 	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
// 		c,
// 		"label.get",
// 		link,
// 	))
// }

// // createLink
// // @Summary Create link
// // @Description create link
// // @Tags links
// // @Accept  json
// // @Produce  json
// // @Param link body models.LinkRequest true "Link"
// // @Success 201 {object} models.APIResponse{data=models.Link}
// // @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// // @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// // @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// // @failure 409 {object} models.APIResponseError{error=models.APIResponseError409}
// // @Router /link [post].
// func (h *v1) createLink(c echo.Context) error {
// 	var link models.Link
// 	if err := c.Bind(&link); err != nil {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.create", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		}))
// 	}

// 	if err := link.Validate(); err != nil {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.create", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		}))
// 	}

// 	// Check if link path or name already exists in namespace
// 	links, err := h.DB.ListLinks(c.Request().Context())
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "link.create", models.APIResponseError500{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error(),
// 		}))
// 	}

// 	for _, l := range links {
// 		if l.NameSpace == link.NameSpace {
// 			switch {
// 			case l.Name == link.Name:
// 				return c.JSON(http.StatusConflict, model.NewAPIResponseError(c, "link.create", models.APIResponseError409{
// 					Code:    http.StatusConflict,
// 					Message: "link name already exists in namespace",
// 				}))
// 			case l.SourcePath == link.SourcePath:
// 				return c.JSON(http.StatusConflict, model.NewAPIResponseError(c, "link.create", models.APIResponseError409{
// 					Code:    http.StatusConflict,
// 					Message: "link path already exists in namespace",
// 				}))
// 			}
// 		}
// 	}

// 	if err := h.DB.CreateLink(c.Request().Context(), link); err != nil {
// 		switch {
// 		case errors.Is(err, models.ErrAlreadyExists):
// 			return c.JSON(http.StatusConflict, model.NewAPIResponseError(c, "link.create", models.APIResponseError409{
// 				Code:    http.StatusConflict,
// 				Message: err.Error(),
// 			}))
// 		case errors.Is(err, models.ErrInvalid):
// 			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.create", models.APIResponseError400{
// 				Code:    http.StatusBadRequest,
// 				Message: err.Error(),
// 			}))
// 		default:
// 			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "link.create", models.APIResponseError500{
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			}))
// 		}
// 	}

// 	return c.JSON(http.StatusCreated, model.NewAPIResponse[models.Link](
// 		c,
// 		"label.get",
// 		link,
// 	))
// }

// // updateLink
// // @Summary Update link
// // @Description update link
// // @Tags links
// // @Accept  json
// // @Produce  json
// // @Param link body models.Link true "Link"
// // @Success 200 {object} models.APIResponse{data=models.Link}
// // @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// // @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// // @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// // @Router /link [put].
// func (h *v1) updateLink(c echo.Context) error {
// 	var link models.Link
// 	if err := c.Bind(&link); err != nil {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.update", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		}))
// 	}

// 	if err := link.Validate(); err != nil {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.update", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error(),
// 		}))
// 	}

// 	if err := h.DB.UpdateLink(c.Request().Context(), link); err != nil {
// 		switch {
// 		case errors.Is(err, models.ErrNotFound):
// 			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "link.update", models.APIResponseError404{
// 				Code:    http.StatusNotFound,
// 				Message: err.Error(),
// 			}))
// 		case errors.Is(err, models.ErrInvalid):
// 			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.update", models.APIResponseError400{
// 				Code:    http.StatusBadRequest,
// 				Message: err.Error(),
// 			}))
// 		default:
// 			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "link.update", models.APIResponseError500{
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			}))
// 		}
// 	}

// 	return c.JSON(http.StatusOK, model.NewAPIResponse[models.Link](
// 		c,
// 		"label.get",
// 		link,
// 	))
// }

// // deleteLink
// // @Summary Delete link
// // @Description delete link
// // @Tags links
// // @Produce  json
// // @Param id path string true "Link ID"
// // @Success 200 {object} models.APIResponse{}
// // @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// // @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// // @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// // @Router /link/{id} [delete].
// func (h *v1) deleteLink(c echo.Context) error {
// 	linkID := c.Param("id")
// 	if linkID == "" {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "link.delete", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: "linkID is empty",
// 		}))
// 	}

// 	if err := h.DB.DeleteLink(c.Request().Context(), models.LinkID(linkID)); err != nil {
// 		switch {
// 		case errors.Is(err, models.ErrNotFound):
// 			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "link.update", models.APIResponseError404{
// 				Code:    http.StatusNotFound,
// 				Message: err.Error(),
// 			}))
// 		default:
// 			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "link.update", models.APIResponseError500{
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			}))
// 		}
// 	}

// 	return c.JSON(http.StatusOK, model.NewAPIResponse[struct{}](
// 		c,
// 		"label.delete",
// 		struct{}{},
// 	))
// }

// // listLinksByLabel
// // @Summary List all links by label
// // @Description get all links by label
// // @Tags links
// // @Produce  json
// // @Param id path string true "Label ID"
// // @Success 200 {object} models.APIResponse{data=[]models.Link}
// // @failure 500 {object} models.APIResponseError{error=models.APIResponseError500}
// // @failure 404 {object} models.APIResponseError{error=models.APIResponseError404}
// // @failure 400 {object} models.APIResponseError{error=models.APIResponseError400}
// // @Router /label/{id}/links [get].
// func (h *v1) listLinksByLabel(c echo.Context) error {
// 	labelID := c.Param("id")
// 	if labelID == "" {
// 		return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "links.list.label", models.APIResponseError400{
// 			Code:    http.StatusBadRequest,
// 			Message: "labelID is empty",
// 		}))
// 	}

// 	links, err := h.DB.ListLinksByLabel(c.Request().Context(), models.LabelID(labelID))
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, models.ErrNotFound):
// 			return c.JSON(http.StatusNotFound, model.NewAPIResponseError(c, "links.list.label", models.APIResponseError404{
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			}))
// 		case errors.Is(err, models.ErrInvalid):
// 			return c.JSON(http.StatusBadRequest, model.NewAPIResponseError(c, "links.list.label", models.APIResponseError400{
// 				Code:    http.StatusBadRequest,
// 				Message: err.Error(),
// 			}))
// 		default:
// 			return c.JSON(http.StatusInternalServerError, model.NewAPIResponseError(c, "links.list.label", models.APIResponseError500{
// 				Code:    http.StatusInternalServerError,
// 				Message: err.Error(),
// 			}))
// 		}
// 	}

// 	return c.JSON(http.StatusOK, model.NewAPIResponse[[]models.Link](
// 		c,
// 		"links.list.label",
// 		links,
// 	))
// }
