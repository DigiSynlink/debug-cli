package settings

import (
	"github.com/urfave/cli/v2"
)

func RegisterCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "api",
		Usage: "Send API Command",
		Flags: []cli.Flag{&cli.StringFlag{Name: "host", Usage: "Host address", Required: true}},
		Subcommands: []*cli.Command{
			{
				Name:    "reboot",
				Aliases: []string{"r"},
				Usage:   "Reboot device",
				Action:  RebootEntry,
			},
		},
	})

}
