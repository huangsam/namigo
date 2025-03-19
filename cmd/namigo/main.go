package main

import (
	"fmt"
	"os"

	"github.com/huangsam/namigo/cmd/namigo/sub"
	"github.com/huangsam/namigo/internal/model"
	"github.com/urfave/cli/v2"
)

// checkSizeFlag checks for valid size flag.
func checkSizeFlag(ctx *cli.Context, i int) error {
	if i <= 0 {
		return fmt.Errorf("size %d is invalid", i)
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

func main() {
	app := &cli.App{
		Name:  "namigo",
		Usage: "Your naming pal, written in Go 🐶",
		Commands: []*cli.Command{
			{ // Lots of goodies to come. Stay tuned!
				Name:  "generate",
				Usage: "Generate names with AI",
				Subcommands: []*cli.Command{
					{
						Name:  "prompt",
						Usage: "Generate prompt for AI chatbots",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "purpose",
								Usage: "Project purpose",
							},
							&cli.StringFlag{
								Name:  "theme",
								Usage: "Project theme",
							},
							&cli.StringFlag{
								Name:  "demographics",
								Usage: "Target demographics",
							},
							&cli.StringFlag{
								Name:  "interests",
								Usage: "Target interests",
							},
							&cli.IntFlag{
								Name:   "size",
								Usage:  "Max size for generated names",
								Value:  10,
								Action: checkSizeFlag,
							},
							&cli.IntFlag{
								Name:   "length",
								Usage:  "Max length for each generated name",
								Value:  20,
								Action: checkLengthFlag,
							},
						},
						Action: sub.GeneratePromptAction,
					},
				},
			},
			{
				Name:  "search",
				Usage: "Search for terms across entities",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:   "size",
						Usage:  "Max size for search results",
						Value:  10,
						Action: checkSizeFlag,
					},
					&cli.StringFlag{
						Name:   "format",
						Usage:  fmt.Sprintf("Output format can be %s or %s", model.TextValue, model.JSONValue),
						Value:  model.TextValue,
						Action: checkFormatFlag,
					},
				},
				Subcommands: []*cli.Command{
					{
						Name:   "package",
						Usage:  "Search for packages",
						Action: sub.SearchPackageAction,
					},
					{
						Name:   "dns",
						Usage:  "Search for DNS records",
						Action: sub.SearchDNSAction,
					},
					{
						Name:   "email",
						Usage:  "Search for email records",
						Action: sub.SearchEmailAction,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("💥 Error: %v\n", err.Error())
		os.Exit(1)
	}
}
