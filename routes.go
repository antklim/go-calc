package calc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	errReadBody         = errors.New("body read failed")
	errParseBody        = errors.New("parsing JSON failed")
	serverErrorResponse = []byte(`{"message": "internal server error"}`)
	notFoundResponse    = []byte(`{"message": "not found"}`)
)

// Routes available in calc service.
func Routes(client HTTPClient) map[string]http.HandlerFunc {
	return map[string]http.HandlerFunc{
		"/do":     doHandler(),
		"/remote": remoteHandler(client),
		"/":       notFoundHandler,
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeRawBody(w, r, notFoundResponse, http.StatusNotFound)
}

func doHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := unmarshalBody(r, &req); err != nil {
			log.Println(err)
			response := NewErrorResponse(err)
			writeBody(w, r, response, http.StatusBadRequest)
			return
		}

		if err := req.Validate(); err != nil {
			log.Println(err)
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

func remoteHandler(client HTTPClient) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			url = "http://example.com/"
		}

		req, err := http.NewRequest(http.MethodPost, url, r.Body)
		if err != nil {
			log.Println(err)
			response := NewErrorResponse(err)
			writeBody(w, r, response, http.StatusInternalServerError)
			return
		}

		res, err := client.Do(req)
		if err != nil {
			log.Println(err)
			response := NewErrorResponse(err)
			writeBody(w, r, response, http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			response := NewErrorResponse(err)
			writeBody(w, r, response, http.StatusInternalServerError)
			return
		}

		if res.StatusCode != http.StatusOK {
			err := fmt.Errorf("remote call failed, returned status code %d", res.StatusCode)
			log.Println(err)
			response := NewErrorResponse(err)
			writeBody(w, r, response, http.StatusInternalServerError)
			return
		}

		writeRawBody(w, req, body, http.StatusOK)
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
		body = serverErrorResponse
		status = http.StatusInternalServerError
	}

	writeRawBody(w, req, body, status)
}

func writeRawBody(w http.ResponseWriter, req *http.Request, body []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(append(body, byte('\n')))
}
