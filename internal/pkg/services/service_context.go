package services

import "encoding/json"

// ServiceContext stores data related to the service execution
type ServiceContext struct {
	properties map[string]interface{}
}

// NewServiceContext creates a new service context
func NewServiceContext() *ServiceContext {
	return &ServiceContext{properties: make(map[string]interface{})}
}

func (c *ServiceContext) String() string {
	b, err := c.MarshalJSON()
	if err != nil {
		return string(b)
	}
	return err.Error()
}

// MarshalJSON marshals to JSON
func (c *ServiceContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ServiceContext map[string]interface{} `json:"serviceContext"`
	}{
		ServiceContext: c.properties,
	})
}

// SetServiceName sets service name to service context
func (c *ServiceContext) SetServiceName(s string) {
	c.properties["servicename"] = s
}

// GetServiceName returns service name to service context
func (c *ServiceContext) GetServiceName() string {
	value := c.properties["servicename"]
	if s, ok := value.(string); ok {
		return s
	}
	return ""
}

// SetDebugEnabled enables or disables debug
func (c *ServiceContext) SetDebugEnabled(b bool) {
	c.properties["debug.enabled"] = b
}

// IsDebugEnabled returns true if debug enabled
func (c *ServiceContext) IsDebugEnabled() bool {
	value := c.properties["debug.enabled"]
	b := false
	b, _ = value.(bool)
	return b
}

// SetRequestAsString store the request as string in the service context
func (c *ServiceContext) SetRequestAsString(s string) {
	c.properties["debug.requestAsString"] = s
}
