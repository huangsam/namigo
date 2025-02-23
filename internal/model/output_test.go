package model_test

import (
	"testing"

	"github.com/huangsam/namigo/internal/model"
)

func TestGetOutputMode(t *testing.T) {
	tests := []struct {
		name     string
		mode     string
		expected model.OutputMode
	}{
		{"Text mode", "text", model.TextMode},
		{"JSON mode", "json", model.JSONMode},
		{"Default mode", "unknown", model.TextMode},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := model.GetOutputMode(tt.mode); got != tt.expected {
				t.Errorf("GetOutputMode(%v) = %v, want %v", tt.mode, got, tt.expected)
			}
		})
	}
}
