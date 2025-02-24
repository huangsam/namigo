// Package generate deals with prompt generation.
package generate

import (
	"strings"
	"text/template"
)

// GeneratePrompt generates the prompt for AI chatbots to generate names.
func GeneratePrompt(purpose, demographics, interests, theme string) (string, error) {
	data := promptData{
		Purpose:      purpose,
		Demographics: demographics,
		Interests:    interests,
		Theme:        theme,
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
