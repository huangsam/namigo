// Package generate deals with prompt generation.
package generate

import (
	"strings"
	"text/template"
)

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
