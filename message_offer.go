package dhcp

import (
	"net"
)

func (msg Message) isValidOFFER() (valid bool, reason string) {
	if msg.YourIPAddr.Equal(net.IPv4zero) {
		return false, "OFFER message requires 'your' ip address to be != 0.0.0.0"
	}

	return true, "ok"
}
