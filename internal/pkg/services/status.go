package services

import "encoding/json"

// StatusCode it is used to identify the error type
type StatusCode int

const (
	// Success means success
	Success StatusCode = 0
	// ServiceNotFound means the service asked to run does not exist
	ServiceNotFound StatusCode = -1
	// InvalidInput means it is invalid for service execution
	InvalidInput StatusCode = -2
	// DateNotFound means that no index value was found for given date
	DateNotFound StatusCode = -3
	// IndexIDNotFound means the given index to update a value was not found in the application
	IndexIDNotFound StatusCode = -4
)

// Status provides a status about the service executed
type Status struct {
	Code  StatusCode
	Error error
}

// MarshalJSON converts to JSON
func (s *Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code    StatusCode `json:"code"`
		Message string     `json:"message"`
	}{
		Code:    s.Code,
		Message: s.Error.Error(),
	})
}
