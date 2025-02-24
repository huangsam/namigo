package sub

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

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
