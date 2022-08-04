package device

import (
	"fmt"
	"net"
	"time"

	"github.com/digisynlink/debug-cli/utils"
	"github.com/urfave/cli/v2"
)

var boardcastAddr = "239.255.255.252:9900"
var magicString = "AreYouDigisynlink?"

func DiscoverEntry(cCtx *cli.Context) error {
	ifi := cCtx.String("interface")
	magicString = cCtx.String("magic-string")
	boardcastAddr = cCtx.String("boardcast-address")
	logger.Info("Boardcast address: ", boardcastAddr)

	if ifi == "" {
		logger.Warn("No interface specified, system will assign random interface for boardcast, which is not desired in most cases.")
	} else {
		has, err := HasInterface(ifi)
		if err != nil {
			return err
		}
		if !has {
			return fmt.Errorf("interface %s not found", ifi)
		}
	}

	return LookForDevice(ifi)
}

func LookForDevice(ifi string) error {
	logger := utils.GetInstance()

	addr, err := net.ResolveUDPAddr("udp4", boardcastAddr)
	if err != nil {
		return err
	}

	var iface *net.Interface
	if ifi != "" {
		iface, err = net.InterfaceByName(ifi)
		if err != nil {
			return err
		}
	}

	conn, err := net.ListenMulticastUDP("udp4", iface, addr)
	if err != nil {
		return err
	}

	defer conn.Close()

	// Send a boardcase
	logger.Info("Sending boardcast...")
	_, err = conn.WriteTo([]byte(magicString), addr)

	if err != nil {
		return err
	}

	buf := make([]byte, 1024)

	for {
		logger.Info("Waiting for a response... 30s Deadline")
		conn.SetReadDeadline(time.Now().Add(time.Second * 30))
		n, dst, err := conn.ReadFrom(buf)
		if n > 0 {
			fmt.Printf("%s sent this: %s\n", dst, buf[:n])
		}

		if err != nil {
			logger.Info("Exit with Error: ", err)
			logger.Info("It may be a timeout, so we will exit now. Restart the program to try again.")
			return nil
		}
	}
}
