package search

import (
	"fmt"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

// SearchFormatter formats different types of results.
type SearchFormatter struct{}

// FormatGo formats Go package results.
func (f *SearchFormatter) FormatGo(result any) string {
	res := result.(model.GoPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [golang] %s (%s) ->\n\t%s", res.Name, res.Path, desc)
}

// FormatNPM formats NPM package results.
func (f *SearchFormatter) FormatNPM(result any) string {
	res := result.(model.NPMPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [npm] %s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, desc)
}

// FormatPyPI formats PyPI package results.
func (f *SearchFormatter) FormatPyPI(result any) string {
	res := result.(model.PyPIPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [pypi] %s by %s ->\n\t%s", res.Name, res.Author, desc)
}

// FormatDNS formats DNS results.
func (f *SearchFormatter) FormatDNS(result any) string {
	res := result.(model.DNSResult)
	var content string
	if len(res.IPList) > 3 {
		content = fmt.Sprintf("The first 3 IPs are %v", res.IPList[0:3])
	} else {
		content = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [dns] %s w/ %d IPs ->\n\t%v", res.FQDN, len(res.IPList), content)
}

// SearchPortfolio has result slices and task helpers.
type SearchPortfolio struct {
	Results struct {
		Golang []model.GoPackageResult
		NPM    []model.NPMPackageResult
		PyPI   []model.PyPIPackageResult
		DNS    []model.DNSResult
	}
	Errs struct {
		Golang error
		NPM    error
		PyPI   error
		DNS    error
	}
	wg *sync.WaitGroup
}

// NewSearchPortfolio creates a new portfolio instance.
func NewSearchPortfolio() *SearchPortfolio {
	return &SearchPortfolio{wg: &sync.WaitGroup{}}
}

// Size returns the number of results collected.
func (p *SearchPortfolio) Size() int {
	return (len(p.Results.NPM) +
		len(p.Results.Golang) +
		len(p.Results.PyPI) +
		len(p.Results.DNS))
}

// Run invokes a goroutine and increments wg counter.
func (p *SearchPortfolio) Run(f func(ptf *SearchPortfolio)) {
	p.wg.Add(1)
	go f(p)
}

// Done decrements wg counter.
func (p *SearchPortfolio) Done() {
	p.wg.Done()
}

// Wait blocks the main thread until all runners are complete.
func (p *SearchPortfolio) Wait() {
	p.wg.Wait()
}

// Errors returns all errors found.
func (p *SearchPortfolio) Errors() []error {
	coll := []error{}
	if p.Size() == 0 {
		coll = append(coll, ErrPorftolioEmpty)
	}
	if p.Errs.Golang != nil {
		coll = append(coll, p.Errs.Golang)
	}
	if p.Errs.NPM != nil {
		coll = append(coll, p.Errs.NPM)
	}
	if p.Errs.PyPI != nil {
		coll = append(coll, p.Errs.PyPI)
	}
	if p.Errs.DNS != nil {
		coll = append(coll, p.Errs.DNS)
	}
	return coll
}
