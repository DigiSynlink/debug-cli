package device

import (
	"fmt"
	"net"
	"strings"
)

var announceAddr = "239.255.255.254:9900"
var port = "9999"
var boardcastAddr = "255.255.255.255"
var magicString = "AreYouDigisynLink?"
var magicEcho = "iAmDigisynLink!"

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

type DeviceEcho struct {
	BoardModel   string
	BoardVersion string
	hostname     string
	ip           string
}

// iAmDigisynLink!
// /board_model=dmx208
//
// /board_version=v1.0
// hostname=DLA04-31801
//
// ip=169.254.242.216/16

func ParseDeviceEcho(data []byte) (*DeviceEcho, error) {
	str := string(data)
	str = strings.TrimSpace(str)
	strArr := strings.Split(str, "\n")

	if len(strArr) < 3 {
		return nil, fmt.Errorf("invalid data: %s", str)
	}

	if strArr[0] != magicEcho {
		return nil, fmt.Errorf("invalid data: %s", str)
	}

	echo := &DeviceEcho{}
	for _, line := range strArr {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "/board_model") {
			boardModel := strings.TrimPrefix(line, "/board_model=")
			boardModel = strings.TrimSpace(boardModel)
			echo.BoardModel = boardModel
		}
		if strings.HasPrefix(line, "/board_version") {
			boardVersion := strings.TrimPrefix(line, "/board_version=")
			boardVersion = strings.TrimSpace(boardVersion)
			echo.BoardVersion = boardVersion
		}
		if strings.HasPrefix(line, "hostname") {
			hostname := strings.TrimPrefix(line, "hostname=")
			hostname = strings.TrimSpace(hostname)
			echo.hostname = hostname
		}
		if strings.HasPrefix(line, "ip") {
			ip := strings.TrimPrefix(line, "ip=")
			ip = strings.TrimSpace(ip)
			echo.ip = ip
		}
	}

	return echo, nil
}

func GetNetworkInterface(ifi string) (iface *net.Interface, err error) {
	return net.InterfaceByName(ifi)
}
