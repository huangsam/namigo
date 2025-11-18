package core

import (
	"context"
	"strings"
	"testing"

	"github.com/urfave/cli/v3"
)

func TestGetStringFromReader_CLIValue(t *testing.T) {
	flagName := "flag"
	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{Name: flagName},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			result, err := GetStringFromReader(cmd, flagName, "prompt", strings.NewReader(""))
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if result != "testvalue" {
				t.Errorf("expected 'testvalue', got '%s'", result)
			}
			return nil
		},
	}

	err := cmd.Run(context.Background(), []string{"test", "--flag", "testvalue"})
	if err != nil {
		t.Fatalf("command run failed: %v", err)
	}
}

func TestGetStringFromReader_ReaderInput(t *testing.T) {
	cmd := &cli.Command{}
	result, err := GetStringFromReader(cmd, "flag", "prompt", strings.NewReader("input\n"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "input" {
		t.Errorf("expected 'input', got '%s'", result)
	}
}

func TestGetStringFromReader_EmptyAfterTrim(t *testing.T) {
	cmd := &cli.Command{}
	_, err := GetStringFromReader(cmd, "flag", "prompt", strings.NewReader("   \n"))
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
