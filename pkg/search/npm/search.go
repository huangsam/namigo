package npm

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByScrape searches for NPM packages by scraping www.npmjs.com.
func SearchByScrape(name string, max int) ([]model.NPMPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	pipeline := util.NewDocumentPipeline(client, ScrapeListing(name))
	doc, err := pipeline.Execute()
	if err != nil {
		return []model.NPMPackage{}, err
	}

	result := []model.NPMPackage{}

	docWorker(doc, &result, max)

	return result, nil
}

// SearchByAPI searches for NPM packages by querying registry.npmjs.com.
func SearchByAPI(name string, max int) ([]model.NPMPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	bl, err := util.RESTAPIQuery(client, APIListing(name, max))
	if err != nil {
		return []model.NPMPackage{}, err
	}

	var listingRes NPMAPIListingResponse
	if err := json.Unmarshal(bl, &listingRes); err != nil {
		return []model.NPMPackage{}, err
	}

	return apiWorker(&listingRes, name)
}
