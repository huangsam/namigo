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

// SearchPackageAction searches term for finding packages.
func SearchPackageAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return errors.New("Provide at least one search term")
	}
	maxResults := c.Int("max")

	ptf := newSearchPortfolio()
	ptfErrorCount := 0

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := golang.SearchByScrape(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load Golang results")
			ptf.results.golang = searchResults
		} else {
			fmt.Println("🔴 Cannot get Golang results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := npm.SearchByScrape(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load NPM results")
			ptf.results.npm = searchResults
		} else {
			fmt.Println("🔴 Cannot get NPM results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := pypi.SearchByAPI(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load PyPI results")
			ptf.results.pypi = searchResults
		} else {
			fmt.Println("🔴 Cannot get PyPI results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.wait()

	if ptfErrorCount == ptf.count() {
		return errors.New("portfolio collection failed")
	}

	if ptf.isEmpty() {
		fmt.Printf("🌧️ No results\n")
	} else {
		fmt.Printf("🍺 Prepare results\n\n")
	}

	time.Sleep(500 * time.Millisecond)

	f := &searchFormatter{}

	util.PrintResults(ptf.results.golang, "Golang", f.formatGo)
	util.PrintResults(ptf.results.npm, "NPM", f.formatNPM)
	util.PrintResults(ptf.results.pypi, "PyPI", f.formatPyPI)

	return nil
}

// SearchDNSAction searches term for finding DNS records.
func SearchDNSAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return errors.New("Provide at least one search term")
	}
	maxResults := c.Int("max")

	ptf := newSearchPortfolio()
	ptfErrorCount := 0

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if probeResults, err := dns.SearchByProbe(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load DNS results")
			ptf.results.dns = probeResults
		} else {
			fmt.Println("🔴 Cannot get DNS results:", err.Error())
			ptfErrorCount++
		}
	})

	ptf.wait()

	if ptfErrorCount == ptf.count() {
		return errors.New("portfolio collection failed")
	}

	if ptf.isEmpty() {
		fmt.Printf("🌧️ No results\n")
	} else {
		fmt.Printf("🍺 Prepare results\n\n")
	}

	time.Sleep(500 * time.Millisecond)

	f := &searchFormatter{}

	util.PrintResults(ptf.results.dns, "DNS", f.formatDNS)

	return nil
}
