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
	EmailLabel  = "Email"  // Label for email records

	MaxLineLength = 80 // Maximum line length
	MaxIPLength   = 3  // Maximum IP length
)

// ResultFormatter formats results as strings.
type ResultFormatter interface {
	// Format formats a single result.
	Format(result any) string

	// Label returns the canonical label for an entity.
	Label() string
}

// GoFormatter formats Go package.
type GoFormatter struct{}

func (f *GoFormatter) Format(result any) string {
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

func (f *NPMFormatter) Format(result any) string {
	res := result.(model.NPMPackage)
	desc := res.Description
	if len(desc) > MaxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s ->\n\t%s", f.Label(), res.Name, desc)
}

func (f *NPMFormatter) Label() string {
	return NPMLabel
}

// PyPIFormatter formats PyPI package.
type PyPIFormatter struct{}

func (f *PyPIFormatter) Format(result any) string {
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

func (f *DNSFormatter) Format(result any) string {
	res := result.(model.DNSRecord)
	var desc string
	if len(res.IPList) > MaxIPLength {
		desc = fmt.Sprintf("The first %d IPs are %v", MaxIPLength, res.IPList[:3])
	} else {
		desc = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [%s] %s w/ %d IPs ->\n\t%v", f.Label(), res.FQDN, len(res.IPList), desc)
}

func (f *DNSFormatter) Label() string {
	return DNSLabel
}

// EmailFormatter formats Email record.
type EmailFormatter struct{}

func (f *EmailFormatter) Format(result any) string {
	res := result.(model.EmailRecord)
	return fmt.Sprintf("📨 [%s] %s ->\n\tvalid-syntax=%v, valid-domain=%v", f.Label(), res.Addr, res.HasValidSyntax, res.HasValidDomain)
}

func (f *EmailFormatter) Label() string {
	return EmailLabel
}
