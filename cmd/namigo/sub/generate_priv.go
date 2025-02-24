package sub

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// getInput gets input for prompt generation.
func getInput(c *cli.Context, flag, prompt string) (string, error) {
	value := c.String(flag)
	if value == "" {
		value = getInputHelper(prompt)
		if value == "" {
			return "", errors.New(flag + " is required")
		}
	}
	return value, nil
}

// getInputHelper prompts the user for input.
func getInputHelper(prompt string) string {
	fmt.Printf("%s: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	response, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(response)
}
