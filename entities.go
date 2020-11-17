package calc

import (
	"fmt"
)

// Request describes service request data structure.
type Request struct {
	Operation string    `json:"operation"`
	Arguments []float64 `json:"arguments"`
}

// Validate implements request validation.
func (r *Request) Validate() error {
	operation, ok := Operations[r.Operation]
	if !ok {
		return fmt.Errorf("unsupported operation %s", r.Operation)
	}

	return operation.validation(r.Arguments)
}

// Response describes service response data structure.
type Response struct {
	Operation string    `json:"operation"`
	Arguments []float64 `json:"arguments"`
	Result    float64   `json:"result"`
}

// ErrorResponse describes service error response data structure.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse ...
func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Message: err.Error()}
}
