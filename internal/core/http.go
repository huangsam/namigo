package core

import (
	"errors"
	"io"
	"net/http"
)

// RequestBuilder builds a HTTP request for client.
type RequestBuilder func() (*http.Request, error)

// ResponseHandler handles side effects for a response.
type ResponseHandler func(*http.Response) error

// RESTAPIQuery accesses an API endpoint and returns bytes.
func RESTAPIQuery(c *http.Client, builder RequestBuilder) ([]byte, error) {
	req, err := builder()
	if err != nil {
		return nil, err
	}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("Bad status: " + res.Status)
	}
	defer dismiss(res.Body.Close)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
