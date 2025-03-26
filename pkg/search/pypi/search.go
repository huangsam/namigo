package pypi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/model/extern"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByAPI searches for PyPI packages by querying pypi.org.
func SearchByAPI(name string, size int) ([]model.PyPIPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	bl, err := util.RESTAPIQuery(client, APIList())
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

	result := []model.PyPIPackage{}
	errors := []error{}
	var mu sync.Mutex

	util.StartCommonWorkers(func() {
		apiWorker(client, taskChan, &mu, &result, &errors, size)
	})

	if len(result) == 0 && len(errors) > 0 {
		return result, fmt.Errorf("no results with %d errors", len(errors))
	}
	return result, nil
}
