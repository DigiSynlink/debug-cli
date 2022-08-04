package find

import (
	"fmt"
	"net"
	"time"

	"github.com/digisynlink/debug-cli/utils"
)

var boardcastAddr = "224.0.1.129:9999"
var magicString = "AreYouDigisynlink?"

func LookForDevice() error {
	logger := utils.GetInstance()
	pc, err := net.ListenPacket("udp4", ":9999")
	if err != nil {
		return err
	}
	defer pc.Close()

	addr, err := net.ResolveUDPAddr("udp4", boardcastAddr)
	if err != nil {
		return err
	}

	// Send a boardcase
	logger.Info("Sending boardcast...")
	_, err = pc.WriteTo([]byte(magicString), addr)

	if err != nil {
		return err
	}

	buf := make([]byte, 1024)
	for {
		logger.Info("Waiting for a response...")
		pc.SetReadDeadline(time.Now().Add(time.Second * 10))
		n, dst, err := pc.ReadFrom(buf)
		if n > 0 {
			fmt.Printf("%s sent this: %s\n", dst, buf[:n])
		}

		if err != nil {
			logger.Info("Exit with Error: ", err)
			logger.Info("It may be a timeout, but it's not a problem. Restart the program and try again.")
			return nil
		}
	}
}
