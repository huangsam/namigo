package golang

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/core"
	"github.com/huangsam/namigo/internal/model"
)

// SearchByScrape searches for Go packages by scraping pkg.go.dev.
func SearchByScrape(name string, size int) ([]model.GoPackage, error) {
	return SearchByScrapeWithBuilder(name, size, ScrapeList(name))
}

// SearchByScrapeWithBuilder searches using a custom request builder.
func SearchByScrapeWithBuilder(name string, size int, builder core.RequestBuilder) ([]model.GoPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	pipeline := core.NewDocumentPipeline(client, builder)
	doc, err := pipeline.Execute()
	if err != nil {
		return nil, err
	}

	result := []model.GoPackage{}
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
