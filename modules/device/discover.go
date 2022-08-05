package device

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/digisynlink/debug-cli/utils"
	"github.com/urfave/cli/v2"
)

func DiscoverEntry(cCtx *cli.Context) error {
	ifi := cCtx.String("interface")
	magicString = cCtx.String("magic-string")
	boardcastAddr = cCtx.String("boardcast-address")
	bindAddr := cCtx.String("bind-address")

	if ifi == "" && bindAddr == "" {
		return errors.New("please specify either --interface or --bind-address")
	}

	if ifi != "" && bindAddr != "" {
		return errors.New("please specify either --interface or --bind-address")
	}

	logger.Info("Boardcast address: ", boardcastAddr)

	return LookForDevice(bindAddr, ifi)
}

func LookForDevice(bindAddr string, iface string) error {
	logger := utils.GetInstance()

	if iface != "" {
		logger.Info("Interface specified, using auto-binding from interface: ", iface)

		ifi, err := GetNetworkInterface(iface)
		if err != nil {
			return err
		}

		addrs, err := ifi.Addrs()
		if err != nil {
			return err
		}

		if len(addrs) == 0 {
			logger.Error("No addresses found associated with interface, unable to redirect boardcast. Retry after a default DHCP address to be assigned.")
			return nil
		}

		found := false
		for _, addr := range addrs {
			addr_str := addr.String()
			if strings.Contains(addr_str, ":") {
				logger.Debug("Found IPv6 address: ", addr_str, " skipping...")
				continue
			}
			ip := strings.Split(addr_str, "/")
			if len(ip) == 2 {
				found = true
				bindAddr = ip[0]
				logger.Info("Binding to address: ", bindAddr)
				break
			}
		}

		if !found {
			logger.Error("No ipv4 addresses found associated with interface, unable to redirect boardcast. Retry after a default DHCP address to be assigned.")
			return nil
		}
	}

	conn, err := net.ListenPacket("udp4", bindAddr+":"+port)

	if err != nil {
		logger.Debug("Error while ListenPacket")
		return err
	}

	addr, err := net.ResolveUDPAddr("udp4", boardcastAddr+":"+port)
	if err != nil {
		logger.Debug("Error while Resolve UDP address")
		return err
	}

	defer conn.Close()

	logger.Info("Sending boardcast...")
	n, err := conn.WriteTo([]byte(magicString), addr)
	logger.Debug("Sent ", n, " bytes")

	if err != nil {
		return err
	}
	buf := make([]byte, 1024)

	for {
		logger.Info("Waiting for responses... 1min Deadline")
		conn.SetReadDeadline(time.Now().Add(time.Minute))
		n, dst, err := conn.ReadFrom(buf)
		if n > 0 {
			logger.Debug(fmt.Sprintf("%s sent this: %s\n", dst, buf[:n]))
			res, err := ParseDeviceEcho(buf[:n])
			if err != nil {
				logger.Warn("Error while parsing device echo: ", err)
			} else {
				logger.Info("Found device: ", res.BoardModel, " at: ", res.ip)
			}
		}

		if err != nil {
			logger.Info("Exit with Error: ", err)
			logger.Info("It may be a timeout, so we will exit now. Restart the program to try again.")
			return err
		}
	}
}
