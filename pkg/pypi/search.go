package pypi

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(name string, max int) ([]model.PyPIPackageResult, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	bl, err := util.RESTAPIQuery(client, listing())
	if err != nil {
		return []model.PyPIPackageResult{}, err
	}

	var listingRes PypiListingResponse
	if err := json.Unmarshal(bl, &listingRes); err != nil {
		return []model.PyPIPackageResult{}, err
	}

	taskChan := make(chan string)
	go func() {
		for _, project := range listingRes.Projects {
			if strings.HasPrefix(project.Name, name) {
				taskChan <- project.Name
			}
		}
		close(taskChan)
	}()

	result := []model.PyPIPackageResult{}
	resultCount := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for pkg := range taskChan {
				bd, err := util.RESTAPIQuery(client, detail(pkg))
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
				mu.Lock()
				if resultCount < max {
					result = append(result, model.PyPIPackageResult{Name: pkg, Description: description, Author: author})
					resultCount++
				}
				mu.Unlock()
				if resultCount >= max {
					break
				}
			}
		}()
	}

	wg.Wait()
	return result, nil
}
