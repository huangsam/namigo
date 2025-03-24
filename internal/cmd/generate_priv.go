package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// GetString gets string from CLI or manual input.
func GetString(c *cli.Context, flag, prompt string) (string, error) {
	value := c.String(flag)
	if value == "" {
		fmt.Printf("%s: ", prompt)
		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		value = strings.TrimSpace(response)
		if value == "" {
			return "", errors.New(flag + " is empty after space trimming")
		}
	}
	return value, nil
}
