package search

import (
	"fmt"

	"github.com/huangsam/namigo/internal/model"
)

const (
	maxLineLength = 80 // Maximum line length
	maxIPCount    = 3  // Maximum IP count
)

// LineFunc formats search records as strings.
type LineFunc func(label string, result model.SearchRecord) string

// GoLine formats Go package.
func GoLine(label string, result model.SearchRecord) string {
	res := result.(*model.GoPackage)
	desc := res.Description
	if len(desc) > maxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s (%s) ->\n\t%s", label, res.Name, res.Path, desc)
}

// NPMLine formats NPM package.
func NPMLine(label string, result model.SearchRecord) string {
	res := result.(*model.NPMPackage)
	desc := res.Description
	if len(desc) > maxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s ->\n\t%s", label, res.Name, desc)
}

// PyPILine formats PyPI package.
func PyPILine(label string, result model.SearchRecord) string {
	res := result.(*model.PyPIPackage)
	desc := res.Description
	if len(desc) > maxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s by %s ->\n\t%s", label, res.Name, res.Author, desc)
}

// DNSLine formats DNS record.
func DNSLine(label string, result model.SearchRecord) string {
	res := result.(*model.DNSRecord)
	var desc string
	if len(res.IPList) > maxIPCount {
		desc = fmt.Sprintf("The first %d IPs are %v", maxIPCount, res.IPList[:3])
	} else {
		desc = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [%s] %s w/ %d IPs ->\n\t%v", label, res.FQDN, len(res.IPList), desc)
}

// EmailLine formats Email record.
func EmailLine(label string, result model.SearchRecord) string {
	res := result.(*model.EmailRecord)
	return fmt.Sprintf("📨 [%s] %s ->\n\tvalid-syntax=%v, valid-domain=%v", label, res.Addr, res.HasValidSyntax, res.HasValidDomain)
}
