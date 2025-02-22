package sub

import (
	"errors"
	"fmt"
	"time"

	"github.com/huangsam/namigo/internal/util"
	"github.com/huangsam/namigo/pkg/search"
	"github.com/huangsam/namigo/pkg/search/dns"
	"github.com/huangsam/namigo/pkg/search/golang"
	"github.com/huangsam/namigo/pkg/search/npm"
	"github.com/huangsam/namigo/pkg/search/pypi"
	"github.com/urfave/cli/v2"
)

const (
	golangLabel = "Golang"
	npmLabel    = "NPM"
	pypiLabel   = "PyPI"
	dnsLabel    = "DNS"
)

var ErrMissingSearchTerm = errors.New("missing search term")

// getOutputMode returns an OutputMode instance.
func getOutputMode(mode string) util.OutputMode {
	switch mode {
	case "text":
		return util.TextMode
	case "json":
		return util.JSONMode
	default:
		return util.TextMode
	}
}

// SearchPackageAction searches for packages.
func SearchPackageAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxResults := c.Int("max")
	outputMode := getOutputMode(c.String("mode"))

	ptf := search.NewSearchPortfolio()

	ptf.Run(func(ptf *search.SearchPortfolio) {
		defer ptf.Done()
		fmt.Printf("🟡 Search for %s results\n", golangLabel)
		if searchResults, err := golang.SearchByScrape(searchTerm, maxResults); err == nil {
			ptf.Results.Golang = searchResults
		} else {
			ptf.Errs.Golang = err
		}
	})

	ptf.Run(func(ptf *search.SearchPortfolio) {
		defer ptf.Done()
		fmt.Printf("🟡 Search for %s results\n", npmLabel)
		if searchResults, err := npm.SearchByScrape(searchTerm, maxResults); err == nil {
			ptf.Results.NPM = searchResults
		} else {
			ptf.Errs.NPM = err
		}
	})

	ptf.Run(func(ptf *search.SearchPortfolio) {
		defer ptf.Done()
		fmt.Printf("🟡 Search for %s results\n", pypiLabel)
		if searchResults, err := pypi.SearchByAPI(searchTerm, maxResults); err == nil {
			ptf.Results.PyPI = searchResults
		} else {
			ptf.Errs.PyPI = err
		}
	})

	ptf.Wait()
	if errs := ptf.Errors(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("💀 Error: %s\n", err)
		}
		return search.ErrPorftolioFailure
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputMode)
	time.Sleep(500 * time.Millisecond)

	f := &search.SearchFormatter{}

	util.PrintResults(ptf.Results.Golang, golangLabel, f.FormatGo, outputMode)
	util.PrintResults(ptf.Results.NPM, npmLabel, f.FormatNPM, outputMode)
	util.PrintResults(ptf.Results.PyPI, pypiLabel, f.FormatPyPI, outputMode)

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

	ptf := search.NewSearchPortfolio()

	ptf.Run(func(ptf *search.SearchPortfolio) {
		defer ptf.Done()
		fmt.Printf("🟡 Search for %s results\n", dnsLabel)
		if probeResults, err := dns.SearchByProbe(searchTerm, maxResults); err == nil {
			ptf.Results.DNS = probeResults
		} else {
			ptf.Errs.DNS = err
		}
	})

	ptf.Wait()
	if errs := ptf.Errors(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("💀 Error: %s\n", err)
		}
		return search.ErrPorftolioFailure
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputMode)
	time.Sleep(500 * time.Millisecond)

	f := &search.SearchFormatter{}

	util.PrintResults(ptf.Results.DNS, dnsLabel, f.FormatDNS, outputMode)
	return nil
}
