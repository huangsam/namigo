package pypi

import (
	"net/http"
	"net/url"

	"github.com/huangsam/namigo/internal/util"
)

// PyPIAPIListingResponse represents listing info from the PyPI API.
type PyPIAPIListingResponse struct {
	Meta struct {
		LastSerial int    `json:"_last-serial"`
		APIVersion string `json:"api-version"`
	} `json:"meta"`
	Projects []struct {
		LastSerial int    `json:"_last-serial"`
		Name       string `json:"name"`
	} `json:"projects"`
}

// PyPIAPIDetailResponse represents detailed info from the PyPI API.
type PyPIAPIDetailResponse struct {
	Info struct {
		Author      string `json:"author"`
		Description string `json:"description"`
		Summary     string `json:"summary"`
		Version     string `json:"version"`
	}
}

// APIListing builds a request for PyPI list view.
func APIListing() util.RequestBuilder {
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

// APIDetail builds a request for PyPI APIDetail view.
func APIDetail(pkg string) util.RequestBuilder {
	return func() (*http.Request, error) {
		url := url.URL{
			Scheme: "https",
			Host:   "pypi.org",
			Path:   "pypi/" + pkg + "/json",
		}
		return http.NewRequest("GET", url.String(), nil)
	}
}
