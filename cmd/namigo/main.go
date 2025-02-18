package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// https://registry.npmjs.com/namigo or scrape
// https://pypi.org/pypi/namigo/json or scrape

// dig namigo.com
// repeat for org, net, io, ai, dev, tech, store, shop, co

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	res, err := http.Get("https://www.npmjs.com/search?q=hello")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
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
