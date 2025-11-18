package sub

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
	"github.com/huangsam/namigo/pkg/search/dns"
	"github.com/huangsam/namigo/pkg/search/email"
	"github.com/huangsam/namigo/pkg/search/golang"
	"github.com/huangsam/namigo/pkg/search/npm"
	"github.com/huangsam/namigo/pkg/search/pypi"
	"github.com/urfave/cli/v2"
)

// ErrMissingSearchTerm is returned when the search term is missing.
var ErrMissingSearchTerm = errors.New("missing search term")

// SearchRunner encapsulates the logic for running searches.
type SearchRunner struct {
	output io.Writer
}

// NewSearchRunner creates a new SearchRunner instance.
func NewSearchRunner(output io.Writer) *SearchRunner {
	return &SearchRunner{output: output}
}

// RunPackageSearch executes a package search with the given parameters.
func (sr *SearchRunner) RunPackageSearch(searchTerm string, maxSize int, outputFormat search.FormatOption) error {
	ptf := search.NewSearchPortfolio(outputFormat, sr.output)

	ptf.Register(func() (model.SearchResult, error) {
		key := model.GoKey
		_, _ = fmt.Fprintf(sr.output, "🔍 Search for %s results\n", key)
		values, err := golang.SearchByScrape(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := make([]model.SearchRecord, len(values))
		for i := range values {
			records[i] = &values[i]
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	ptf.Register(func() (model.SearchResult, error) {
		key := model.NPMKey
		_, _ = fmt.Fprintf(sr.output, "🔍 Search for %s results\n", key)
		values, err := npm.SearchByAPI(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := make([]model.SearchRecord, len(values))
		for i := range values {
			records[i] = &values[i]
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	ptf.Register(func() (model.SearchResult, error) {
		key := model.PyPIKey
		_, _ = fmt.Fprintf(sr.output, "🔍 Search for %s results\n", key)
		values, err := pypi.SearchByAPI(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := make([]model.SearchRecord, len(values))
		for i := range values {
			records[i] = &values[i]
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	if err := ptf.Run(); err != nil {
		return err
	}
	ptf.Display()

	return nil
}

// SearchPackageAction searches for packages.
func SearchPackageAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxSize := c.Int("size")
	outputFormat := search.GetFormatOption(c.String("format"))

	runner := NewSearchRunner(os.Stdout)
	return runner.RunPackageSearch(searchTerm, maxSize, outputFormat)
}

// RunDNSSearch executes a DNS search with the given parameters.
func (sr *SearchRunner) RunDNSSearch(searchTerm string, maxSize int, outputFormat search.FormatOption) error {
	ptf := search.NewSearchPortfolio(outputFormat, sr.output)

	ptf.Register(func() (model.SearchResult, error) {
		key := model.DNSKey
		_, _ = fmt.Fprintf(sr.output, "🔍 Search for %s results\n", key)
		values, err := dns.SearchByProbe(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := make([]model.SearchRecord, len(values))
		for i := range values {
			records[i] = &values[i]
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	if err := ptf.Run(); err != nil {
		return err
	}
	ptf.Display()

	return nil
}

// SearchDNSAction searches for DNS records.
func SearchDNSAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxSize := c.Int("size")
	outputFormat := search.GetFormatOption(c.String("format"))

	runner := NewSearchRunner(os.Stdout)
	return runner.RunDNSSearch(searchTerm, maxSize, outputFormat)
}

// RunEmailSearch executes an email search with the given parameters.
func (sr *SearchRunner) RunEmailSearch(searchTerm string, maxSize int, outputFormat search.FormatOption) error {
	ptf := search.NewSearchPortfolio(outputFormat, sr.output)

	ptf.Register(func() (model.SearchResult, error) {
		key := model.EmailKey
		_, _ = fmt.Fprintf(sr.output, "🔍 Search for %s results\n", key)
		values, err := email.SearchByProbe(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := make([]model.SearchRecord, len(values))
		for i := range values {
			records[i] = &values[i]
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	if err := ptf.Run(); err != nil {
		return err
	}
	ptf.Display()

	return nil
}

// SearchEmailAction searches for email records.
func SearchEmailAction(c *cli.Context) error {
	searchTerm := c.Args().First()
	if len(searchTerm) == 0 {
		return ErrMissingSearchTerm
	}
	maxSize := c.Int("size")
	outputFormat := search.GetFormatOption(c.String("format"))

	runner := NewSearchRunner(os.Stdout)
	return runner.RunEmailSearch(searchTerm, maxSize, outputFormat)
}
