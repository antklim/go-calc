package calc

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/antklim/go-calc/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestNotFoundHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/not-found", nil)
	if err != nil {
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(notFoundHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.JSONEq(t, `{"message": "not found"}`, rr.Body.String())
}

func TestDoHandler(t *testing.T) {
	reqBody := `{"operation": "add", "arguments": [1,2]}`
	req, err := http.NewRequest("POST", "/do", strings.NewReader(reqBody))
	if err != nil {
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(doHandler())
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	resBody := `{"operation": "add", "arguments": [1,2], "result": 3}`
	assert.JSONEq(t, resBody, rr.Body.String())
}

func TestRemoteHandlerWithClientMock(t *testing.T) {
	res := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(`{"foo": "bar"}`)),
	}
	clientMock := mocks.HTTPClient{}
	clientMock.On("Do", mock.Anything, mock.Anything).Return(res, nil)

	req, err := http.NewRequest("GET", "/remote", nil)
	if err != nil {
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(remoteHandler(&clientMock))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	resBody := `{"foo": "bar"}`
	assert.JSONEq(t, resBody, rr.Body.String())
}

func TestRemoteHandlerWithStubServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))

	defer ts.Close()

	url := fmt.Sprintf("/remote?url=%s", ts.URL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		require.NoError(t, err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(remoteHandler(&http.Client{}))
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.JSONEq(t, `{"message": "remote call failed, returned status code 503"}`, rr.Body.String())
}
