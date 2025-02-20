package sub

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/golang"
	"github.com/huangsam/namigo/pkg/npm"
	"github.com/huangsam/namigo/pkg/pypi"
	"github.com/urfave/cli/v2"
)

// portfolio is a collection of result slices.
type portfolio struct {
	results struct {
		golang []model.GoPackageResult
		npm    []model.NPMPackageResult
		pypi   []model.PyPIPackageResult
	}
	wg *sync.WaitGroup
}

// newPortfolio creates a new portfolio instance.
func newPortfolio() *portfolio {
	return &portfolio{wg: &sync.WaitGroup{}}
}

// isEmpty checks if the portfolio has zero results.
func (p *portfolio) isEmpty() bool {
	return len(p.results.npm)+len(p.results.golang)+len(p.results.pypi) == 0
}

// run invokes a function as a goroutine and passes a WaitGroup into it.
func (p *portfolio) run(f func(wg *sync.WaitGroup)) {
	p.wg.Add(1)
	go f(p.wg)
}

// wait blocks the main thread until all runners are complete.
func (p *portfolio) wait() {
	p.wg.Wait()
}

// SearchAction searches term for finding packages.
func SearchAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return errors.New("Provide at least one search term")
	}

	maxResultsToPrint := c.Int("max")

	ptf := newPortfolio()

	errorCount := 0

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := golang.SearchByScrape(searchTerm); err == nil {
			fmt.Println("🟢 Load Golang results")
			ptf.results.golang = searchResults
		} else {
			fmt.Println("🔴 Cannot get Golang results:", err.Error())
			errorCount++
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := npm.SearchByScrape(searchTerm); err == nil {
			fmt.Println("🟢 Load NPM results")
			ptf.results.npm = searchResults
		} else {
			fmt.Println("🔴 Cannot get NPM results:", err.Error())
			errorCount++
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := pypi.SearchByAPI(searchTerm); err == nil {
			fmt.Println("🟢 Load PyPI results")
			ptf.results.pypi = searchResults
		} else {
			fmt.Println("🔴 Cannot get PyPI results:", err.Error())
			errorCount++
		}
	})

	ptf.wait()

	if errorCount > 0 {
		return fmt.Errorf("found %d portfolio errors", errorCount)
	}

	if ptf.isEmpty() {
		fmt.Printf("🌧️ No results...\n")
	} else {
		fmt.Printf("🍺 Prepare results...\n\n")
	}

	time.Sleep(500 * time.Millisecond)

	printResults := func(results any, label string, maxResults int, format func(any) string) {
		switch res := results.(type) {
		case []model.GoPackageResult:
			if len(res) > 0 {
				fmt.Printf("%d %s results found. First %d are:\n", len(res), label, maxResults)
				for i, r := range res {
					if i >= maxResults {
						break
					}
					fmt.Println(format(r))
				}
				fmt.Println()
			}
		case []model.NPMPackageResult:
			if len(res) > 0 {
				fmt.Printf("%d %s results found. First %d are:\n", len(res), label, maxResults)
				for i, r := range res {
					if i >= maxResults {
						break
					}
					fmt.Println(format(r))
				}
				fmt.Println()
			}
		case []model.PyPIPackageResult:
			if len(res) > 0 {
				fmt.Printf("%d %s results found. First %d are:\n", len(res), label, maxResults)
				for i, r := range res {
					if i >= maxResults {
						break
					}
					fmt.Println(format(r))
				}
				fmt.Println()
			}
		}
	}

	formatGo := func(result any) string {
		res := result.(model.GoPackageResult)
		if len(res.Description) > 80 || len(res.Description) == 0 {
			return fmt.Sprintf("\t[golang] %s %s ->\n\t\t%.80s...", res.Name, res.Path, res.Description)
		}
		return fmt.Sprintf("\t[golang] %s %s ->\n\t\t%s", res.Name, res.Path, res.Description)
	}

	formatNPM := func(result any) string {
		res := result.(model.NPMPackageResult)
		if len(res.Description) > 80 || len(res.Description) == 0 {
			return fmt.Sprintf("\t[npm] %s [exact=%v] ->\n\t\t%.80s...", res.Name, res.IsExactMatch, res.Description)
		}
		return fmt.Sprintf("\t[npm] %s [exact=%v] ->\n\t\t%s", res.Name, res.IsExactMatch, res.Description)
	}

	formatPyPI := func(result any) string {
		res := result.(model.PyPIPackageResult)
		if len(res.Description) > 80 || len(res.Description) == 0 {
			return fmt.Sprintf("\t[pypi] %s by %s ->\n\t\t%.80s...", res.Name, res.Author, res.Description)
		}
		return fmt.Sprintf("\t[pypi] %s by %s ->\n\t\t%s", res.Name, res.Author, res.Description)
	}

	printResults(ptf.results.golang, "Golang", maxResultsToPrint, formatGo)
	printResults(ptf.results.npm, "NPM", maxResultsToPrint, formatNPM)
	printResults(ptf.results.pypi, "PyPI", maxResultsToPrint, formatPyPI)

	return nil
}
