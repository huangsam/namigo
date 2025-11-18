package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// GetString gets string from CLI or manual input.
func GetString(c *cli.Context, flag, prompt string) (string, error) {
	return GetStringFromReader(c, flag, prompt, os.Stdin)
}

// GetStringFromReader gets string from CLI or reader input.
func GetStringFromReader(c *cli.Context, flag, prompt string, reader io.Reader) (string, error) {
	value := c.String(flag)
	if value == "" {
		fmt.Printf("%s: ", prompt)
		scanner := bufio.NewScanner(reader)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return "", err
			} else {
				return "", errors.New("no input provided")
			}
		}
		value = strings.TrimSpace(scanner.Text())
		if value == "" {
			return "", errors.New(flag + " is empty after space trimming")
		}
	}
	return value, nil
}
