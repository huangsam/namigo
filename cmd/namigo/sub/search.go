package sub

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/huangsam/namigo/internal/util"
	"github.com/huangsam/namigo/pkg/dns"
	"github.com/huangsam/namigo/pkg/golang"
	"github.com/huangsam/namigo/pkg/npm"
	"github.com/huangsam/namigo/pkg/pypi"
	"github.com/urfave/cli/v2"
)

// SearchAction searches term for finding packages.
func SearchAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return errors.New("Provide at least one search term")
	}

	ptf := newSearchPortfolio()
	ptfErrorCount := 0

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := golang.SearchByScrape(searchTerm); err == nil {
			fmt.Println("🟢 Load Golang results")
			ptf.results.golang = searchResults
		} else {
			fmt.Println("🔴 Cannot get Golang results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := npm.SearchByScrape(searchTerm); err == nil {
			fmt.Println("🟢 Load NPM results")
			ptf.results.npm = searchResults
		} else {
			fmt.Println("🔴 Cannot get NPM results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := pypi.SearchByAPI(searchTerm); err == nil {
			fmt.Println("🟢 Load PyPI results")
			ptf.results.pypi = searchResults
		} else {
			fmt.Println("🔴 Cannot get PyPI results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if probeResults, err := dns.SearchByProbe(searchTerm); err == nil {
			fmt.Println("🟢 Load DNS results")
			ptf.results.DNS = probeResults
		} else {
			fmt.Println("🔴 Cannot get DNS results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.wait()

	if ptfErrorCount > 0 {
		return fmt.Errorf("found %d portfolio errors", ptfErrorCount)
	}

	if ptf.isEmpty() {
		fmt.Printf("🌧️ No results\n")
	} else {
		fmt.Printf("🍺 Prepare results\n\n")
	}

	time.Sleep(500 * time.Millisecond)

	f := &searchFormatter{}

	maxResultCount := c.Int("max")
	util.PrintResults(ptf.results.golang, "Golang", maxResultCount, f.formatGo)
	util.PrintResults(ptf.results.npm, "NPM", maxResultCount, f.formatNPM)
	util.PrintResults(ptf.results.pypi, "PyPI", maxResultCount, f.formatPyPI)
	util.PrintResults(ptf.results.DNS, "DNS", maxResultCount, f.formatDNS)

	return nil
}
