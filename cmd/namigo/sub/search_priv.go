package sub

import (
	"fmt"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

// searchFormatter formats different types of results.
type searchFormatter struct{}

// formatGo formats Go package results.
func (f *searchFormatter) formatGo(result any) string {
	res := result.(model.GoPackageResult)
	if len(res.Description) > 80 || len(res.Description) == 0 {
		return fmt.Sprintf("\t[golang] %s %s ->\n\t\t%.80s...", res.Name, res.Path, res.Description)
	}
	return fmt.Sprintf("\t[golang] %s %s ->\n\t\t%s", res.Name, res.Path, res.Description)
}

// formatNPM formats NPM package results.
func (f *searchFormatter) formatNPM(result any) string {
	res := result.(model.NPMPackageResult)
	if len(res.Description) > 80 || len(res.Description) == 0 {
		return fmt.Sprintf("\t[npm] %s [exact=%v] ->\n\t\t%.80s...", res.Name, res.IsExactMatch, res.Description)
	}
	return fmt.Sprintf("\t[npm] %s [exact=%v] ->\n\t\t%s", res.Name, res.IsExactMatch, res.Description)
}

// formatPyPI formats PyPI package results.
func (f *searchFormatter) formatPyPI(result any) string {
	res := result.(model.PyPIPackageResult)
	if len(res.Description) > 80 || len(res.Description) == 0 {
		return fmt.Sprintf("\t[pypi] %s by %s ->\n\t\t%.80s...", res.Name, res.Author, res.Description)
	}
	return fmt.Sprintf("\t[pypi] %s by %s ->\n\t\t%s", res.Name, res.Author, res.Description)
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
	return fmt.Sprintf("\t[dns] %s w/ %d IPs ->\n\t\t%v", res.FQDN, len(res.IPList), content)
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
