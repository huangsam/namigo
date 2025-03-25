package search_test

import (
	"net"
	"testing"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
)

func TestGoLine_Format(t *testing.T) {
	line := &search.GoLine{}
	result := model.GoPackage{
		Name:        "example-go-package",
		Path:        "github.com/example/go-package",
		Description: "This is an example Go package used for testing purposes.",
	}
	expected := "📦 [Golang] example-go-package (github.com/example/go-package) ->\n\tThis is an example Go package used for testing purposes."
	if got := line.Format(&result); got != expected {
		t.Errorf("GoFormatter.Format() = %v, want %v", got, expected)
	}
}

func TestNPMLine_Format(t *testing.T) {
	line := &search.NPMLine{}
	result := model.NPMPackage{
		Name:        "example-npm-package",
		Description: "This is an example NPM package used for testing purposes.",
	}
	expected := "📦 [NPM] example-npm-package ->\n\tThis is an example NPM package used for testing purposes."
	if got := line.Format(&result); got != expected {
		t.Errorf("NPMFormatter.Format() = %v, want %v", got, expected)
	}
}

func TestPyPILine_Format(t *testing.T) {
	line := &search.PyPILine{}
	result := model.PyPIPackage{
		Name:        "example-pypi-package",
		Author:      "example-author",
		Description: "This is an example PyPI package used for testing purposes.",
	}
	expected := "📦 [PyPI] example-pypi-package by example-author ->\n\tThis is an example PyPI package used for testing purposes."
	if got := line.Format(&result); got != expected {
		t.Errorf("PyPIFormatter.Format() = %v, want %v", got, expected)
	}
}

func TestDNSLine_Format(t *testing.T) {
	line := &search.DNSLine{}
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
	if got := line.Format(&result); got != expected {
		t.Errorf("DNSFormatter.Format() = %v, want %v", got, expected)
	}
}
func TestEmailLine_Format(t *testing.T) {
	line := &search.EmailLine{}
	result := model.EmailRecord{
		Addr:           "example@example.com",
		HasValidSyntax: true,
		HasValidDomain: false,
	}
	expected := "📨 [Email] example@example.com ->\n\tvalid-syntax=true, valid-domain=false"
	if got := line.Format(&result); got != expected {
		t.Errorf("EmailFormatter.Format() = %v, want %v", got, expected)
	}
}
