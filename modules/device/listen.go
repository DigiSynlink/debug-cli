package device

import (
	"fmt"
	"net"
	"time"

	"github.com/urfave/cli/v2"
)

var announceAddr = "239.255.255.254:9900"

func ListenEntry(cCtx *cli.Context) error {
	ifi := cCtx.String("interface")
	announceAddr = cCtx.String("listen-address")

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

	return Listen(ifi)
}

func Listen(ifi string) error {
	logger.Info("Listening on interface: ", ifi)
	logger.Info("Announce address: ", announceAddr)
	addr, err := net.ResolveUDPAddr("udp4", announceAddr)

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

	pc, err := net.ListenMulticastUDP("udp4", iface, addr)

	if err != nil {
		return err
	}

	defer pc.Close()

	buf := make([]byte, 1024)
	logger.Warn("To stop Listen, press Ctrl+C")
	for {
		logger.Info("Waiting for a response... 30s Deadline")
		pc.SetReadDeadline(time.Now().Add(time.Second * 30))
		n, dst, err := pc.ReadFrom(buf)
		if n > 0 {
			fmt.Printf("%s sent this: %s\n", dst, buf[:n])
		}

		if err != nil {
			logger.Info("Exit with Error: ", err)
			logger.Info("It may be a timeout, but it's not a problem. Try again...")
		}
	}
}
