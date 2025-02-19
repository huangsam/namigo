package main

import (
	"fmt"

	"github.com/huangsam/namigo/pkg/golang"
	"github.com/huangsam/namigo/pkg/npm"
	"github.com/huangsam/namigo/pkg/pypi"
)

const maxResultsToPrint = 10

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	searchTerm := "hello"

	for i, res := range npm.SearchByScrape(searchTerm) {
		if i >= maxResultsToPrint {
			break
		}
		content := fmt.Sprintf("[npm] %s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, res.Description)
		fmt.Println(content)
	}

	for i, res := range golang.SearchByScrape(searchTerm) {
		if i >= maxResultsToPrint {
			break
		}
		content := fmt.Sprintf("[golang] %s %s ->\n\t%s", res.Name, res.Path, res.Description)
		fmt.Println(content)
	}

	for i, res := range pypi.SearchByAPI(searchTerm) {
		if i >= maxResultsToPrint {
			break
		}
		content := fmt.Sprintf("[pypi] %s", res.Name)
		fmt.Println(content)
	}
}
