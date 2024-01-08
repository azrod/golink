package short

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/azrod/golink/api/model"
	"github.com/azrod/golink/models"
	"github.com/azrod/golink/pkg/sb"
)

type Short model.Handlers

func NewHandlers(db sb.Client, e *echo.Echo) *Short {
	h := &Short{
		DB:         db,
		EchoServer: e,
	}

	h.EchoServer.GET("/:path", h.redirectLink)

	return h
}

func (s *Short) redirectLink(c echo.Context) error {
	path := c.Param("path")
	// path := fmt.Sprintf("/%s", c.Param("path"))

	// Split path by "/"
	paths := strings.Split(path, "/")

	var (
		links []models.Link
		link  models.Link
		ns    models.Namespace
		err   error
	)

	switch {
	case len(paths) == 1:
		link, err = s.DB.GetLinkByPath(c.Request().Context(), "/"+path, "default")
		if err != nil {
			if errors.Is(err, models.ErrNotFound) {
				// Link is not found, find if namespace exists and redirect to namespace page
				ns, err = s.DB.GetNamespace(c.Request().Context(), paths[0])
				if err != nil {
					goto RETURNERROR
				}

				return c.HTML(http.StatusTemporaryRedirect, "/u/"+ns.Name)
			}
			goto RETURNERROR
		}
		goto REDIRECT
	default:
		// find link by path and namespace
		ns, err = s.DB.GetNamespace(c.Request().Context(), paths[0])
		if err != nil {
			goto RETURNERROR
		}

		links, err = s.DB.ListLinks(c.Request().Context(), ns.Name)
		if err != nil {
			goto RETURNERROR
		}

		for _, l := range links {
			if l.SourcePath == fmt.Sprintf("/%s", paths[1]) && l.NameSpace == ns.Name {
				link = l
				goto REDIRECT
			}
		}

		err = models.ErrNotFound
		goto RETURNERROR
	}

RETURNERROR:
	switch {
	// TODO add html PAGE if UserAgent is a web browser and json output for other UserAgents
	case errors.Is(err, models.ErrNotFound):
		return c.HTML(http.StatusNotFound, "Not Found")
	case err != nil:
		return c.JSON(http.StatusInternalServerError, err)
	}

REDIRECT:
	return c.Redirect(http.StatusTemporaryRedirect, link.TargetURL)
}
