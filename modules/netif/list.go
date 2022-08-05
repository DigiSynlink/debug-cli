package netif

import (
	"fmt"
	"net"

	"github.com/digisynlink/debug-cli/utils"
	"github.com/google/gopacket/pcap"
	"github.com/urfave/cli/v2"
)

var logger = utils.GetInstance()

func ListEntry(cCtx *cli.Context) error {
	return list(cCtx.Bool("verbose"))
}

func list(verbose bool) error {
	logger.Info("Devices found:")
	if verbose {
		devices, err := pcap.FindAllDevs()
		if err != nil {
			return err
		}
		logger.Debug(devices)
		for _, device := range devices {
			fmt.Println()
			fmt.Println("========= Start Interface ==========")
			fmt.Println("Name: ", device.Name)
			fmt.Println("Description: ", device.Description)
			for _, address := range device.Addresses {
				fmt.Println("- IP address: ", address.IP)
			}
			fmt.Println("=========== End ==============")
		}
	} else {
		infs, _ := net.Interfaces()
		for _, f := range infs {
			fmt.Println("- '" + f.Name + "'")
		}
	}

	return nil
}
