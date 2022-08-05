package netif

import (
	"github.com/urfave/cli/v2"
)

func RegisterCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "interface",
		Usage: "Utility for finding interface parameters, especially for windows user",
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"r"},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "verbose",
						Usage: "Showing more information with Pcap(currently not implemented due to cgo dependency)",
						Value: false,
					},
				},
				Usage:  "List all network interfaces",
				Action: ListEntry,
			},
		},
	})

}
