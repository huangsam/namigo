package sub

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// promptTemplate is a Go template for prompting
const promptTemplate = `Given the following parameters:

Project Purpose: {{.Purpose}}
Audience: {{.Audience}}
Theme: {{.Theme}}

Please generate a few possible names for a side business / project.

For each of the names generated, please provide the following:

- An explanation of the name's fit for project purpose, audience and theme.
- First impressions from the audience when they hear the name.
- Pros and cons of the name.
- Any other thoughts you have about the name.

Format the output as a JSON array of objects, stack ranked based on your
assessment of their suitability. Provide a suitability score in each JSON
array element.
`

// getInput prompts the user for input.
func getInput(c *cli.Context, flag, prompt string) (string, error) {
	value := c.String(flag)
	if value == "" {
		value = promptUser(prompt)
		if value == "" {
			return "", errors.New(flag + " is required")
		}
	}
	return value, nil
}

// promptUser prompts the user for input.
func promptUser(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(response)
}
