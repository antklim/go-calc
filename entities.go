package calc

// Request describes service request data structure.
type Request struct {
	Operation string    `json:"operation"`
	Operands  []float64 `json:"operands"`
}

func (r *Request) Validate() error {
	// if r.Operation == Unknown {
	// 	return errors.New("unknown operation")
	// }

	// if r.Operands == nil {
	// 	return errors.New("operands list cannot be empty")
	// }

	// if len(r.Operands) < 1 {

	// }

	return nil
}

// Response describes service response data structure.
type Response struct {
	Operation string    `json:"operation"`
	Operands  []float64 `json:"operands"`
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
