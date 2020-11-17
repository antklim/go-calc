package calc

import (
	"fmt"
	"net/http"
)

func DoHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := UnmarshalRequest(r.Body)
		if err != nil {
			fmt.Fprintf(w, "error %v\n", err)
			return
		}

		if err := req.Validate(); err != nil {
			fmt.Fprintf(w, "error %v\n", err)
			return
		}

		operation := Operations[req.Operation]
		result := Do(operation.lambda, req.Arguments)

		fmt.Fprintf(w, "This is result of Do handler %.5f\n", result)
	})
}
