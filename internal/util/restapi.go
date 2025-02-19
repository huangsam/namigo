package util

import (
	"errors"
	"io"
	"net/http"
)

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
	if res.StatusCode >= 300 {
		return nil, errors.New("Bad status: " + res.Status)
	}
	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
