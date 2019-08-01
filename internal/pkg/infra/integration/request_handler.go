package integration

import "github.com/mmpereira2github/financial-go/internal/pkg/services"

// RequestHandler is a kind of Interface Adapter to provide input to service execution and return output in the approprieted format
type RequestHandler interface {
	GetServiceID() string
	IsDebugRequiredByRequest() bool
	UnmarshalServiceInput(serviceContext *services.ServiceContext, serviceInput interface{}) (interface{}, *services.Status)
}

// ErrorDetail provides the status related to error and the service execution context
type ErrorDetail struct {
	Status         *services.Status
	ServiceContext *services.ServiceContext
}

// InvokeService invokes the service associated with the given handler
func InvokeService(requestHandler RequestHandler) (serviceOutput interface{}, errorDetail *ErrorDetail) {
	servicename := requestHandler.GetServiceID()
	serviceContext := services.NewServiceContext()
	serviceContext.SetServiceName(servicename)
	serviceContext.SetDebugEnabled(requestHandler.IsDebugRequiredByRequest())

	var status *services.Status
	if entry, status := services.Manager.GetServiceEntryByID(servicename); status == nil {
		var input = entry.InputFactory()
		var output interface{}
		input, status = requestHandler.UnmarshalServiceInput(serviceContext, input)
		if status == nil {
			output, status = entry.Invoke(input)
			if output != nil {
				return output, nil
			}
		}
	}

	return nil, &ErrorDetail{status, serviceContext}
}
