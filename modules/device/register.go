package device

import (
	"github.com/digisynlink/debug-cli/utils"
	"github.com/urfave/cli/v2"
)

var logger = utils.GetInstance()

func RegisterCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:    "device",
		Aliases: []string{"d"},
		Usage:   "Device Actions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "interface",
				Usage:       "Interface to use",
				DefaultText: "eth0",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:    "discover",
				Aliases: []string{"d"},
				Usage:   "Device Discovery",
				Action:  DiscoverEntry,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "boardcast-address",
						Aliases: []string{"baddr"},
						Usage:   "Address to use",
						Value:   boardcastAddr,
					},
					&cli.StringFlag{
						Name:  "magic-string",
						Usage: "Magic string to use",
						Value: magicString,
					},
				},
			},
			{
				Name:    "listen",
				Aliases: []string{"l"},
				Usage:   "Listen to Announcement",
				Action:  ListenEntry,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "listen-address",
						Aliases: []string{"laddr"},
						Usage:   "Address to use",
						Value:   announceAddr,
					},
				},
			},
		},
	})
}
