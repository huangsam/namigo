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
		if searchResults, err := golang.SearchByScrape(searchTerm, maxResults); err == nil {
			fmt.Printf("🟢 Load %s results\n", golangLabel)
			ptf.results.golang = searchResults
		} else {
			fmt.Printf("🔴 Cannot get %s results: %s\n", golangLabel, err)
			ptf.errs.golang = err
		}
	})

	ptf.run(func(ptf *searchPortfolio) {
		defer ptf.done()
		if searchResults, err := npm.SearchByScrape(searchTerm, maxResults); err == nil {
			fmt.Printf("🟢 Load %s results\n", npmLabel)
			ptf.results.npm = searchResults
		} else {
			fmt.Printf("🔴 Cannot get %s results: %s\n", npmLabel, err)
			ptf.errs.npm = err
		}
	})

	ptf.run(func(ptf *searchPortfolio) {
		defer ptf.done()
		if searchResults, err := pypi.SearchByAPI(searchTerm, maxResults); err == nil {
			fmt.Printf("🟢 Load %s results\n", pypiLabel)
			ptf.results.pypi = searchResults
		} else {
			fmt.Printf("🔴 Cannot get %s results: %s\n", pypiLabel, err)
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
		if probeResults, err := dns.SearchByProbe(searchTerm, maxResults); err == nil {
			fmt.Printf("🟢 Load %s results\n", dnsLabel)
			ptf.results.dns = probeResults
		} else {
			fmt.Printf("🔴 Cannot get %s results: %s\n", dnsLabel, err)
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
