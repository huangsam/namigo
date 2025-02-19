package main

import (
	"fmt"

	"github.com/huangsam/namigo/pkg/golang"
	"github.com/huangsam/namigo/pkg/npm"
	"github.com/huangsam/namigo/pkg/pypi"
)

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	searchTerm := "hello"

	for _, res := range npm.SearchByScrape(searchTerm) {
		content := fmt.Sprintf("[npm] %s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, res.Description)
		fmt.Println(content)
	}

	for _, res := range golang.SearchByScrape(searchTerm) {
		content := fmt.Sprintf("[golang] %s %s ->\n\t%s", res.Name, res.Path, res.Description)
		fmt.Println(content)
	}

	for _, res := range pypi.SearchByAPI(searchTerm) {
		content := fmt.Sprintf("[pypi] %s", res.Name)
		fmt.Println(content)
	}
}
