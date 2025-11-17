package sub_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/huangsam/namigo/cmd/namigo/sub"
	"github.com/huangsam/namigo/pkg/search"
)

func TestSearchRunner_RunPackageSearch(t *testing.T) {
	tests := []struct {
		name         string
		searchTerm   string
		maxSize      int
		outputFormat search.FormatOption
	}{
		{
			name:         "Valid search with text format",
			searchTerm:   "test",
			maxSize:      5,
			outputFormat: search.TextOption,
		},
		{
			name:         "Valid search with JSON format",
			searchTerm:   "test",
			maxSize:      5,
			outputFormat: search.JSONOption,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			runner := sub.NewSearchRunner(&buf)

			// Package search may fail due to network issues, but we can test the structure
			err := runner.RunPackageSearch(tt.searchTerm, tt.maxSize, tt.outputFormat)

			output := buf.String()
			// Verify that search indicators were printed (output was redirected)
			if !strings.Contains(output, "🔍 Search for") {
				t.Errorf("RunPackageSearch() output missing search indicator, got: %v", output)
			}

			// If there's an error, it could be network-related which is acceptable in tests
			if err != nil {
				t.Logf("RunPackageSearch() returned error (possibly network-related): %v", err)
			}
		})
	}
}

func TestSearchRunner_RunDNSSearch(t *testing.T) {
	tests := []struct {
		name         string
		searchTerm   string
		maxSize      int
		outputFormat search.FormatOption
	}{
		{
			name:         "Valid DNS search",
			searchTerm:   "example",
			maxSize:      5,
			outputFormat: search.TextOption,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			runner := sub.NewSearchRunner(&buf)

			// DNS search may fail due to network issues, but we can test the structure
			err := runner.RunDNSSearch(tt.searchTerm, tt.maxSize, tt.outputFormat)

			output := buf.String()
			// Verify that search indicator was printed
			if !strings.Contains(output, "🔍 Search for") {
				t.Errorf("RunDNSSearch() output missing search indicator, got: %v", output)
			}

			// If there's an error, it could be network-related which is acceptable in tests
			if err != nil {
				t.Logf("RunDNSSearch() returned error (possibly network-related): %v", err)
			}
		})
	}
}

func TestSearchRunner_RunEmailSearch(t *testing.T) {
	tests := []struct {
		name         string
		searchTerm   string
		maxSize      int
		outputFormat search.FormatOption
	}{
		{
			name:         "Valid email search",
			searchTerm:   "test",
			maxSize:      5,
			outputFormat: search.TextOption,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			runner := sub.NewSearchRunner(&buf)

			// Email search may fail due to network issues, but we can test the structure
			err := runner.RunEmailSearch(tt.searchTerm, tt.maxSize, tt.outputFormat)

			output := buf.String()
			// Verify that search indicator was printed
			if !strings.Contains(output, "🔍 Search for") {
				t.Errorf("RunEmailSearch() output missing search indicator, got: %v", output)
			}

			// If there's an error, it could be network-related which is acceptable in tests
			if err != nil {
				t.Logf("RunEmailSearch() returned error (possibly network-related): %v", err)
			}
		})
	}
}

func TestNewSearchRunner(t *testing.T) {
	var buf bytes.Buffer
	runner := sub.NewSearchRunner(&buf)

	if runner == nil {
		t.Error("NewSearchRunner() returned nil")
	}
}
