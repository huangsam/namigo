package npm

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/model/extern"
)

// docWorker runs serial logic for NPM search.
func docWorker(doc *goquery.Document, maxResults int) []model.NPMPackage {
	result := []model.NPMPackage{}
	doc.Find("main section").Each(func(i int, section *goquery.Selection) {
		if len(result) >= maxResults {
			return
		}

		pkg := section.Find("h3").Text()

		description := section.Find("p").Text()
		if len(description) == 0 {
			description = model.NoDescription
		} else {
			description = strings.TrimSpace(description)
		}

		result = append(result, model.NPMPackage{Name: pkg, Description: description})
	})
	return result
}

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
