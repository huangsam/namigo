package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/golang"
	"github.com/huangsam/namigo/pkg/npm"
	"github.com/huangsam/namigo/pkg/pypi"
)

// maxResultsToPrint is the max number of results to print for each result collection.
const maxResultsToPrint = 10

// portfolio is a collection of result slices.
type portfolio struct {
	npmResults    []model.NPMPackageResult
	golangResults []model.GoPackageResult
	pypiResults   []model.PyPIPackageResult
}

func (p *portfolio) isEmpty() bool {
	return len(p.npmResults)+len(p.golangResults)+len(p.pypiResults) == 0
}

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	searchTerm := "hello"

	var ptf portfolio

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if searchResults, err := npm.SearchByScrape(searchTerm); err == nil {
			fmt.Println("🟢 Load NPM results")
			ptf.npmResults = searchResults
		} else {
			fmt.Println("🔴 Cannot get NPM results:", err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if searchResults, err := golang.SearchByScrape(searchTerm); err == nil {
			fmt.Println("🟢 Load Golang results")
			ptf.golangResults = searchResults
		} else {
			fmt.Println("🔴 Cannot get Golang results:", err.Error())
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if searchResults, err := pypi.SearchByAPI(searchTerm); err == nil {
			fmt.Println("🟢 Load PyPI results")
			ptf.pypiResults = searchResults
		} else {
			fmt.Println("🔴 Cannot get PyPI results:", err.Error())
		}
	}()

	wg.Wait()

	if ptf.isEmpty() {
		fmt.Printf("👎 No results...\n")
	} else {
		fmt.Printf("🍺 Prepare results...\n\n")
	}

	time.Sleep(500 * time.Millisecond)

	for i, res := range ptf.npmResults {
		if i >= maxResultsToPrint {
			break
		}
		content := fmt.Sprintf("[npm] %s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, res.Description)
		fmt.Println(content)
	}

	for i, res := range ptf.golangResults {
		if i >= maxResultsToPrint {
			break
		}
		content := fmt.Sprintf("[golang] %s %s ->\n\t%s", res.Name, res.Path, res.Description)
		fmt.Println(content)
	}

	for i, res := range ptf.pypiResults {
		if i >= maxResultsToPrint {
			break
		}
		content := fmt.Sprintf("[pypi] %s ->\n\t%s", res.Name, res.Description)
		fmt.Println(content)
	}
}
