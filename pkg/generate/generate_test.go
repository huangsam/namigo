package generate_test

import (
	"testing"

	"github.com/huangsam/namigo/pkg/generate"
)

func TestGeneratePrompt(t *testing.T) {
	tests := []struct {
		name         string
		purpose      string
		theme        string
		demographics string
		interests    string
		size         int
		length       int
		wantErr      bool
	}{
		{
			name:         "Valid input",
			purpose:      "Create a new project",
			theme:        "Technology",
			demographics: "Developers",
			interests:    "Open Source",
			size:         5,
			length:       10,
			wantErr:      false,
		},
		{
			name:         "Empty input",
			purpose:      "",
			theme:        "",
			demographics: "",
			interests:    "",
			size:         0,
			length:       0,
			wantErr:      false,
		},
		{
			name:         "Negative size and length",
			purpose:      "Test negative values",
			theme:        "Negative",
			demographics: "Testers",
			interests:    "Testing",
			size:         -1,
			length:       -1,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generate.Prompt(tt.purpose, tt.theme, tt.demographics, tt.interests, tt.size, tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratePrompt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Errorf("GeneratePrompt() got empty string, want non-empty string")
			}
		})
	}
}
