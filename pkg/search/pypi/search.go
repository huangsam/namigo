package pypi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

const goroutineCount = 4

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(name string, max int) ([]model.PyPIPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	bl, err := util.RESTAPIQuery(client, Listing())
	if err != nil {
		return []model.PyPIPackage{}, err
	}

	var listingRes PypiListingResponse
	if err := json.Unmarshal(bl, &listingRes); err != nil {
		return []model.PyPIPackage{}, err
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

	result := []model.PyPIPackage{}
	errorCount := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go apiWorker(client, taskChan, &wg, &mu, &result, &errorCount, max)
	}

	wg.Wait()
	if len(result) == 0 && errorCount > 0 {
		return result, fmt.Errorf("no results with %d errors", errorCount)
	}
	return result, nil
}
