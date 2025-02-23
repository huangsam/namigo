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

// Portfolio has entity helpers and task helpers.
type Portfolio struct {
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

// NewPortfolio creates a new portfolio instance.
func NewPortfolio() *Portfolio {
	return &Portfolio{wg: &sync.WaitGroup{}}
}

// Size returns the number of results collected.
func (p *Portfolio) Size() int {
	return (len(p.Results.NPM) +
		len(p.Results.Golang) +
		len(p.Results.PyPI) +
		len(p.Results.DNS))
}

// Run invokes a goroutine and increments wg counter.
func (p *Portfolio) Run(f func(ptf *Portfolio)) {
	p.wg.Add(1)
	go f(p)
}

// Done decrements wg counter.
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
	if p.Size() == 0 {
		errs = append(errs, ErrPorftolioEmpty)
	}
	if p.Errs.Golang != nil {
		errs = append(errs, p.Errs.Golang)
	}
	if p.Errs.NPM != nil {
		errs = append(errs, p.Errs.NPM)
	}
	if p.Errs.PyPI != nil {
		errs = append(errs, p.Errs.PyPI)
	}
	if p.Errs.DNS != nil {
		errs = append(errs, p.Errs.DNS)
	}
	return errs
}
