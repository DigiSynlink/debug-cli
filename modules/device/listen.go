package device

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/urfave/cli/v2"
)

var announceAddr = "239.255.255.254:9900"

type DeviceAnnouncement struct {
	Action          string `json:"action"`
	Active          string `json:"active"`
	Hostname        string `json:"hostname"`
	DeviceName      string `json:"device_name"`
	DeviceCategory  string `json:"device_category"`
	DeviceType      string `json:"device_type"`
	ChannelNumber   string `json:"channel_num"`
	InChannelNames  string `json:"in_channel_names"`
	OutChannelNames string `json:"out_channel_names"`
	IP              string `json:"ip"`
	IPConflict      string `json:"ip_conflict"`
	OnlineVersion   string `json:"online_version"`
}

func (d *DeviceAnnouncement) cleanUp() {
	name, err := url.QueryUnescape(d.DeviceName)
	if err != nil {
		logger.Error("Error decoding the string %v", err)
		return
	}
	d.DeviceName = name
}

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
	res := make(map[string]DeviceAnnouncement)

	for {
		logger.Info("Waiting for a response... 1min Deadline")
		pc.SetReadDeadline(time.Now().Add(time.Minute * 1))
		n, dst, err := pc.ReadFrom(buf)
		if n > 0 {
			logger.Debug(fmt.Sprintf("%s sent this: %s\n", dst, buf[:n]))

			// Trying to parse response
			announcement := DeviceAnnouncement{}
			err := json.Unmarshal(buf[:n], &announcement)
			if err != nil {
				logger.Error("Error parsing announcement: ", err)
			} else {
				announcement.cleanUp()
				if _, ok := res[announcement.DeviceName]; !ok {
					res[announcement.DeviceName] = announcement
					logger.Info("Device: ", announcement.DeviceName, "(", announcement.IP, ")", " is online")
				}
			}
		}

		if len(res) > 0 {
			logger.Info("Current Online Machines:")
			for k, v := range res {
				logger.Info("\t", k, ": ", v)
			}
		} else {
			logger.Warn("No online devices at this time...")
		}

		if err != nil {
			logger.Info("Exit with Error: ", err)
			logger.Info("It may be a timeout, so it's not a problem. Try again...")
		}
	}
}
