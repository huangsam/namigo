package npm

import (
	"net/http"
	"net/url"

	"github.com/huangsam/namigo/internal/util"
)

// listing builds the request for NPM list view.
func listing(name string) util.RequestBuilder {
	return func() (*http.Request, error) {
		encodedName := url.PathEscape(name)
		params := url.Values{"q": []string{encodedName}}
		url := url.URL{
			Scheme:   "https",
			Host:     "www.npmjs.com",
			Path:     "search",
			RawQuery: params.Encode(),
		}
		return http.NewRequest("GET", url.String(), nil)
	}
}
