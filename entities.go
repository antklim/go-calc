package calc

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Request describes service request data structure.
type Request struct {
	Operation string    `json:"operation"`
	Arguments []float64 `json:"arguments"`
}

func UnmarshalRequest(r io.Reader) (*Request, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		return nil, err
	}

	return &req, nil
}

// Validate ...
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
	Message string         `json:"message"`
	Errors  []ServiceError `json:"errors,omitempy"`
}

// ServiceError describes service error data structure (validation, operation errors, etc).
type ServiceError struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}
