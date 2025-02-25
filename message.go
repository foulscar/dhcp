package dhcp

import (
	"net"
        "slices"
)

type MsgTypeCode byte

const (
	MsgTypeRequest MsgTypeCode = 0x01
	MsgTypeReply   MsgTypeCode = 0x02
)

type Message struct {
	MessageType        MsgTypeCode
	HardwareAddrType   byte
	HardwareAddrLen    uint8
	HopCount           uint8
	TransactionID      uint32
	SecsElapsed        uint16
	Flags              uint16
	ClientIPAddr       net.IP
	YourIPAddr         net.IP
	ServerIPAddr       net.IP
	GatewayIPAddr      net.IP
	ClientHardwareAddr net.HardwareAddr
	ServerHostname     string
	BootFilename       string
	Options            Options
}

func isDHCPMessage(data []byte) bool {
        if slices.Equal(data[236:240], MagicCookie) {
                return true
        }
        return false
}
