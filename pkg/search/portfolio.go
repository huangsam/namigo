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
	resultMap    map[model.SearchKey][]model.SearchRecord
	errorMap     map[model.SearchKey]error
	lineMap      map[model.SearchKey]SearchRecordLine
	formatOption FormatOption
	resultFuncs  []SearchResultFunc
}

// NewSearchPortfolio creates a new portfolio instance.
func NewSearchPortfolio(option FormatOption) *SearchPortfolio {
	return &SearchPortfolio{
		resultMap: map[model.SearchKey][]model.SearchRecord{},
		lineMap: map[model.SearchKey]SearchRecordLine{
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

// display prints results based on the specified parameters.
func display(
	key model.SearchKey,
	records []model.SearchRecord,
	line SearchRecordLine,
	option FormatOption,
) {
	if len(records) == 0 {
		return
	}
	label := key.String()
	switch option {
	case JSONOption:
		data, err := json.MarshalIndent(&model.SearchJSON{Label: label, Result: records}, "", "  ")
		if err != nil {
			fmt.Printf("Cannot print %s for %s: %v\n", option, key, err)
			return
		}
		fmt.Printf("%s\n", data)
	case TextOption:
		for _, record := range records {
			fmt.Println(line.Format(label, record))
		}
	}
}
