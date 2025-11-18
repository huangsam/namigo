package dns

import (
	"net"
	"testing"

	"github.com/huangsam/namigo/internal/core"
)

func TestSearchByProbeWithLookup(t *testing.T) {
	lookup := func(domain string) ([]net.IP, error) {
		if domain == "test.com" {
			return []net.IP{net.ParseIP("1.2.3.4")}, nil
		}
		return nil, nil
	}

	result, err := SearchByProbeWithLookup("test", 10, lookup) // Increase size to get all
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	found := false
	for _, r := range result {
		if r.FQDN == "test.com" && len(r.IPList) > 0 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find test.com with IPs")
	}
}

func TestIsValidDomainName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid simple", "example", true},
		{"valid with numbers", "test123", true},
		{"valid with hyphen", "test-name", true},
		{"valid mixed case", "TestName", true},
		{"empty string", "", false},
		{"too long", "thisdomainnameiswaytoolongandshouldfailvalidationaccordingtoourrules", false},
		{"starts with hyphen", "-test", false},
		{"ends with hyphen", "test-", false},
		{"contains space", "test name", false},
		{"contains dot", "test.name", false},
		{"contains underscore", "test_name", false},
		{"contains special char", "test@name", false},
		{"only numbers", "123", true},
		{"single char", "a", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := core.IsValidDomainName(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidDomainName(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSearchByProbe_InvalidDomain(t *testing.T) {
	_, err := SearchByProbe("invalid@domain", 5)
	if err == nil {
		t.Error("expected error for invalid domain name")
	}
	if err.Error() != "invalid domain name: invalid@domain" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestSearchByProbe_ValidDomain(t *testing.T) {
	// This will fail due to network calls, but we just want to test validation passes
	_, err := SearchByProbe("test", 5)
	// We expect this to fail with network errors, not validation errors
	if err != nil && err.Error() == "invalid domain name: test" {
		t.Error("validation should have passed for valid domain")
	}
}
