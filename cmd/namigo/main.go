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
			{
				Name:  "search",
				Usage: "Search for terms across entities",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "max",
						Usage: "Max number of results to display",
						Value: 10,
					},
					&cli.StringFlag{
						Name:  "format",
						Usage: "Output format: text/json",
						Value: "text",
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
