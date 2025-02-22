package search

import (
	"fmt"

	"github.com/huangsam/namigo/internal/model"
)

// SearchFormatter formats different types of results.
type SearchFormatter struct{}

// FormatGo formats a single Go package result.
func (f *SearchFormatter) FormatGo(result any) string {
	res := result.(model.GoPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [golang] %s (%s) ->\n\t%s", res.Name, res.Path, desc)
}

// FormatNPM formats a single NPM package result.
func (f *SearchFormatter) FormatNPM(result any) string {
	res := result.(model.NPMPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [npm] %s [exact=%v] ->\n\t%s", res.Name, res.IsExactMatch, desc)
}

// FormatPyPI formats a single PyPI package result.
func (f *SearchFormatter) FormatPyPI(result any) string {
	res := result.(model.PyPIPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [pypi] %s by %s ->\n\t%s", res.Name, res.Author, desc)
}

// FormatDNS formats a single DNS result.
func (f *SearchFormatter) FormatDNS(result any) string {
	res := result.(model.DNSResult)
	var content string
	if len(res.IPList) > 3 {
		content = fmt.Sprintf("The first 3 IPs are %v", res.IPList[0:3])
	} else {
		content = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [dns] %s w/ %d IPs ->\n\t%v", res.FQDN, len(res.IPList), content)
}
