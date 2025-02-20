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
				Usage: "Search term for finding packages",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "max",
						Usage: "Max number of results to display",
						Value: 10,
					},
				},
				Action: sub.SearchAction,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
