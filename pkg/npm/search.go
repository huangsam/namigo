package npm

import (
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for NPM packages by scraping www.npmjs.com.
func SearchByScrape(name string) ([]model.NPMPackageResult, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	pipeline := util.NewDocumentPipeline(client, listing(name))
	doc, err := pipeline.Execute()
	if err != nil {
		return []model.NPMPackageResult{}, err
	}

	result := []model.NPMPackageResult{}

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

	return result, nil
}
