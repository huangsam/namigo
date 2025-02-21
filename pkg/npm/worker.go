package npm

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
)

func worker(doc *goquery.Document, result *[]model.NPMPackageResult, maxResults int) {
	doc.Find("main section").Each(func(i int, section *goquery.Selection) {
		if len(*result) >= maxResults {
			return
		}

		pkg := section.Find("h3").Text()

		match := section.Find("span#pkg-list-exact-match").Text()

		description := section.Find("p").Text()
		if len(description) == 0 {
			description = model.NoDescription
		} else {
			description = strings.Trim(description, " \n\t")
		}

		*result = append(*result, model.NPMPackageResult{
			Name:         pkg,
			Description:  description,
			IsExactMatch: len(match) > 0,
		})
	})
}
