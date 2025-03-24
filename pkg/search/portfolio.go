package search

import (
	"reflect"
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
	Email  []model.EmailRecord
}

type PortfolioErrors struct {
	Golang error
	NPM    error
	PyPI   error
	DNS    error
	Email  error
}

type PortfolioFormatters struct {
	Golang GoFormatter
	NPM    NPMFormatter
	PyPI   PyPIFormatter
	DNS    DNSFormatter
	Email  EmailFormatter
}

// NewPortfolio creates a new portfolio instance.
func NewPortfolio() *Portfolio {
	return &Portfolio{wg: &sync.WaitGroup{}}
}

// Size returns the number of results collected.
func (p *Portfolio) Size() int {
	totalSize := 0
	v := reflect.ValueOf(p.Res)
	for i := range v.NumField() {
		field := v.Field(i)
		if field.Kind() == reflect.Slice {
			totalSize += field.Len()
		}
	}
	return totalSize
}

// Errors returns all errors found.
func (p *Portfolio) Errors() []error {
	errs := []error{}
	v := reflect.ValueOf(p.Err)
	for i := range v.NumField() {
		field := v.Field(i)
		if field.Interface() != nil {
			errs = append(errs, field.Interface().(error))
		}
	}
	return errs
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
