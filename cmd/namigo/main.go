package main

import (
	"fmt"

	"github.com/huangsam/namigo/pkg/npm"
)

// JavaScript
// https://registry.npmjs.com/namigo
// https://registry.npmjs.com/-/v1/search?text=namigo
// scrape

// Python
// https://pypi.org/pypi/namigo/json
// scrape

// Rust
// https://crates.io/api/v1/crates?page=1&per_page=10&q=height
// https://crates.io/api/v1/crates/anipwatch
// scrape

// Go
// https://index.golang.org/index
// scrape

// DNS
// dig namigo.com
// whois namigo.com
// repeat for org, net, io, ai, dev, tech, store, shop, co

// Email
// lookup namigo@gmail.com
// repeat for the DNS permutations above

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	searchTerm := "hello"

	for _, res := range npm.Search(searchTerm) {
		content := fmt.Sprintf("%s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, res.Description)
		fmt.Println(content)
	}
}
