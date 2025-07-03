package pypi

import (
	"net/http"
	"net/url"

	"github.com/huangsam/namigo/internal/core"
)

// APIList builds a request for PyPI list view.
func APIList() core.RequestBuilder {
	return func() (*http.Request, error) {
		url := url.URL{
			Scheme: "https",
			Host:   "pypi.org",
			Path:   "simple/",
		}
		req, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Accept", "application/vnd.pypi.simple.v1+json")
		return req, nil
	}
}

// APIDetail builds a request for PyPI detail view.
func APIDetail(pkg string) core.RequestBuilder {
	return func() (*http.Request, error) {
		url := url.URL{
			Scheme: "https",
			Host:   "pypi.org",
			Path:   "pypi/" + pkg + "/json",
		}
		return http.NewRequest("GET", url.String(), nil)
	}
}
