package sub

import (
	"errors"
	"fmt"

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
	outputFormat := search.GetFormatOption(c.String("format"))

	ptf := search.NewSearchPortfolio(outputFormat)

	ptf.Register(func(sp *search.SearchPortfolio) (model.SearchResult, error) {
		key := model.GoKey
		fmt.Printf("🔍 Search for %s results\n", key)
		values, err := golang.SearchByScrape(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := []model.SearchRecord{}
		for _, value := range values {
			records = append(records, &value)
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	ptf.Register(func(sp *search.SearchPortfolio) (model.SearchResult, error) {
		key := model.NPMKey
		fmt.Printf("🔍 Search for %s results\n", key)
		values, err := npm.SearchByAPI(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := []model.SearchRecord{}
		for _, value := range values {
			records = append(records, &value)
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	ptf.Register(func(sp *search.SearchPortfolio) (model.SearchResult, error) {
		key := model.PyPIKey
		fmt.Printf("🔍 Search for %s results\n", key)
		values, err := pypi.SearchByAPI(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := []model.SearchRecord{}
		for _, value := range values {
			records = append(records, &value)
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

	ptf := search.NewSearchPortfolio(outputFormat)

	ptf.Register(func(sp *search.SearchPortfolio) (model.SearchResult, error) {
		key := model.DNSKey
		fmt.Printf("🔍 Search for %s results\n", key)
		values, err := dns.SearchByProbe(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := []model.SearchRecord{}
		for _, value := range values {
			records = append(records, &value)
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

	ptf := search.NewSearchPortfolio(outputFormat)

	ptf.Register(func(sp *search.SearchPortfolio) (model.SearchResult, error) {
		key := model.EmailKey
		fmt.Printf("🔍 Search for %s results\n", key)
		values, err := email.SearchByProbe(searchTerm, maxSize)
		if err != nil {
			return model.SearchResult{}, err
		}
		records := []model.SearchRecord{}
		for _, value := range values {
			records = append(records, &value)
		}
		return model.SearchResult{Key: key, Records: records}, nil
	})

	if err := ptf.Run(); err != nil {
		return err
	}
	ptf.Display()

	return nil
}
