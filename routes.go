package calc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

// Routes available in calc service.
var Routes = map[string]http.HandlerFunc{
	"/do": doHandler(),
	"/":   http.NotFound,
}

var (
	errReadBody  = errors.New("body read failed")
	errParseBody = errors.New("parsing JSON failed")
	serverError  = []byte(`{"message": "internal server error"}`)
)

func doHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := unmarshalBody(r, &req); err != nil {
			response := NewErrorResponse(err)
			writeBody(w, r, response, http.StatusBadRequest)
			return
		}

		if err := req.Validate(); err != nil {
			response := NewErrorResponse(err)
			writeBody(w, r, response, http.StatusUnprocessableEntity)
			return
		}

		operation := Operations[req.Operation]
		result := Do(operation.lambda, req.Arguments)

		response := Response{
			Operation: req.Operation,
			Arguments: req.Arguments,
			Result:    result,
		}

		writeBody(w, r, response, http.StatusOK)
	})
}

func unmarshalBody(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return errReadBody
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, v); err != nil {
		log.Println(err)
		return errParseBody
	}

	return nil
}

func writeBody(w http.ResponseWriter, req *http.Request, response interface{}, status int) {
	body, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		body = serverError
		status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(append(body, byte('\n')))
}
