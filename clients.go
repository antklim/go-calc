package calc

import "net/http"

// HTTPClient interface.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHTTPClient ...
func NewHTTPClient() HTTPClient {
	return &http.Client{}
}
