package sub

import (
	"errors"
	"fmt"
	"time"

	"github.com/huangsam/namigo/internal/cmd"
	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
	"github.com/huangsam/namigo/pkg/search/dns"
	"github.com/huangsam/namigo/pkg/search/email"
	"github.com/huangsam/namigo/pkg/search/golang"
	"github.com/huangsam/namigo/pkg/search/npm"
	"github.com/huangsam/namigo/pkg/search/pypi"
	"github.com/urfave/cli/v2"
)

var ErrMissingSearchTerm = errors.New("missing search term")

// SearchPackageAction searches for packages.
func SearchPackageAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxSize := c.Int("size")
	outputFormat := model.GetOutputFormat(c.String("format"))

	ptf := search.NewPortfolio()

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.Golang.Label())
		if searchResults, err := golang.SearchByScrape(searchTerm, maxSize); err == nil {
			ptf.Res.Golang = searchResults
		} else {
			ptf.Err.Golang = err
		}
	})

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.NPM.Label())
		if searchResults, err := npm.SearchByAPI(searchTerm, maxSize); err == nil {
			ptf.Res.NPM = searchResults
		} else {
			ptf.Err.NPM = err
		}
	})

	ptf.Run(func(ptf *search.Portfolio) {
		defer ptf.Done()
		fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.PyPI.Label())
		if searchResults, err := pypi.SearchByAPI(searchTerm, maxSize); err == nil {
			ptf.Res.PyPI = searchResults
		} else {
			ptf.Err.PyPI = err
		}
	})

	ptf.Wait()
	if err := cmd.CheckPortfolio(ptf); err != nil {
		return err
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputFormat)
	time.Sleep(500 * time.Millisecond)
	cmd.DisplayResults(ptf.Res.Golang, &ptf.Fmt.Golang, outputFormat)
	cmd.DisplayResults(ptf.Res.NPM, &ptf.Fmt.NPM, outputFormat)
	cmd.DisplayResults(ptf.Res.PyPI, &ptf.Fmt.PyPI, outputFormat)

	return nil
}

// SearchDNSAction searches for DNS records.
func SearchDNSAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxSize := c.Int("size")
	outputFormat := model.GetOutputFormat(c.String("format"))

	ptf := search.NewPortfolio()

	fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.DNS.Label())
	if searchResults, err := dns.SearchByProbe(searchTerm, maxSize); err == nil {
		ptf.Res.DNS = searchResults
	} else {
		ptf.Err.DNS = err
	}

	if err := cmd.CheckPortfolio(ptf); err != nil {
		return err
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputFormat)
	time.Sleep(500 * time.Millisecond)
	cmd.DisplayResults(ptf.Res.DNS, &ptf.Fmt.DNS, outputFormat)

	return nil
}

// SearchEmailAction searches for email records.
func SearchEmailAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxSize := c.Int("size")
	outputFormat := model.GetOutputFormat(c.String("format"))

	ptf := search.NewPortfolio()

	fmt.Printf("🔍 Search for %s results\n", ptf.Fmt.Email.Label())
	if searchResults, err := email.SearchByProbe(searchTerm, maxSize); err == nil {
		ptf.Res.Email = searchResults
	} else {
		ptf.Err.Email = err
	}

	if err := cmd.CheckPortfolio(ptf); err != nil {
		return err
	}

	fmt.Printf("🍺 Prepare %s results\n\n", outputFormat)
	time.Sleep(500 * time.Millisecond)
	cmd.DisplayResults(ptf.Res.Email, &ptf.Fmt.Email, outputFormat)

	return nil
}
