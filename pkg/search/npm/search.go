package npm

import (
	"net/http"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for NPM packages by scraping www.npmjs.com.
func SearchByScrape(name string, max int) ([]model.NPMPackageResult, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	pipeline := util.NewDocumentPipeline(client, Listing(name))
	doc, err := pipeline.Execute()
	if err != nil {
		return []model.NPMPackageResult{}, err
	}

	result := []model.NPMPackageResult{}

	worker(doc, &result, max)

	return result, nil
}
