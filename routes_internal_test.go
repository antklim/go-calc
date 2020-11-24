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
)

func TestNotFoundHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/not-found", nil)
	rr := httptest.NewRecorder()

	notFoundHandler(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusNotFound, res.StatusCode)
	assert.JSONEq(t, `{"message": "not found"}`, string(resBody))
}

func TestDoHandler(t *testing.T) {
	req := httptest.NewRequest("POST", "/do", strings.NewReader(`{"operation": "add", "arguments": [1,2]}`))
	rr := httptest.NewRecorder()

	doHandler()(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.JSONEq(t, `{"operation": "add", "arguments": [1,2], "result": 3}`, string(resBody))
}

func TestRemoteHandlerWithClientMock(t *testing.T) {
	mockRes := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(`{"foo": "bar"}`)),
	}
	clientMock := mocks.HTTPClient{}
	clientMock.On("Do", mock.AnythingOfType("*http.Request")).Return(mockRes, nil)

	req := httptest.NewRequest("GET", "/remote", nil)
	rr := httptest.NewRecorder()

	remoteHandler(&clientMock)(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.JSONEq(t, `{"foo": "bar"}`, string(resBody))
}

func TestRemoteHandlerWithStubServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer ts.Close()

	url := fmt.Sprintf("/remote?url=%s", ts.URL)
	req := httptest.NewRequest("GET", url, nil)
	rr := httptest.NewRecorder()

	remoteHandler(&http.Client{})(rr, req)
	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.JSONEq(t, `{"message": "remote call failed, returned status code 503"}`, string(resBody))
}
