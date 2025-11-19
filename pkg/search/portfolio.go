package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/huangsam/namigo/v2/internal/model"
)

// resultDelay is the delay before displaying results.
const resultDelay = 500 * time.Millisecond

// searchTypeCount is the number of search types.
const searchTypeCount = 5

var (
	// ErrPorftolioEmpty is returned when the portfolio is empty.
	ErrPorftolioEmpty = errors.New("portfolio collection empty")

	// ErrPorftolioFailure is returned when the portfolio fails.
	ErrPorftolioFailure = errors.New("portfolio collection failure")
)

// ResultFunc is a function that returns search results.
type ResultFunc func() (model.SearchResult, error)

// Portfolio has entity helpers and task helpers.
type Portfolio struct {
	resultMap map[model.SearchKey][]model.SearchRecord
	lineMap   map[model.SearchKey]LineFunc
	option    FormatOption
	funcs     []ResultFunc
	errors    []error
	output    io.Writer
}

// NewSearchPortfolio creates a new portfolio instance.
func NewSearchPortfolio(format FormatOption, output io.Writer) *Portfolio {
	return &Portfolio{
		resultMap: make(map[model.SearchKey][]model.SearchRecord, searchTypeCount), // Pre-allocate for 5 search types
		lineMap: map[model.SearchKey]LineFunc{
			model.GoKey:    GoLine,
			model.NPMKey:   NPMLine,
			model.PyPIKey:  PyPILine,
			model.DNSKey:   DNSLine,
			model.EmailKey: EmailLine,
		},
		option: format,
		funcs:  make([]ResultFunc, 0, searchTypeCount), // Pre-allocate capacity
		errors: make([]error, 0, searchTypeCount),      // Pre-allocate capacity
		output: output,
	}
}

// Register invokes a goroutine and increments internal WaitGroup counter.
func (p *Portfolio) Register(f ResultFunc) {
	p.funcs = append(p.funcs, f)
}

// Run blocks the main thread until all goroutines complete.
func (p *Portfolio) Run() error {
	var wg sync.WaitGroup
	var emu sync.Mutex
	var rmu sync.Mutex
	for _, fn := range p.funcs {
		wg.Go(func() {
			if result, err := fn(); err != nil {
				emu.Lock() // Critical section
				p.errors = append(p.errors, err)
				emu.Unlock()
			} else {
				rmu.Lock() // Critical section
				p.resultMap[result.Key] = result.Records
				rmu.Unlock()
			}
		})
	}
	wg.Wait()

	if len(p.errors) > 0 {
		return ErrPorftolioFailure
	}
	if len(p.resultMap) == 0 {
		return ErrPorftolioEmpty
	}
	return nil
}

// Display prints results across all results.
func (p *Portfolio) Display() {
	_, _ = fmt.Fprintf(p.output, "🍺 Prepare %s results\n\n", p.option)
	time.Sleep(resultDelay)
	for key, records := range p.resultMap {
		label := key.String()
		line := p.lineMap[key] // Assume that this exists
		option := p.option
		if len(records) == 0 {
			return
		}
		switch option {
		case TextOption:
			for _, record := range records {
				_, _ = fmt.Fprintln(p.output, line(label, record))
			}
		case JSONOption:
			if b, err := json.MarshalIndent(&model.SearchRender{Label: label, Result: records}, "", "  "); err != nil {
				_, _ = fmt.Fprintf(p.output, "Cannot display %s: %s\n", label, err.Error())
			} else {
				_, _ = fmt.Fprintf(p.output, "%s\n", b)
			}
		}
	}
}
