package sub

import (
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

const (
	golangLabel = "Golang"
	npmLabel    = "NPM"
	pypiLabel   = "PyPI"
	dnsLabel    = "DNS"
)

// SearchPackageAction searches for packages.
func SearchPackageAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingTerm
	}
	maxResults := c.Int("max")
	outputMode := getOutputMode(c.String("mode"))

	ptf := newSearchPortfolio()
	ptfErrorMap := newErrorMap()

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := golang.SearchByScrape(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load Golang results")
			ptf.results.golang = searchResults
		} else {
			fmt.Println("🔴 Cannot get Golang results:", err.Error())
			ptfErrorMap[golangLabel] = err
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := npm.SearchByScrape(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load NPM results")
			ptf.results.npm = searchResults
		} else {
			fmt.Println("🔴 Cannot get NPM results:", err.Error())
			ptfErrorMap[npmLabel] = err
		}
	})

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if searchResults, err := pypi.SearchByAPI(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load PyPI results")
			ptf.results.pypi = searchResults
		} else {
			fmt.Println("🔴 Cannot get PyPI results:", err.Error())
			ptfErrorMap[pypiLabel] = err
		}
	})

	ptf.wait()
	if err := ptfErrorMap.aggregate(); err != nil {
		return err
	} else if ptf.isEmpty() {
		fmt.Println("🌧️ No results")
		return ErrPorftolioEmpty
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputMode)
	time.Sleep(500 * time.Millisecond)

	f := &searchFormatter{}

	util.PrintResults(ptf.results.golang, golangLabel, f.formatGo, outputMode)
	util.PrintResults(ptf.results.npm, npmLabel, f.formatNPM, outputMode)
	util.PrintResults(ptf.results.pypi, pypiLabel, f.formatPyPI, outputMode)

	return nil
}

// SearchDNSAction searches for DNS records.
func SearchDNSAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingTerm
	}
	maxResults := c.Int("max")
	outputMode := getOutputMode(c.String("mode"))

	ptf := newSearchPortfolio()
	ptfErrorMap := newErrorMap()

	ptf.run(func(wg *sync.WaitGroup) {
		defer wg.Done()
		if probeResults, err := dns.SearchByProbe(searchTerm, maxResults); err == nil {
			fmt.Println("🟢 Load DNS results")
			ptf.results.dns = probeResults
		} else {
			fmt.Println("🔴 Cannot get DNS results:", err.Error())
			ptfErrorMap[dnsLabel] = err
		}
	})

	ptf.wait()
	if err := ptfErrorMap.aggregate(); err != nil {
		return err
	} else if ptf.isEmpty() {
		fmt.Println("🌧️ No results")
		return ErrPorftolioEmpty
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputMode)
	time.Sleep(500 * time.Millisecond)

	f := &searchFormatter{}

	util.PrintResults(ptf.results.dns, dnsLabel, f.formatDNS, outputMode)
	return nil
}
