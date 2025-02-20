package pypi

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// workerCount is the number of workers used to grab package details.
const workerCount = 5

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(name string, max int) ([]model.PyPIPackageResult, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	b, err := util.RESTAPIQuery(client, listing())
	if err != nil {
		return []model.PyPIPackageResult{}, err
	}

	var listingRes PypiListingResponse
	if err := json.Unmarshal(b, &listingRes); err != nil {
		return []model.PyPIPackageResult{}, err
	}

	result := []model.PyPIPackageResult{}

	count := 0
	for _, project := range listingRes.Projects {
		if !strings.HasPrefix(project.Name, name) {
			continue
		}
		bd, err := util.RESTAPIQuery(client, detail(project.Name))
		if err != nil {
			continue
		}
		var detailRes PypiDetailResponse
		if err := json.Unmarshal(bd, &detailRes); err != nil {
			continue
		}
		description := detailRes.Info.Summary
		if len(description) == 0 {
			description = model.NoDescription
		}
		author := detailRes.Info.Author
		if len(author) == 0 {
			author = model.NoAuthor
		}
		result = append(result, model.PyPIPackageResult{Name: project.Name, Description: description, Author: author})
		count++
		if count >= max {
			break
		}
	}

	return result, nil
}
