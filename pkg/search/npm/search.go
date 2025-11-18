package npm

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/huangsam/namigo/internal/core"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/model/extern"
)

// SearchByAPI searches for NPM packages by querying registry.npmjs.com.
func SearchByAPI(name string, size int) ([]model.NPMPackage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	return SearchByAPIWithClient(client, APIList(name, size))
}

// SearchByAPIWithClient searches using a custom client and builder.
func SearchByAPIWithClient(client *http.Client, builder core.RequestBuilder) ([]model.NPMPackage, error) {
	bl, err := core.RESTAPIQuery(client, builder)
	if err != nil {
		return nil, err
	}

	var listRes extern.NPMAPIListResponse
	if err := json.Unmarshal(bl, &listRes); err != nil {
		return nil, err
	}

	result := []model.NPMPackage{}
	for _, object := range listRes.Objects {
		pkg := object.Package.Name

		description := object.Package.Description
		if len(description) == 0 {
			description = model.NoDescription
		} else {
			description = strings.TrimSpace(description)
		}

		result = append(result, model.NPMPackage{Name: pkg, Description: description})
	}
	return result, nil
}
