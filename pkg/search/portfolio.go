package search

import (
	"errors"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

var (
	ErrPorftolioEmpty   = errors.New("portfolio collection empty")
	ErrPorftolioFailure = errors.New("portfolio collection failure")
)

// SearchPortfolio has entity helpers and task helpers.
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
	Formats struct {
		Golang GoFormatter
		NPM    NPMFormatter
		PyPI   PyPIFormatter
		DNS    DNSFormatter
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

// Wait blocks the main thread until all goroutines complete.
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
