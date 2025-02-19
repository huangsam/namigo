package golang

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for Go packages by scraping pkg.go.dev.
func SearchByScrape(name string) []model.GoPackageResult {
	result := []model.GoPackageResult{}

	// Setup document query
	client := &http.Client{Timeout: 5 * time.Second}
	builder := func() (*http.Request, error) {
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
	pipeline := util.NewDocumentPipeline(client, builder)
	doc, err := pipeline.Execute()
	if err != nil {
		return result
	}

	// Run document query
	doc.Find(".SearchSnippet").Each(func(i int, section *goquery.Selection) {
		content := strings.Fields(section.Find("h2").Text())
		pkg, path := content[0], content[1]
		if !strings.Contains(pkg, name) && !strings.Contains(path, name) {
			return
		}

		description := strings.Trim(section.Find("p").Text(), " \n")
		if len(description) == 0 {
			description = "No description"
		}

		result = append(result, model.GoPackageResult{
			Name:        pkg,
			Path:        path,
			Description: description,
		})
	})

	return result
}
