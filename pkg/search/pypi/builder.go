package pypi

import (
	"net/http"
	"net/url"

	"github.com/huangsam/namigo/internal/util"
)

// PypiListingResponse represents the response from the PyPI simple API.
type PypiListingResponse struct {
	Meta struct {
		LastSerial int    `json:"_last-serial"`
		APIVersion string `json:"api-version"`
	} `json:"meta"`
	Projects []struct {
		LastSerial int    `json:"_last-serial"`
		Name       string `json:"name"`
	} `json:"projects"`
}

// PypiDetailResponse represents the response from the PyPI json API.
type PypiDetailResponse struct {
	Info struct {
		Author      string `json:"author"`
		Description string `json:"description"`
		Summary     string `json:"summary"`
		Version     string `json:"version"`
	}
}

// listing builds the request for PyPI list view.
func listing() util.RequestBuilder {
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

// detail builds the request for PyPI detail view.
func detail(pkg string) util.RequestBuilder {
	return func() (*http.Request, error) {
		url := url.URL{
			Scheme: "https",
			Host:   "pypi.org",
			Path:   "pypi/" + pkg + "/json",
		}
		return http.NewRequest("GET", url.String(), nil)
	}
}
