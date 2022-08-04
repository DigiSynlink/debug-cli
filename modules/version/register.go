package version

import (
	"github.com/urfave/cli/v2"
)

func RegisterCommand(app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Show version",
		Action: func(cCtx *cli.Context) error {
			ShowVersion()
			return nil
		},
	})

}
