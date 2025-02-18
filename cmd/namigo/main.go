package main

import (
	"fmt"

	"github.com/huangsam/namigo/pkg/npm"
)

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	searchTerm := "hello"

	for _, res := range npm.Search(searchTerm) {
		content := fmt.Sprintf("%s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, res.Description)
		fmt.Println(content)
	}
}
