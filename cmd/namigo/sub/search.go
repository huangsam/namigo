package sub

import (
	"errors"
	"fmt"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
	"github.com/huangsam/namigo/pkg/search/dns"
	"github.com/huangsam/namigo/pkg/search/golang"
	"github.com/huangsam/namigo/pkg/search/npm"
	"github.com/huangsam/namigo/pkg/search/pypi"
	"github.com/urfave/cli/v2"
)

var (
	ErrMissingSearchTerm = errors.New("missing search term")
	ErrPorftolioEmpty    = errors.New("portfolio collection empty")
	ErrPorftolioFailure  = errors.New("portfolio collection failure")
)

// SearchPackageAction searches for packages.
func SearchPackageAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxResults := c.Int("max")
	outputMode := model.GetOutputMode(c.String("mode"))

	ptf := search.NewPortfolio()

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Formats.Golang.Label())
		if searchResults, err := golang.SearchByScrape(searchTerm, maxResults); err == nil {
			ptf.Results.Golang = searchResults
		} else {
			ptf.Errs.Golang = err
		}
	})

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Formats.NPM.Label())
		if searchResults, err := npm.SearchByAPI(searchTerm, maxResults); err == nil {
			ptf.Results.NPM = searchResults
		} else {
			ptf.Errs.NPM = err
		}
	})

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Formats.PyPI.Label())
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
		return ErrPorftolioFailure
	} else if ptf.Size() == 0 {
		return ErrPorftolioEmpty
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputMode)
	time.Sleep(500 * time.Millisecond)

	displayResults(ptf.Results.Golang, &ptf.Formats.Golang, outputMode)
	displayResults(ptf.Results.NPM, &ptf.Formats.NPM, outputMode)
	displayResults(ptf.Results.PyPI, &ptf.Formats.PyPI, outputMode)

	return nil
}

// SearchDNSAction searches for DNS records.
func SearchDNSAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxResults := c.Int("max")
	outputMode := model.GetOutputMode(c.String("mode"))

	ptf := search.NewPortfolio()

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Formats.DNS.Label())
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
		return ErrPorftolioFailure
	} else if ptf.Size() == 0 {
		return ErrPorftolioEmpty
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputMode)
	time.Sleep(500 * time.Millisecond)

	displayResults(ptf.Results.DNS, &ptf.Formats.DNS, outputMode)
	return nil
}
