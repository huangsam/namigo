package search_test

import (
	"slices"
	"testing"

	"github.com/huangsam/namigo/v2/pkg/search"
)

func TestGetFormatOption(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected search.FormatOption
	}{
		{"Text format", "text", search.TextOption},
		{"JSON format", "json", search.JSONOption},
		{"Default format", "unknown", search.TextOption},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.GetFormatOption(tt.format); got != tt.expected {
				t.Errorf("GetOutputMode(%v) = %v, want %v", tt.format, got, tt.expected)
			}
		})
	}
}

func TestGetAllFormatOptionValues(t *testing.T) {
	expected := []string{search.TextOption.Value, search.JSONOption.Value}
	got := search.GetAllFormatOptionValues()

	if len(got) != len(expected) {
		t.Fatalf("GetAllFormatOptionValues() length = %d, want %d", len(got), len(expected))
	}

	for _, value := range expected {
		found := slices.Contains(got, value)
		if !found {
			t.Errorf("GetAllFormatOptionValues() missing value: %v", value)
		}
	}
}

func TestFormatOption_String(t *testing.T) {
	if got := search.TextOption.String(); got != "PlainText" {
		t.Errorf("TextOption.String() = %v, want %v", got, "PlainText")
	}
	if got := search.JSONOption.String(); got != "JSON" {
		t.Errorf("JSONOption.String() = %v, want %v", got, "JSON")
	}
}
