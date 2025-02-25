package main

import (
	"fmt"

	"github.com/huangsam/namigo/internal/model"
	"github.com/urfave/cli/v2"
)

// checkCountFlag checks for valid count flag.
func checkCountFlag(ctx *cli.Context, i int) error {
	if i <= 0 {
		return fmt.Errorf("count %d is invalid", i)
	}
	return nil
}

// checkLengthFlag checks for valid length flag.
func checkLengthFlag(ctx *cli.Context, i int) error {
	if i <= 0 {
		return fmt.Errorf("length %d is invalid", i)
	}
	return nil
}

// checkFormatFlag checks for valid format flag.
func checkFormatFlag(ctx *cli.Context, s string) error {
	if s != model.TextValue && s != model.JSONValue {
		return fmt.Errorf("format %s is invalid", s)
	}
	return nil
}
