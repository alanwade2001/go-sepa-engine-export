package handler

import (
	"log/slog"
	"net/http"

	"github.com/alanwade2001/go-sepa-engine-export/internal/service"
	"github.com/alanwade2001/go-sepa-infra/routing"

	"github.com/gin-gonic/gin"
)

type Export struct {
	service *service.Export
}

func NewExport(service *service.Export, r *routing.Router) *Export {
	export := &Export{
		service: service,
	}

	r.Router.GET("/exports/:id", export.PostExport)

	return export
}

// postInitiation adds an initiations from JSON received in the request body.
func (d *Export) PostExport(c *gin.Context) {

	id := c.Param("id")

	if newExport, err := d.service.Export(id); err != nil {
		slog.Error("failed to post execution", "Error", err)
		c.IndentedJSON(http.StatusInternalServerError, newExport)
	} else {
		c.XML(http.StatusOK, newExport)
	}

}
