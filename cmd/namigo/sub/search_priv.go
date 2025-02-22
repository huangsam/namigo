package sub

import (
	"fmt"
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
	errs struct {
		golang error
		npm    error
		pypi   error
		dns    error
	}
	wg *sync.WaitGroup
}

// newSearchPortfolio creates a new portfolio instance.
func newSearchPortfolio() *searchPortfolio {
	return &searchPortfolio{wg: &sync.WaitGroup{}}
}

// size returns the number of results collected.
func (p *searchPortfolio) size() int {
	return (len(p.results.npm) +
		len(p.results.golang) +
		len(p.results.pypi) +
		len(p.results.dns))
}

// run invokes a goroutine and increments wg counter.
func (p *searchPortfolio) run(f func(ptf *searchPortfolio)) {
	p.wg.Add(1)
	go f(p)
}

// done decrements wg counter.
func (p *searchPortfolio) done() {
	p.wg.Done()
}

// wait blocks the main thread until all runners are complete.
func (p *searchPortfolio) wait() {
	p.wg.Wait()
}

// errors returns all errors found.
func (p *searchPortfolio) errors() []error {
	coll := []error{}
	if p.size() == 0 {
		coll = append(coll, ErrPorftolioEmpty)
	}
	if p.errs.golang != nil {
		coll = append(coll, p.errs.golang)
	}
	if p.errs.npm != nil {
		coll = append(coll, p.errs.npm)
	}
	if p.errs.pypi != nil {
		coll = append(coll, p.errs.pypi)
	}
	if p.errs.dns != nil {
		coll = append(coll, p.errs.dns)
	}
	return coll
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
