package main

import (
	"fmt"
	"os"

	"github.com/huangsam/namigo/cmd/namigo/sub"
	"github.com/urfave/cli/v2"
)

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
						Name:  "max",
						Usage: "Max number of results to display",
						Value: 10,
						Action: func(ctx *cli.Context, i int) error {
							if i < 0 {
								return fmt.Errorf("max %d is negative", i)
							}
							return nil
						},
					},
					&cli.StringFlag{
						Name:  "format",
						Usage: "Output format: text/json",
						Value: "text",
						Action: func(ctx *cli.Context, s string) error {
							if s != "text" && s != "json" {
								return fmt.Errorf("format %s is invalid", s)
							}
							return nil
						},
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
