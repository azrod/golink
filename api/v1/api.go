package apiv1

import (
	"github.com/labstack/echo/v4"

	"github.com/azrod/golink/api/model"
	"github.com/azrod/golink/pkg/sb"
)

type v1 model.HandlerAPIVersion

func New(db sb.Client, e *echo.Group) {
	h := &v1{
		DB:        db,
		EchoGroup: e,
	}

	v1 := h.EchoGroup.Group("/v1")

	// * Link

	// * Label
	v1.GET("/labels", h.listLabels)
	// v1.GET("/label/:id/links", h.listLinksByLabel)
	v1.GET("/label/:id", h.getLabelByID)
	v1.GET("/label/name/:name", h.getLabelByName)
	v1.POST("/label", h.createLabel)
	v1.PUT("/label", h.updateLabel)
	v1.DELETE("/label/:id", h.deleteLabel)
	v1.DELETE("/label/name/:name", h.deleteLabelByName)

	// * Namespace
	v1.GET("/namespaces", h.listNamespaces)
	v1.GET("/namespace/:name", h.getNamespace)
	v1.GET("/namespace/:name/links", h.listLinksByNamespace)
	v1.GET("/namespace/:name/link/:linkid", h.getLinkFromNamespace)
	v1.GET("/namespace/:name/link/name/:linkname", h.getLinkByNameFromNamespace)
	v1.GET("/namespace/:name/link/path/:linkname", h.getLinkByPathFromNamespace)

	v1.POST("/namespace", h.createNamespace)
	v1.POST("/namespace/:name/link", h.addLinkToNamespace)

	v1.PUT("/namespace/:name/link/:linkid", h.updateLinkInNamespace)

	v1.DELETE("/namespace/:name", h.deleteNamespace)
	v1.DELETE("/namespace/:name/link/:linkid", h.deleteLinkFromNamespace)
	v1.DELETE("/namespace/:name/link/name/:linkname", h.deleteLinkByNameFromNamespace)
}
