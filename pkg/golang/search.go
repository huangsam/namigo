package golang

import (
	"net/http"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for Go packages by scraping pkg.go.dev.
func SearchByScrape(name string, max int) ([]model.GoPackageResult, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	pipeline := util.NewDocumentPipeline(client, listing(name))
	doc, err := pipeline.Execute()
	if err != nil {
		return []model.GoPackageResult{}, err
	}

	result := []model.GoPackageResult{}

	worker(doc, &result, max, name)

	return result, nil
}
