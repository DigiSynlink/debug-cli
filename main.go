package main

import (
	"os"
	"time"

	"github.com/digisynlink/debug-cli/modules/device"
	"github.com/digisynlink/debug-cli/modules/version"
	"github.com/digisynlink/debug-cli/utils"
	"github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

var logger = utils.GetInstance()

func main() {
	app := &cli.App{
		Name:      "debug-cli",
		Version:   version.Version,
		Compiled:  time.Now(),
		Copyright: "Copyright © 2020 digisynlink",
		Usage:     "DigitSynlink network debug tool",
		Commands:  []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "debug-level",
				Aliases: []string{"lvl"},
				Usage:   "Debug level, available values: debug, info, warn, error",
				Value:   "info",
			},
		},
		Before: func(cCtx *cli.Context) error {
			parsed, err := logrus.ParseLevel(cCtx.String("debug-level"))
			if err != nil {
				return err
			}
			logger.SetLevel(parsed)
			return nil
		},
	}

	device.RegisterCommand(app)
	version.RegisterCommand(app)

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}
