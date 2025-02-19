package npm

import (
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for NPM packages by scraping www.npmjs.com.
func SearchByScrape(name string) []model.NPMPackageResult {
	result := []model.NPMPackageResult{}

	// Setup document query
	client := &http.Client{Timeout: 5 * time.Second}
	builder := func() (*http.Request, error) {
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
	pipeline := util.NewDocumentPipeline(client, builder)
	doc, err := pipeline.Execute()
	if err != nil {
		return result
	}

	// Run document query
	doc.Find("main section").Each(func(i int, section *goquery.Selection) {
		pkg := section.Find("h3").Text()
		match := section.Find("span#pkg-list-exact-match").Text()
		description := section.Find("p").Text()

		result = append(result, model.NPMPackageResult{
			Name:         pkg,
			Description:  description,
			IsExactMatch: len(match) > 0,
		})
	})

	return result
}
