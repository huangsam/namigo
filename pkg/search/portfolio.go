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
	resultMap    map[model.SearchRecordKey][]model.SearchRecord
	errorMap     map[model.SearchRecordKey]error
	lineMap      map[model.SearchRecordKey]SearchRecordLine
	formatOption FormatOption
	resultFuncs  []SearchResultFunc
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
		formatOption: option,
		resultFuncs:  []SearchResultFunc{},
	}
}

// Register invokes a goroutine and increments internal WaitGroup counter.
func (p *SearchPortfolio) Register(f SearchResultFunc) {
	p.resultFuncs = append(p.resultFuncs, f)
}

// Run blocks the main thread until all goroutines complete.
func (p *SearchPortfolio) Run() error {
	var wg sync.WaitGroup
	var emu sync.Mutex
	var rmu sync.Mutex
	for _, resultFunc := range p.resultFuncs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if result, err := resultFunc(p); err != nil {
				emu.Lock() // Critical section
				p.errorMap[result.Key] = err
				emu.Unlock()
			} else {
				rmu.Lock() // Critical section
				p.resultMap[result.Key] = result.Records
				rmu.Unlock()
			}
		}()
	}
	wg.Wait()

	if len(p.errorMap) > 0 {
		return ErrPorftolioFailure
	}
	if len(p.resultMap) == 0 {
		return ErrPorftolioEmpty
	}
	return nil
}

// Display prints results across all results
func (p *SearchPortfolio) Display() {
	fmt.Printf("🍺 Prepare %s results\n\n", p.formatOption)
	time.Sleep(resultDelay)
	for key := range p.resultMap {
		results := p.resultMap[key]
		line := p.lineMap[key] // assume that it exists
		option := p.formatOption
		display(key, results, line, option)
	}
}

// jsonWrapper is a helper struct for JSON formatting.
type jsonWrapper struct {
	Label   string               `json:"label"`
	Results []model.SearchRecord `json:"results"`
}

// display prints results based on the specified parameters.
func display(
	key model.SearchRecordKey,
	results []model.SearchRecord,
	line SearchRecordLine,
	option FormatOption,
) {
	label := key.String()
	if len(results) == 0 {
		return
	}
	switch option {
	case JSONOption:
		data, err := json.MarshalIndent(&jsonWrapper{Label: label, Results: results}, "", "  ")
		if err != nil {
			fmt.Printf("Cannot print %s for %s: %v\n", option, label, err)
			return
		}
		fmt.Printf("%s\n", data)
	case TextOption:
		for _, r := range results {
			fmt.Println(line.Format(r))
		}
	}
}
