package search

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/huangsam/namigo/internal/model"
)

const resultDelay = 500 * time.Millisecond

var (
	ErrPorftolioEmpty   = errors.New("portfolio collection empty")
	ErrPorftolioFailure = errors.New("portfolio collection failure")
)

// SearchResultFunc is a function that returns search results.
type SearchResultFunc func(*SearchPortfolio) (model.SearchResult, error)

// SearchPortfolio has entity helpers and task helpers.
type SearchPortfolio struct {
	resultMap map[model.SearchKey][]model.SearchRecord
	lineMap   map[model.SearchKey]SearchLineFunc
	option    FormatOption
	funcs     []SearchResultFunc
	errors    []error
}

// NewSearchPortfolio creates a new portfolio instance.
func NewSearchPortfolio(format FormatOption) *SearchPortfolio {
	return &SearchPortfolio{
		resultMap: map[model.SearchKey][]model.SearchRecord{},
		lineMap: map[model.SearchKey]SearchLineFunc{
			model.GoKey:    GoLine,
			model.NPMKey:   NPMLine,
			model.PyPIKey:  PyPILine,
			model.DNSKey:   DNSLine,
			model.EmailKey: EmailLine,
		},
		option: format,
		funcs:  []SearchResultFunc{},
		errors: []error{},
	}
}

// Register invokes a goroutine and increments internal WaitGroup counter.
func (p *SearchPortfolio) Register(f SearchResultFunc) {
	p.funcs = append(p.funcs, f)
}

// Run blocks the main thread until all goroutines complete.
func (p *SearchPortfolio) Run() error {
	var wg sync.WaitGroup
	var emu sync.Mutex
	var rmu sync.Mutex
	for _, fn := range p.funcs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if result, err := fn(p); err != nil {
				emu.Lock() // Critical section
				p.errors = append(p.errors, err)
				emu.Unlock()
			} else {
				rmu.Lock() // Critical section
				p.resultMap[result.Key] = result.Records
				rmu.Unlock()
			}
		}()
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

// Display prints results across all results
func (p *SearchPortfolio) Display() {
	fmt.Printf("🍺 Prepare %s results\n\n", p.option)
	time.Sleep(resultDelay)
	for key, records := range p.resultMap {
		label := key.String()
		line := p.lineMap[key] // assume that it exists
		option := p.option
		if len(records) == 0 {
			return
		}
		switch option {
		case JSONOption:
			if b, err := json.MarshalIndent(&model.SearchJSON{Label: label, Result: records}, "", "  "); err != nil {
				fmt.Printf("Cannot print %s for %s: %v\n", option, label, err)
			} else {
				fmt.Printf("%s\n", b)
			}
		case TextOption:
			for _, record := range records {
				fmt.Println(line(label, record))
			}
		}
	}
}
