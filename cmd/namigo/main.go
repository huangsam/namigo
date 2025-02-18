package main

import (
	"fmt"

	"github.com/huangsam/namigo/pkg/npm"
)

// https://registry.npmjs.com/namigo or scrape
// https://pypi.org/pypi/namigo/json or scrape

// dig namigo.com
// repeat for org, net, io, ai, dev, tech, store, shop, co

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	searchTerm := "hello"

	for _, res := range npm.Search(searchTerm) {
		content := fmt.Sprintf("%s[%v] ->\n\t%s", res.Name, res.IsExactMatch, res.Description)
		fmt.Println(content)
	}
}
