package sub

import (
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/urfave/cli/v2"
)

// GeneratePromptAction generates a prompt for AI chatbot users.
func GeneratePromptAction(c *cli.Context) error {
	purpose, err := getInput(c, "purpose", "👋 Enter the project purpose: ")
	if err != nil {
		return err
	}

	demographics, err := getInput(c, "demographics", "👋 Enter the audience demographics: ")
	if err != nil {
		return err
	}

	interests, err := getInput(c, "audience", "👋 Enter the audience interests: ")
	if err != nil {
		return err
	}

	theme, err := getInput(c, "theme", "👋 Enter the project theme: ")
	if err != nil {
		return err
	}

	data := struct {
		Purpose      string
		Demographics string
		Interests    string
		Theme        string
	}{
		Purpose:      purpose,
		Demographics: demographics,
		Interests:    interests,
		Theme:        theme,
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
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("🍺 Final result ->\n\n%s\n\n", strings.TrimSpace(builder.String()))
	fmt.Println("🎉 Copy into the AI of your choice!")

	return nil
}
