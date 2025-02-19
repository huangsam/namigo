package golang

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for Go packages by scraping pkg.go.dev.
func SearchByScrape(name string) ([]model.GoPackageResult, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	pipeline := util.NewDocumentPipeline(client, listing(name))
	doc, err := pipeline.Execute()
	if err != nil {
		return []model.GoPackageResult{}, err
	}

	result := []model.GoPackageResult{}

	doc.Find(".SearchSnippet").Each(func(i int, section *goquery.Selection) {
		content := strings.Fields(section.Find("h2").Text())
		pkg, path := content[0], content[1]
		if !strings.Contains(pkg, name) && !strings.Contains(path, name) {
			return
		}

		description := strings.Trim(section.Find("p").Text(), " \n")
		if len(description) == 0 {
			description = model.NoDescription
		}

		result = append(result, model.GoPackageResult{
			Name:        pkg,
			Path:        path,
			Description: description,
		})
	})

	return result, nil
}
