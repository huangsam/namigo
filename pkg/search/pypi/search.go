package pypi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/huangsam/namigo/v2/internal/core"
	"github.com/huangsam/namigo/v2/internal/model"
	"github.com/huangsam/namigo/v2/internal/model/extern"
)

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(client *http.Client, name string, size int) ([]model.PyPIPackage, error) {
	bl, err := core.RESTAPIQuery(client, APIList())
	if err != nil {
		return nil, err
	}

	var listRes extern.PyPIAPIListResponse
	if err := json.Unmarshal(bl, &listRes); err != nil {
		return nil, err
	}

	taskChan := make(chan string)
	go func() {
		for _, project := range listRes.Projects {
			if strings.HasPrefix(project.Name, name) {
				taskChan <- project.Name
			}
		}
		close(taskChan)
	}()

	result := make([]model.PyPIPackage, 0, size) // Pre-allocate with capacity
	errors := []error{}
	var mu sync.Mutex

	core.StartDynamicWorkers(len(listRes.Projects), func() {
		for pkg := range taskChan {
			bd, err := core.RESTAPIQuery(client, APIDetail(pkg))
			if err != nil {
				mu.Lock() // Critical section
				errors = append(errors, err)
				mu.Unlock()
				continue
			}

			var detailRes extern.PyPIAPIDetailResponse
			if err := json.Unmarshal(bd, &detailRes); err != nil {
				mu.Lock() // Critical section
				errors = append(errors, err)
				mu.Unlock()
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

			mu.Lock() // Critical section
			if len(result) < size {
				result = append(result, model.PyPIPackage{Name: pkg, Description: description, Author: author})
			}
			mu.Unlock()
			if len(result) >= size {
				break
			}
		}
	})

	if len(result) == 0 && len(errors) > 0 {
		return result, fmt.Errorf("no results with %d errors", len(errors))
	}
	return result, nil
}
