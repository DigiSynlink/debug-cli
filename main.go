package main

import (
	"fmt"
	"os"

	"github.com/digisynlink/debug-cli/modules/find"
	"github.com/digisynlink/debug-cli/utils"

	"github.com/urfave/cli/v2"
)

var gitCommit string
var buildDate string

func main() {
	logger := utils.GetInstance()
	app := &cli.App{
		Name:  "network-debug",
		Usage: "DigitSynlink network debug tool",
		Commands: []*cli.Command{
			{
				Name:    "find",
				Aliases: []string{"f"},
				Usage:   "Find a device",
				Action: func(cCtx *cli.Context) error {
					return find.LookForDevice()
				},
			},
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Show version",
				Action: func(cCtx *cli.Context) error {
					fmt.Printf("Commit: %s \nBuildDate: %s\n", gitCommit, buildDate)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
