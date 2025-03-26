package search

import (
	"fmt"

	"github.com/huangsam/namigo/internal/model"
)

const (
	maxLineLength = 80 // Maximum line length
	maxIPLength   = 3  // Maximum IP length
)

// SearchRecordLine formats search records as strings.
type SearchRecordLine interface {
	// Format formats a single result.
	Format(result model.SearchRecord) string
}

// GoLine formats Go package.
type GoLine struct{}

func (l *GoLine) Format(result model.SearchRecord) string {
	res := result.(*model.GoPackage)
	desc := res.Description
	if len(desc) > maxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s (%s) ->\n\t%s", l.Key(), res.Name, res.Path, desc)
}

func (l *GoLine) Key() model.SearchRecordKey {
	return model.GoKey
}

// NPMLine formats NPM package.
type NPMLine struct{}

func (l *NPMLine) Format(result model.SearchRecord) string {
	res := result.(*model.NPMPackage)
	desc := res.Description
	if len(desc) > maxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s ->\n\t%s", l.Key(), res.Name, desc)
}

func (l *NPMLine) Key() model.SearchRecordKey {
	return model.NPMKey
}

// PyPILine formats PyPI package.
type PyPILine struct{}

func (l *PyPILine) Format(result model.SearchRecord) string {
	res := result.(*model.PyPIPackage)
	desc := res.Description
	if len(desc) > maxLineLength {
		desc = fmt.Sprintf("%.80s...", desc)
	}
	return fmt.Sprintf("📦 [%s] %s by %s ->\n\t%s", l.Key(), res.Name, res.Author, desc)
}

func (l *PyPILine) Key() model.SearchRecordKey {
	return model.PyPIKey
}

// DNSLine formats DNS record.
type DNSLine struct{}

func (l *DNSLine) Format(result model.SearchRecord) string {
	res := result.(*model.DNSRecord)
	var desc string
	if len(res.IPList) > maxIPLength {
		desc = fmt.Sprintf("The first %d IPs are %v", maxIPLength, res.IPList[:3])
	} else {
		desc = fmt.Sprintf("The IPs are %v", res.IPList)
	}
	return fmt.Sprintf("🌎 [%s] %s w/ %d IPs ->\n\t%v", l.Key(), res.FQDN, len(res.IPList), desc)
}

func (l *DNSLine) Key() model.SearchRecordKey {
	return model.DNSKey
}

// EmailLine formats Email record.
type EmailLine struct{}

func (l *EmailLine) Format(result model.SearchRecord) string {
	res := result.(*model.EmailRecord)
	return fmt.Sprintf("📨 [%s] %s ->\n\tvalid-syntax=%v, valid-domain=%v", l.Key(), res.Addr, res.HasValidSyntax, res.HasValidDomain)
}

func (l *EmailLine) Key() model.SearchRecordKey {
	return model.EmailKey
}
