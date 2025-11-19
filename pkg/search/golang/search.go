package golang

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/v2/internal/core"
	"github.com/huangsam/namigo/v2/internal/model"
)

// SearchByScrape searches for Go packages by scraping pkg.go.dev.
func SearchByScrape(client *http.Client, name string, size int) ([]model.GoPackage, error) {
	return SearchByScrapeWithBuilder(client, name, size, ScrapeList(name))
}

// SearchByScrapeWithBuilder searches using a custom request builder.
func SearchByScrapeWithBuilder(client *http.Client, name string, size int, builder core.RequestBuilder) ([]model.GoPackage, error) {
	pipeline := core.NewDocumentPipeline(client, builder)
	doc, err := pipeline.Execute()
	if err != nil {
		return nil, err
	}

	result := make([]model.GoPackage, 0, size) // Pre-allocate with capacity
	doc.Find(".SearchSnippet").Each(func(_ int, section *goquery.Selection) {
		if len(result) >= size {
			return
		}

		content := strings.Fields(section.Find("h2").Text())
		pkg, path := content[0], strings.Trim(content[1], "()")
		if !strings.Contains(pkg, name) && !strings.Contains(path, name) {
			return
		}

		description := strings.TrimSpace(section.Find("p").Text())
		if len(description) == 0 {
			description = model.NoDescription
		}

		result = append(result, model.GoPackage{
			Name:        pkg,
			Path:        path,
			Description: description,
		})
	})
	return result, nil
}
