package npm

import (
	"strings"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/model/extern"
)

// apiWorker runs serial logic for NPM search.
func apiWorker(listRes *extern.NPMAPIListResponse) ([]model.NPMPackage, error) {
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
