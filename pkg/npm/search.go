package npm

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

func Search(name string) []model.NPMResult {
	// Setup document query
	client := &http.Client{Timeout: 5 * time.Second}
	builder := func() (*http.Request, error) {
		params := url.Values{"q": []string{"hello"}}
		url := url.URL{
			Scheme:   "https",
			Host:     "www.npmjs.com",
			Path:     "search",
			RawQuery: params.Encode(),
		}
		return http.NewRequest("GET", url.String(), nil)
	}
	pipe := util.NewDocumentPipeline(client, builder)
	doc, err := pipe.Execute()
	if err != nil {
		log.Fatal(err)
	}

	result := []model.NPMResult{}

	// Run document query
	doc.Find("main section").Each(func(i int, section *goquery.Selection) {
		pkg := section.Find("h3.db7ee1ac").Text()
		match := section.Find("div.bea55649 span#pkg-list-exact-match").Text()
		description := section.Find("p").Text()

		result = append(result, model.NPMResult{
			Name:         pkg,
			Description:  description,
			IsExactMatch: len(match) > 0,
		})
	})

	return result
}
