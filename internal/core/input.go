package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/urfave/cli/v3"
)

// GetString gets string from CLI or manual input.
func GetString(cmd *cli.Command, flag, prompt string) (string, error) {
	return GetStringFromReader(cmd, flag, prompt, os.Stdin)
}

// GetStringFromReader gets string from CLI or reader input.
func GetStringFromReader(cmd *cli.Command, flag, prompt string, reader io.Reader) (string, error) {
	value := cmd.String(flag)
	if value == "" {
		fmt.Printf("%s: ", prompt)
		scanner := bufio.NewScanner(reader)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return "", err
			}
			return "", errors.New("no input provided")
		}
		value = strings.TrimSpace(scanner.Text())
		if value == "" {
			return "", errors.New(flag + " is empty after space trimming")
		}
	}
	return value, nil
}
