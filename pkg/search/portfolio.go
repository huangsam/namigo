package search

import (
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

// Portfolio has entity helpers and task helpers.
type Portfolio struct {
	Res PortfolioResults
	Err PortfolioErrors
	Fmt PortfolioFormatters
	wg  *sync.WaitGroup
}

type PortfolioResults struct {
	Golang []model.GoPackage
	NPM    []model.NPMPackage
	PyPI   []model.PyPIPackage
	DNS    []model.DNSRecord
}

type PortfolioErrors struct {
	Golang error
	NPM    error
	PyPI   error
	DNS    error
}

type PortfolioFormatters struct {
	Golang GoFormatter
	NPM    NPMFormatter
	PyPI   PyPIFormatter
	DNS    DNSFormatter
}

// NewPortfolio creates a new portfolio instance.
func NewPortfolio() *Portfolio {
	return &Portfolio{wg: &sync.WaitGroup{}}
}

// Size returns the number of results collected.
func (p *Portfolio) Size() int {
	return (len(p.Res.NPM) +
		len(p.Res.Golang) +
		len(p.Res.PyPI) +
		len(p.Res.DNS))
}

// Run invokes a goroutine and increments internal WaitGroup counter.
func (p *Portfolio) Run(f func(ptf *Portfolio)) {
	p.wg.Add(1)
	go f(p)
}

// Done decrements internal WaitGroup counter.
func (p *Portfolio) Done() {
	p.wg.Done()
}

// Wait blocks the main thread until all goroutines complete.
func (p *Portfolio) Wait() {
	p.wg.Wait()
}

// Errors returns all errors found.
func (p *Portfolio) Errors() []error {
	errs := []error{}
	if p.Err.Golang != nil {
		errs = append(errs, p.Err.Golang)
	}
	if p.Err.NPM != nil {
		errs = append(errs, p.Err.NPM)
	}
	if p.Err.PyPI != nil {
		errs = append(errs, p.Err.PyPI)
	}
	if p.Err.DNS != nil {
		errs = append(errs, p.Err.DNS)
	}
	return errs
}
