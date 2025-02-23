package golang

import (
	"net/http"
	"net/url"

	"github.com/huangsam/namigo/internal/util"
)

// Listing builds the request for Go package list view.
func Listing(name string) util.RequestBuilder {
	return func() (*http.Request, error) {
		encodedName := url.PathEscape(name)
		params := url.Values{"q": []string{encodedName}, "m": []string{"package"}}
		url := url.URL{
			Scheme:   "https",
			Host:     "pkg.go.dev",
			Path:     "search",
			RawQuery: params.Encode(),
		}
		return http.NewRequest("GET", url.String(), nil)
	}
}
