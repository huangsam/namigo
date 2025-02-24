package sub

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/urfave/cli/v2"
)

// promptTemplate is a Go template for prompting
const promptTemplate = `Given the following parameters:

Project Purpose: {{.Purpose}}
Audience: {{.Audience}}
Theme: {{.Theme}}

Please generate a few possible names for a side business / project.
`

// GeneratePromptAction generates a prompt for AI chatbot users.
func GeneratePromptAction(c *cli.Context) error {
	purpose, err := getInput(c, "purpose", "👋 Enter the project purpose: ")
	if err != nil {
		return err
	}

	audience, err := getInput(c, "audience", "👋 Enter the project audience: ")
	if err != nil {
		return err
	}

	theme, err := getInput(c, "theme", "👋 Enter the project theme: ")
	if err != nil {
		return err
	}

	data := struct {
		Purpose  string
		Audience string
		Theme    string
	}{
		Purpose:  purpose,
		Audience: audience,
		Theme:    theme,
	}

	tmpl, err := template.New("prompt").Parse(promptTemplate)
	if err != nil {
		return err
	}

	builder := strings.Builder{}
	err = tmpl.Execute(&builder, data)
	if err != nil {
		return err
	}
	content := strings.Trim(builder.String(), " \t\n")
	fmt.Printf("🍺 Final result ->\n\n%s\n\n", content)
	fmt.Println("🎉 Copy into the AI of your choice, and see the names come!")

	return nil
}
