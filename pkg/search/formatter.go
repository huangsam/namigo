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
	// Result formats a single result.
	Result(result any) string

	// Label returns the canonical label for an entity.
	Label() string
}

// GoFormatter formats Go package results.
type GoFormatter struct{}

func (f *GoFormatter) Result(result any) string {
	res := result.(model.GoPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s (%s) ->\n\t%s", f.Label(), res.Name, res.Path, desc)
}

func (f *GoFormatter) Label() string {
	return GolangLabel
}

// NPMFormatter formats NPM package results.
type NPMFormatter struct{}

func (f *NPMFormatter) Result(result any) string {
	res := result.(model.NPMPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s [exact=%v] ->\n\t%s", f.Label(), res.Name, res.IsExactMatch, desc)
}

func (f *NPMFormatter) Label() string {
	return NPMLabel
}

// PyPIFormatter formats PyPI package results.
type PyPIFormatter struct{}

func (f *PyPIFormatter) Result(result any) string {
	res := result.(model.PyPIPackageResult)
	desc := res.Description
	if len(desc) > 80 || len(desc) == 0 {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s by %s ->\n\t%s", f.Label(), res.Name, res.Author, desc)
}

func (f *PyPIFormatter) Label() string {
	return PyPILabel
}

// DNSFormatter formats DNS results.
type DNSFormatter struct{}

func (f *DNSFormatter) Result(result any) string {
	res := result.(model.DNSResult)
	var content string
	if len(res.IPList) > 3 {
		content = fmt.Sprintf("The first 3 IPs are %v", res.IPList[0:3])
	} else {
		content = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [%s] %s w/ %d IPs ->\n\t%v", f.Label(), res.FQDN, len(res.IPList), content)
}

func (f *DNSFormatter) Label() string {
	return DNSLabel
}
