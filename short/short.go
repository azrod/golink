package short

import (
	"errors"
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

	// Split path by "/"
	paths := strings.Split(path, "/")

	const (
		uaTypeBrowser = iota
		uaTypeCLI
	)

	var (
		ns     = "default"
		target = paths[0]

		uaType = func() int {
			ua := c.Request().UserAgent()

			// Check if UserAgent is a web browser
			switch {
			case strings.Contains(ua, "curl") || strings.Contains(ua, "wget"):
				return uaTypeCLI
			default:
				return uaTypeBrowser
			}
		}()
	)

	if len(paths) != 1 {
		ns = paths[0]
		target = paths[1]
	}

	link, err := s.DB.GetLinkByPath(c.Request().Context(), "/"+target, ns)
	if err != nil {
		goto RETURNERROR
	}

	goto REDIRECT

RETURNERROR:

	switch {
	case errors.Is(err, models.ErrNotFound):
		switch uaType {
		case uaTypeBrowser:
			return s.handleHTML404(c)
		case uaTypeCLI:
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "not found",
			})
		}
	case err != nil:
		return c.JSON(http.StatusInternalServerError, err)
	}

REDIRECT:
	switch uaType {
	case uaTypeCLI:
		return c.String(http.StatusTemporaryRedirect, link.TargetURL)
	default:
		return c.Redirect(http.StatusTemporaryRedirect, link.TargetURL)
	}
}
