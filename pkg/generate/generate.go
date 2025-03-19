// Package generate deals with prompt generation.
package generate

import (
	_ "embed"
	"strings"
	"text/template"
)

// promptName is a name for the template below.
const promptName = "name_brainstorm_with_json"

// promptData is the input for the prompt template.
type promptData struct {
	Purpose      string // Project purpose
	Theme        string // Project theme
	Demographics string // Target demographics
	Interests    string // Target interests
	MaxSize      int    // Maximum number of names to generate
	MaxLength    int    // Maximum length of each name
}

//go:embed project.template
var promptTemplate string

// GeneratePrompt generates the prompt for AI chatbots.
func GeneratePrompt(purpose, theme, demographics, interests string, size, length int) (string, error) {
	data := promptData{
		Purpose:      purpose,
		Theme:        theme,
		Demographics: demographics,
		Interests:    interests,
		MaxSize:      size,
		MaxLength:    length,
	}
	tmpl, err := template.New(promptName).Parse(promptTemplate)
	if err != nil {
		return "", err
	}
	builder := strings.Builder{}
	err = tmpl.Execute(&builder, data)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(builder.String()), err
}
