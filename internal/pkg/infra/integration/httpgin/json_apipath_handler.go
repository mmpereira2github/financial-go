package httpgin

import (
	"net/http"

	"github.com/mmpereira2github/financial-go/internal/pkg/infra/integration"
	"github.com/mmpereira2github/financial-go/internal/pkg/services"

	"github.com/gin-gonic/gin"
)

// handleServiceInvocation identifies and invokes the service specified in URL
func jsonAPIPathHandler(c *gin.Context) {
	handler := NewRequestHandler(c)
	output, errorDetail := integration.InvokeService(handler)
	if output != nil {
		c.JSON(http.StatusOK, output)
	} else {
		switch errorDetail.Status.Code {
		case services.InvalidInput:
			c.JSON(http.StatusBadRequest, errorDetail)
		case services.ServiceNotFound:
			c.JSON(http.StatusNotFound, errorDetail)
		default:
			c.JSON(http.StatusInternalServerError, errorDetail)
		}
	}
}
