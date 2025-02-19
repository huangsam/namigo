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

// PypiSimpleResponse represents the response from the PyPI simple API.
type PypiSimpleResponse struct {
	Meta struct {
		LastSerial int    `json:"_last-serial"`
		APIVersion string `json:"api-version"`
	} `json:"meta"`
	Projects []struct {
		LastSerial int    `json:"_last-serial"`
		Name       string `json:"name"`
	} `json:"projects"`
}

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(name string) []model.PyPIPackageResult {
	client := &http.Client{Timeout: 5 * time.Second}
	builder := func() (*http.Request, error) {
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

	b, err := util.RESTAPIQuery(client, builder)
	if err != nil {
		log.Fatal(err.Error())
	}

	var simpleRes PypiSimpleResponse
	if err := json.Unmarshal(b, &simpleRes); err != nil {
		log.Fatal(err.Error())
	}

	result := []model.PyPIPackageResult{}
	for _, project := range simpleRes.Projects {
		if strings.HasPrefix(project.Name, name) {
			result = append(result, model.PyPIPackageResult{Name: project.Name})
		}
		if len(result) >= 50 { // That's probably enough!
			break
		}
	}
	return result
}
