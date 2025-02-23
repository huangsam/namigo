package search

import (
	"fmt"

	"github.com/huangsam/namigo/internal/model"
)

const (
	GolangLabel = "Golang" // Label for Golang packages
	NPMLabel    = "NPM"    // Label for NPM packages
	PyPILabel   = "PyPI"   // Label for PyPI packages
	DNSLabel    = "DNS"    // Label for DNS records

	MaxLineLength = 80 // Maximum line length
	MaxIPLength   = 3  // Maximum IP length
)

// Formatter is an interface for formatting different types of results.
type Formatter interface {
	// FormatResult formats a single result.
	FormatResult(result any) string

	// Label returns the canonical label for an entity.
	Label() string
}

// GoFormatter formats Go package.
type GoFormatter struct{}

func (f *GoFormatter) FormatResult(result any) string {
	res := result.(model.GoPackage)
	desc := res.Description
	if len(desc) > MaxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s (%s) ->\n\t%s", f.Label(), res.Name, res.Path, desc)
}

func (f *GoFormatter) Label() string {
	return GolangLabel
}

// NPMFormatter formats NPM package.
type NPMFormatter struct{}

func (f *NPMFormatter) FormatResult(result any) string {
	res := result.(model.NPMPackage)
	desc := res.Description
	if len(desc) > MaxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s [exact=%v] ->\n\t%s", f.Label(), res.Name, res.IsExactMatch, desc)
}

func (f *NPMFormatter) Label() string {
	return NPMLabel
}

// PyPIFormatter formats PyPI package.
type PyPIFormatter struct{}

func (f *PyPIFormatter) FormatResult(result any) string {
	res := result.(model.PyPIPackage)
	desc := res.Description
	if len(desc) > MaxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s by %s ->\n\t%s", f.Label(), res.Name, res.Author, desc)
}

func (f *PyPIFormatter) Label() string {
	return PyPILabel
}

// DNSFormatter formats DNS record.
type DNSFormatter struct{}

func (f *DNSFormatter) FormatResult(result any) string {
	res := result.(model.DNSRecord)
	var desc string
	if len(res.IPList) > MaxIPLength {
		desc = fmt.Sprintf("The first %d IPs are %v", MaxIPLength, res.IPList[0:3])
	} else {
		desc = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [%s] %s w/ %d IPs ->\n\t%v", f.Label(), res.FQDN, len(res.IPList), desc)
}

func (f *DNSFormatter) Label() string {
	return DNSLabel
}
