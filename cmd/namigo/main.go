package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/huangsam/namigo/internal/util"
)

// https://registry.npmjs.com/namigo or scrape
// https://pypi.org/pypi/namigo/json or scrape

// dig namigo.com
// repeat for org, net, io, ai, dev, tech, store, shop, co

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

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

	// Run document query
	doc.Find("main section").Each(func(i int, section *goquery.Selection) {
		pkg, ok := section.Find("div.bea55649 a").Attr("href")
		if ok {
			fmt.Printf("pkg: %s\n", pkg)
		}

		match := section.Find("div.bea55649 span#pkg-list-exact-match").Text()
		fmt.Printf("match: %s\n", match)

		description := section.Find("p").Text()
		fmt.Printf("description: %s\n", description)

		fmt.Println()
	})
}
