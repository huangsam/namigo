package search_test

import (
	"net"
	"testing"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
)

func TestGoFormatter_FormatResult(t *testing.T) {
	formatter := &search.GoFormatter{}
	result := model.GoPackage{
		Name:        "example-go-package",
		Path:        "github.com/example/go-package",
		Description: "This is an example Go package used for testing purposes.",
	}
	expected := "📦 [Golang] example-go-package (github.com/example/go-package) ->\n\tThis is an example Go package used for testing purposes."
	if got := formatter.Format(result); got != expected {
		t.Errorf("GoFormatter.FormatResult() = %v, want %v", got, expected)
	}
}

func TestNPMFormatter_FormatResult(t *testing.T) {
	formatter := &search.NPMFormatter{}
	result := model.NPMPackage{
		Name:         "example-npm-package",
		IsExactMatch: true,
		Description:  "This is an example NPM package used for testing purposes.",
	}
	expected := "📦 [NPM] example-npm-package [exact=true] ->\n\tThis is an example NPM package used for testing purposes."
	if got := formatter.Format(result); got != expected {
		t.Errorf("NPMFormatter.FormatResult() = %v, want %v", got, expected)
	}
}

func TestPyPIFormatter_FormatResult(t *testing.T) {
	formatter := &search.PyPIFormatter{}
	result := model.PyPIPackage{
		Name:        "example-pypi-package",
		Author:      "example-author",
		Description: "This is an example PyPI package used for testing purposes.",
	}
	expected := "📦 [PyPI] example-pypi-package by example-author ->\n\tThis is an example PyPI package used for testing purposes."
	if got := formatter.Format(result); got != expected {
		t.Errorf("PyPIFormatter.FormatResult() = %v, want %v", got, expected)
	}
}

func TestDNSFormatter_FormatResult(t *testing.T) {
	formatter := &search.DNSFormatter{}
	result := model.DNSRecord{
		FQDN: "example.com",
		IPList: []net.IP{
			net.ParseIP("192.168.1.1"),
			net.ParseIP("192.168.1.2"),
			net.ParseIP("192.168.1.3"),
			net.ParseIP("192.168.1.4"),
		},
	}
	expected := "🌎 [DNS] example.com w/ 4 IPs ->\n\tThe first 3 IPs are [192.168.1.1 192.168.1.2 192.168.1.3]"
	if got := formatter.Format(result); got != expected {
		t.Errorf("DNSFormatter.FormatResult() = %v, want %v", got, expected)
	}
}
func TestEmailFormatter_FormatResult(t *testing.T) {
	formatter := &search.EmailFormatter{}
	result := model.EmailRecord{
		Addr:           "example@example.com",
		HasValidSyntax: true,
		HasValidDomain: false,
	}
	expected := "📨 [Email] example@example.com ->\n\tvalid-syntax=true, valid-domain=false"
	if got := formatter.Format(result); got != expected {
		t.Errorf("EmailFormatter.FormatResult() = %v, want %v", got, expected)
	}
}
