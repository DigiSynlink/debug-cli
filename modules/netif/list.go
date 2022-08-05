package netif

import (
	"fmt"
	"net"

	"github.com/digisynlink/debug-cli/utils"
	"github.com/urfave/cli/v2"
)

var logger = utils.GetInstance()

func ListEntry(cCtx *cli.Context) error {
	return list(cCtx.Bool("verbose"))
}

func list(verbose bool) error {
	logger.Info("Devices found:")
	infs, _ := net.Interfaces()
	for _, f := range infs {
		fmt.Println("- '" + f.Name + "'")
	}

	return nil
}
