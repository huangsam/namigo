package golang

import (
	"net/http"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for Go packages by scraping pkg.go.dev.
func SearchByScrape(name string, size int) ([]model.GoPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	pipeline := util.NewDocumentPipeline(client, ScrapeList(name))
	doc, err := pipeline.Execute()
	if err != nil {
		return nil, err
	}
	return docWorker(doc, size, name), nil
}
