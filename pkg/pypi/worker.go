package pypi

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

func worker(
	client *http.Client,
	taskChan chan string,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	result *[]model.PyPIPackageResult,
	errorCount *int,
	maxResults int,
) {
	defer wg.Done()
	for pkg := range taskChan {
		bd, err := util.RESTAPIQuery(client, detail(pkg))
		if err != nil {
			*errorCount++
			continue
		}
		var detailRes PypiDetailResponse
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
		mu.Lock()
		if len(*result) < maxResults {
			*result = append(*result, model.PyPIPackageResult{Name: pkg, Description: description, Author: author})
		}
		mu.Unlock()
		if len(*result) >= maxResults {
			break
		}
	}
}
