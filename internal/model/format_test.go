package model_test

import (
	"testing"

	"github.com/huangsam/namigo/internal/model"
)

func TestGetOutputFormat(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected model.FormatOption
	}{
		{"Text format", "text", model.TextOption},
		{"JSON format", "json", model.JSONOption},
		{"Default format", "unknown", model.TextOption},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := model.GetOutputFormat(tt.format); got != tt.expected {
				t.Errorf("GetOutputMode(%v) = %v, want %v", tt.format, got, tt.expected)
			}
		})
	}
}
