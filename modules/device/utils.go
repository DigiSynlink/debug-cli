package device

import "net"

func HasInterface(ifname string) (bool, error) {
	ifis, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	logger.Debug("Interfaces: ", ifis)

	for _, s := range ifis {
		if s.Name == ifname {
			return true, nil
		}
	}

	return false, nil
}
