package golang

import (
	"net/http"
	"net/url"

	"github.com/huangsam/namigo/v2/internal/core"
)

// ScrapeList builds the request for Go package list view.
func ScrapeList(name string) core.RequestBuilder {
	return ScrapeListWithBaseURL(name, "https://pkg.go.dev")
}

// ScrapeListWithBaseURL builds the request with a custom base URL.
func ScrapeListWithBaseURL(name, baseURL string) core.RequestBuilder {
	return func() (*http.Request, error) {
		encodedName := url.PathEscape(name)
		params := url.Values{"q": []string{encodedName}, "m": []string{"package"}}
		fullURL := baseURL + "/search?" + params.Encode()
		return http.NewRequest("GET", fullURL, nil)
	}
}
