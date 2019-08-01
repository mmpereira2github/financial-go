package httpgin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mmpereira2github/financial-go/internal/pkg/infra/integration"
	"github.com/mmpereira2github/financial-go/internal/pkg/services"
	"github.com/pkg/errors"
)

type jsonRequestHandler struct {
	c *gin.Context
}

func (h *jsonRequestHandler) GetServiceID() string {
	return h.c.Param("servicename")
}

func (h *jsonRequestHandler) IsDebugRequiredByRequest() bool {
	return h.c.GetHeader("debug") != ""
}

func (h *jsonRequestHandler) UnmarshalServiceInput(serviceContext *services.ServiceContext, serviceInput interface{}) (interface{}, *services.Status) {
	var err error
	jsonAsString := "not-available"
	if serviceContext.IsDebugEnabled() {
		err = h.c.ShouldBindBodyWith(serviceInput, binding.JSON)

		if buf, ok := h.c.Get(gin.BodyBytesKey); ok {
			jsonAsString = string(buf.([]byte)[0:])
			serviceContext.SetRequestAsString(jsonAsString)
		}
	} else {
		err = h.c.ShouldBindJSON(serviceInput)
	}
	if err == nil {
		return serviceInput, nil
	}
	return serviceInput, &services.Status{
		Code:  services.InvalidInput,
		Error: errors.Wrapf(err, "request=[%s]", jsonAsString),
	}
}

// NewRequestHandler creates a RequestHandler to handle JSON requests
func NewRequestHandler(c *gin.Context) integration.RequestHandler { return &jsonRequestHandler{c} }
