package sub

import (
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/urfave/cli/v2"
)

// PromptData is the input for the prompt template.
type PromptData struct {
	Purpose      string
	Demographics string
	Interests    string
	Theme        string
}

// GeneratePromptAction generates a prompt for AI chatbot users.
func GeneratePromptAction(c *cli.Context) error {
	purpose, err := getInput(c, "purpose", "👋 Enter project purpose: ")
	if err != nil {
		return err
	}

	theme, err := getInput(c, "theme", "👋 Enter project theme: ")
	if err != nil {
		return err
	}

	demographics, err := getInput(c, "demographics", "👋 Enter target demographics: ")
	if err != nil {
		return err
	}

	interests, err := getInput(c, "audience", "👋 Enter target interests: ")
	if err != nil {
		return err
	}

	data := PromptData{
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
