package core

import (
	"flag"
	"strings"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestGetStringFromReader_CLIValue(t *testing.T) {
	flagName := "flag"
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{Name: flagName},
		},
	}
	set := flag.NewFlagSet("test", 0)
	set.String(flagName, "testvalue", "")
	c := cli.NewContext(app, set, nil)

	result, err := GetStringFromReader(c, flagName, "prompt", strings.NewReader(""))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "testvalue" {
		t.Errorf("expected 'testvalue', got '%s'", result)
	}
}

func TestGetStringFromReader_ReaderInput(t *testing.T) {
	app := &cli.App{}
	c := &cli.Context{}
	c.App = app

	result, err := GetStringFromReader(c, "flag", "prompt", strings.NewReader("input\n"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "input" {
		t.Errorf("expected 'input', got '%s'", result)
	}
}

func TestGetStringFromReader_EmptyAfterTrim(t *testing.T) {
	app := &cli.App{}
	c := &cli.Context{}
	c.App = app

	_, err := GetStringFromReader(c, "flag", "prompt", strings.NewReader("   \n"))
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
