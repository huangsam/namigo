package pypi

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// workerCount is the number of workers used to grab package details.
const workerCount = 5

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(name string) []model.PyPIPackageResult {
	client := &http.Client{Timeout: 5 * time.Second}

	b, err := util.RESTAPIQuery(client, listing())
	if err != nil {
		log.Fatal(err.Error())
	}

	var simpleRes PypiListingResponse
	if err := json.Unmarshal(b, &simpleRes); err != nil {
		log.Fatal(err.Error())
	}

	result := []model.PyPIPackageResult{}

	taskChan := make(chan string)

	go func() {
		count := 0
		for _, project := range simpleRes.Projects {
			if strings.HasPrefix(project.Name, name) {
				taskChan <- project.Name
				count++
			}
			if count >= 50 { // That's probably enough!
				break
			}
		}
		close(taskChan)
	}()

	worker := func() {
		for item := range taskChan {
			bd, err := util.RESTAPIQuery(client, detail(item))
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
			result = append(result, model.PyPIPackageResult{Name: item, Description: description})
		}
	}

	doneChan := make(chan struct{})
	for i := 0; i < workerCount; i++ {
		go func() {
			worker()
			doneChan <- struct{}{}
		}()
	}
	for i := 0; i < workerCount; i++ {
		<-doneChan
	}

	return result
}
