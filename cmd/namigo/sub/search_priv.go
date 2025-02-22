package sub

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/util"
)

// searchFormatter formats different types of results.
type searchFormatter struct{}

// formatGo formats Go package results.
func (f *searchFormatter) formatGo(result any) string {
	res := result.(model.GoPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [golang] %s (%s) ->\n\t%s", res.Name, res.Path, desc)
}

// formatNPM formats NPM package results.
func (f *searchFormatter) formatNPM(result any) string {
	res := result.(model.NPMPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [npm] %s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, desc)
}

// formatPyPI formats PyPI package results.
func (f *searchFormatter) formatPyPI(result any) string {
	res := result.(model.PyPIPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [pypi] %s by %s ->\n\t%s", res.Name, res.Author, desc)
}

// formatDNS formats DNS results.
func (f *searchFormatter) formatDNS(result any) string {
	res := result.(model.DNSResult)
	var content string
	if len(res.IPList) > 3 {
		content = fmt.Sprintf("The first 3 IPs are %v", res.IPList[0:3])
	} else {
		content = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [dns] %s w/ %d IPs ->\n\t%v", res.FQDN, len(res.IPList), content)
}

// searchPortfolio has result slices and task helpers.
type searchPortfolio struct {
	results struct {
		golang []model.GoPackageResult
		npm    []model.NPMPackageResult
		pypi   []model.PyPIPackageResult
		dns    []model.DNSResult
	}
	wg *sync.WaitGroup
	c  int
}

// newSearchPortfolio creates a new portfolio instance.
func newSearchPortfolio() *searchPortfolio {
	return &searchPortfolio{wg: &sync.WaitGroup{}}
}

// isEmpty checks if the portfolio has zero results.
func (p *searchPortfolio) isEmpty() bool {
	return (len(p.results.npm) +
		len(p.results.golang) +
		len(p.results.pypi) +
		len(p.results.dns)) == 0
}

// run invokes a function as a goroutine and passes a WaitGroup into it.
func (p *searchPortfolio) run(f func(wg *sync.WaitGroup)) {
	p.wg.Add(1)
	p.c++
	go f(p.wg)
}

// wait blocks the main thread until all runners are complete.
func (p *searchPortfolio) wait() {
	p.wg.Wait()
}

// count returns the goroutine count.
func (p *searchPortfolio) count() int {
	return p.c
}

// getOutputMode returns an OutputMode instance.
func getOutputMode(mode string) util.OutputMode {
	switch mode {
	case "text":
		return util.TextMode
	case "json":
		return util.JSONMode
	default:
		return util.TextMode
	}
}

// errorMap is a custom error mapping.
type errorMap map[string]error

// newErrorMap creates a new errorMap.
func newErrorMap() errorMap {
	return map[string]error{
		golangLabel: nil,
		npmLabel:    nil,
		pypiLabel:   nil,
		dnsLabel:    nil,
	}
}

// aggregate creates a combined error.
func (em *errorMap) aggregate() error {
	errs := []string{}
	for _, err := range *em {
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, " && "))
	}
	return nil
}
