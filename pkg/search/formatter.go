package search

import (
	"fmt"

	"github.com/huangsam/namigo/internal/model"
)

const (
	GolangLabel = "Golang"
	NPMLabel    = "NPM"
	PyPILabel   = "PyPI"
	DNSLabel    = "DNS"
)

// Formatter is an interface for formatting different types of results.
type Formatter interface {
	Format(result any) string
	Label() string
}

// GoFormatter formats Go package results.
type GoFormatter struct{}

// Format formats a single Go package result.
func (f *GoFormatter) Format(result any) string {
	res := result.(model.GoPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s (%s) ->\n\t%s", f.Label(), res.Name, res.Path, desc)
}

// Label returns the label for GoFormatter.
func (f *GoFormatter) Label() string {
	return GolangLabel
}

// NPMFormatter formats NPM package results.
type NPMFormatter struct{}

// Format formats a single NPM package result.
func (f *NPMFormatter) Format(result any) string {
	res := result.(model.NPMPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s [exact=%v] ->\n\t%s", f.Label(), res.Name, res.IsExactMatch, desc)
}

// Label returns the label for NPMFormatter.
func (f *NPMFormatter) Label() string {
	return NPMLabel
}

// PyPIFormatter formats PyPI package results.
type PyPIFormatter struct{}

// Format formats a single PyPI package result.
func (f *PyPIFormatter) Format(result any) string {
	res := result.(model.PyPIPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s by %s ->\n\t%s", f.Label(), res.Name, res.Author, desc)
}

// Label returns the label for PyPIFormatter.
func (f *PyPIFormatter) Label() string {
	return PyPILabel
}

// DNSFormatter formats DNS results.
type DNSFormatter struct{}

// Format formats a single DNS result.
func (f *DNSFormatter) Format(result any) string {
	res := result.(model.DNSResult)
	var content string
	if len(res.IPList) > 3 {
		content = fmt.Sprintf("The first 3 IPs are %v", res.IPList[0:3])
	} else {
		content = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [%s] %s w/ %d IPs ->\n\t%v", f.Label(), res.FQDN, len(res.IPList), content)
}

// Label returns the label for DNSFormatter.
func (f *DNSFormatter) Label() string {
	return DNSLabel
}
