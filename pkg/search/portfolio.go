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

type SearchResultFunc func(*SearchPortfolio) (model.SearchResult, error)

// SearchPortfolio has entity helpers and task helpers.
type SearchPortfolio struct {
	resultMap map[model.SearchRecordKey][]model.SearchRecord
	lineMap   map[model.SearchRecordKey]SearchRecordLine
	option    FormatOption
	callers   []SearchResultFunc
	errors    []error
}

// NewSearchPortfolio creates a new portfolio instance.
func NewSearchPortfolio(option FormatOption) *SearchPortfolio {
	return &SearchPortfolio{
		resultMap: map[model.SearchRecordKey][]model.SearchRecord{},
		lineMap: map[model.SearchRecordKey]SearchRecordLine{
			model.GoKey:    &GoLine{},
			model.NPMKey:   &NPMLine{},
			model.PyPIKey:  &PyPILine{},
			model.DNSKey:   &DNSLine{},
			model.EmailKey: &EmailLine{},
		},
		option:  option,
		callers: []SearchResultFunc{},
		errors:  []error{},
	}
}

// Size returns the number of results collected.
func (p *SearchPortfolio) Size() int {
	total := 0
	for _, records := range p.resultMap {
		total += len(records)
	}
	return total
}

// Register invokes a goroutine and increments internal WaitGroup counter.
func (p *SearchPortfolio) Register(f SearchResultFunc) {
	p.callers = append(p.callers, f)
}

// Run blocks the main thread until all goroutines complete.
func (p *SearchPortfolio) Run() error {
	var wg sync.WaitGroup
	for _, caller := range p.callers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if result, err := caller(p); err != nil {
				p.errors = append(p.errors, err)
			} else {
				p.resultMap[result.Key] = result.Records
			}
		}()
	}
	wg.Wait()

	if len(p.errors) > 0 {
		return ErrPorftolioFailure
	}
	if p.Size() == 0 {
		return ErrPorftolioEmpty
	}
	return nil
}

// displayResults displays results based on the specified parameters.
func displayResults[T model.SearchRecord](
	key model.SearchRecordKey,
	results []T,
	formatter SearchRecordLine,
	format FormatOption,
) {
	label := key.String()
	if len(results) == 0 {
		return
	}
	switch format {
	case JSONOption:
		type wrapper struct {
			Label   string `json:"label"`
			Results []T    `json:"results"`
		}
		data, err := json.MarshalIndent(&wrapper{Label: label, Results: results}, "", "  ")
		if err != nil {
			fmt.Printf("Cannot print %s for %s: %v\n", format, label, err)
			return
		}
		fmt.Printf("%s\n", data)
	case TextOption:
		for _, r := range results {
			fmt.Println(formatter.Format(r))
		}
	}
}

// Display displays results across all results
func (p *SearchPortfolio) Display() {
	fmt.Printf("🍺 Prepare %s results\n\n", p.option)
	time.Sleep(resultDelay)
	for key := range p.resultMap {
		results := p.resultMap[key]
		line := p.lineMap[key] // assume that it exists
		option := p.option
		displayResults(key, results, line, option)
	}
}
