package npm

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/model/extern"
	"github.com/huangsam/namigo/internal/util"
)

// SearchByAPI searches for NPM packages by querying registry.npmjs.com.
func SearchByAPI(name string, size int) ([]model.NPMPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	bl, err := util.RESTAPIQuery(client, APIList(name, size))
	if err != nil {
		return nil, err
	}

	var listRes extern.NPMAPIListResponse
	if err := json.Unmarshal(bl, &listRes); err != nil {
		return nil, err
	}

	return apiWorker(&listRes)
}
