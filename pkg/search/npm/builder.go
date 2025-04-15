package npm

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/huangsam/namigo/internal/util"
)

// APIList builds the request for NPM list view.
func APIList(name string, size int) util.RequestBuilder {
	return func() (*http.Request, error) {
		params := url.Values{}
		params.Add("text", name)
		params.Add("size", strconv.Itoa(size))
		url := url.URL{
			Scheme:   "https",
			Host:     "registry.npmjs.com",
			Path:     "-/v1/search",
			RawQuery: params.Encode(),
		}
		return http.NewRequest("GET", url.String(), nil)
	}
}
