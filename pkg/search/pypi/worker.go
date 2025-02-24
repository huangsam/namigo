package pypi

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/model/extern"
	"github.com/huangsam/namigo/internal/util"
)

// apiWorker runs concurrent logic for PyPI search.
func apiWorker(
	client *http.Client,
	taskChan chan string,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	result *[]model.PyPIPackage,
	errorCount *int,
	maxResults int,
) {
	defer wg.Done()
	for pkg := range taskChan {
		bd, err := util.RESTAPIQuery(client, APIDetail(pkg))
		if err != nil {
			mu.Lock() // Critical section
			*errorCount++
			mu.Unlock()
			continue
		}
		var detailRes extern.PyPIAPIDetailResponse
		if err := json.Unmarshal(bd, &detailRes); err != nil {
			*errorCount++
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
		if len(*result) < maxResults {
			*result = append(*result, model.PyPIPackage{Name: pkg, Description: description, Author: author})
		}
		mu.Unlock()
		if len(*result) >= maxResults {
			break
		}
	}
}
