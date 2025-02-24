package golang

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
)

// docWorker runs serial logic for Golang search.
func docWorker(doc *goquery.Document, maxResults int, name string) []model.GoPackage {
	result := []model.GoPackage{}
	doc.Find(".SearchSnippet").Each(func(i int, section *goquery.Selection) {
		if len(result) >= maxResults {
			return
		}

		content := strings.Fields(section.Find("h2").Text())
		pkg, path := content[0], strings.Trim(content[1], "()")
		if !strings.Contains(pkg, name) && !strings.Contains(path, name) {
			return
		}

		description := strings.Trim(section.Find("p").Text(), " \n")
		if len(description) == 0 {
			description = model.NoDescription
		}

		result = append(result, model.GoPackage{
			Name:        pkg,
			Path:        path,
			Description: description,
		})
	})
	return result
}
