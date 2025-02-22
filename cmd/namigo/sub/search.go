package sub

import (
	"fmt"
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
		return ErrMissingSearchTerm
	}
	maxResults := c.Int("max")
	outputMode := getOutputMode(c.String("mode"))

	ptf := newSearchPortfolio()

	ptf.run(func(ptf *searchPortfolio) {
		defer ptf.done()
		fmt.Printf("🟡 Search for %s results\n", golangLabel)
		if searchResults, err := golang.SearchByScrape(searchTerm, maxResults); err == nil {
			ptf.results.golang = searchResults
		} else {
			ptf.errs.golang = err
		}
	})

	ptf.run(func(ptf *searchPortfolio) {
		defer ptf.done()
		fmt.Printf("🟡 Search for %s results\n", npmLabel)
		if searchResults, err := npm.SearchByScrape(searchTerm, maxResults); err == nil {
			ptf.results.npm = searchResults
		} else {
			ptf.errs.npm = err
		}
	})

	ptf.run(func(ptf *searchPortfolio) {
		defer ptf.done()
		fmt.Printf("🟡 Search for %s results\n", pypiLabel)
		if searchResults, err := pypi.SearchByAPI(searchTerm, maxResults); err == nil {
			ptf.results.pypi = searchResults
		} else {
			ptf.errs.pypi = err
		}
	})

	ptf.wait()
	if errs := ptf.errors(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("💀 Error: %s\n", err)
		}
		return ErrPorftolioFailure
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
		return ErrMissingSearchTerm
	}
	maxResults := c.Int("max")
	outputMode := getOutputMode(c.String("mode"))

	ptf := newSearchPortfolio()

	ptf.run(func(ptf *searchPortfolio) {
		defer ptf.done()
		fmt.Printf("🟡 Search for %s results\n", dnsLabel)
		if probeResults, err := dns.SearchByProbe(searchTerm, maxResults); err == nil {
			ptf.results.dns = probeResults
		} else {
			ptf.errs.dns = err
		}
	})

	ptf.wait()
	if errs := ptf.errors(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("💀 Error: %s\n", err)
		}
		return ErrPorftolioFailure
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputMode)
	time.Sleep(500 * time.Millisecond)

	f := &searchFormatter{}

	util.PrintResults(ptf.results.dns, dnsLabel, f.formatDNS, outputMode)
	return nil
}
