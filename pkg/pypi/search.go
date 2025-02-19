package pypi

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// PypiListingResponse represents the response from the PyPI simple API.
type PypiListingResponse struct {
	Meta struct {
		LastSerial int    `json:"_last-serial"`
		APIVersion string `json:"api-version"`
	} `json:"meta"`
	Projects []struct {
		LastSerial int    `json:"_last-serial"`
		Name       string `json:"name"`
	} `json:"projects"`
}

// PypiDetailResponse represents the response from the PyPI json API.
type PypiDetailResponse struct {
	Info struct {
		Author      string `json:"author"`
		Description string `json:"description"`
		Summary     string `json:"summary"`
		Version     string `json:"version"`
	}
}

func listing() (*http.Request, error) {
	url := url.URL{
		Scheme: "https",
		Host:   "pypi.org",
		Path:   "simple/",
	}
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.pypi.simple.v1+json")
	return req, nil
}

func detail(pkg string) util.RequestBuilder {
	return func() (*http.Request, error) {
		url := url.URL{
			Scheme: "https",
			Host:   "pypi.org",
			Path:   "pypi/" + pkg + "/json",
		}
		return http.NewRequest("GET", url.String(), nil)
	}
}

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(name string) []model.PyPIPackageResult {
	client := &http.Client{Timeout: 5 * time.Second}

	b, err := util.RESTAPIQuery(client, listing)
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
		for _, project := range simpleRes.Projects {
			if strings.HasPrefix(project.Name, name) {
				taskChan <- project.Name
			}
			if len(result) >= 50 { // That's probably enough!
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
	workerCount := 5
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
