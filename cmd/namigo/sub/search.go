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
	outputFormat := model.GetOutputFormat(c.String("format"))

	ptf := search.NewPortfolio()

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.Golang.Label())
		if searchResults, err := golang.SearchByScrape(searchTerm, maxResults); err == nil {
			ptf.Res.Golang = searchResults
		} else {
			ptf.Err.Golang = err
		}
	})

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.NPM.Label())
		if searchResults, err := npm.SearchByAPI(searchTerm, maxResults); err == nil {
			ptf.Res.NPM = searchResults
		} else {
			ptf.Err.NPM = err
		}
	})

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.PyPI.Label())
		if searchResults, err := pypi.SearchByAPI(searchTerm, maxResults); err == nil {
			ptf.Res.PyPI = searchResults
		} else {
			ptf.Err.PyPI = err
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

	fmt.Printf("🍺 Prepare %s results\n\n", outputFormat)
	time.Sleep(500 * time.Millisecond)

	displayResults(ptf.Res.Golang, &ptf.Fmt.Golang, outputFormat)
	displayResults(ptf.Res.NPM, &ptf.Fmt.NPM, outputFormat)
	displayResults(ptf.Res.PyPI, &ptf.Fmt.PyPI, outputFormat)

	return nil
}

// SearchDNSAction searches for DNS records.
func SearchDNSAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxResults := c.Int("max")
	outputFormat := model.GetOutputFormat(c.String("format"))

	ptf := search.NewPortfolio()

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.DNS.Label())
		if probeResults, err := dns.SearchByProbe(searchTerm, maxResults); err == nil {
			ptf.Res.DNS = probeResults
		} else {
			ptf.Err.DNS = err
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

	fmt.Printf("🍺 Prepare %s results\n\n", outputFormat)
	time.Sleep(500 * time.Millisecond)

	displayResults(ptf.Res.DNS, &ptf.Fmt.DNS, outputFormat)
	return nil
}
