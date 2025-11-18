package core

import (
	"errors"
	"testing"
)

func TestDismiss(t *testing.T) {
	called := false
	f := func() error {
		called = true
		return errors.New("test error")
	}
	dismiss(f)
	if !called {
		t.Error("function was not called")
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
			result := IsValidDomainName(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidDomainName(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
