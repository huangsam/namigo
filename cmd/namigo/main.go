package main

import (
	"fmt"
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

func main() {
	fmt.Println("Hello Namigo 🐶")
	fmt.Println()

	searchTerm := "hello"

	var ptf portfolio

	ch := make(chan struct{})
	defer close(ch)

	go func() {
		ptf.npmResults = npm.SearchByScrape(searchTerm)
		ch <- struct{}{}
	}()

	go func() {
		ptf.golangResults = golang.SearchByScrape(searchTerm)
		ch <- struct{}{}
	}()

	go func() {
		ptf.pypiResults = pypi.SearchByAPI(searchTerm)
		ch <- struct{}{}
	}()

	fmt.Printf("Loading")
	for i := 0; i < 3; i++ {
		<-ch
		fmt.Printf(".")
	}
	fmt.Println("done!")
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
