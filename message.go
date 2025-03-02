package dhcp

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Message struct {
	BOOTPMessageType   BOOTPMessageType
	HardwareAddrType   HardwareAddrType
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

func (msg Message) String() string {
	var sb strings.Builder

	sb.WriteString("BOOTP Message Type: ")
	sb.WriteString(BOOTPMessageTypeToString[msg.BOOTPMessageType])

	sb.WriteString("\nHardware Type: ")
	sb.WriteString(HardwareAddrTypeToString[msg.HardwareAddrType])

	sb.WriteString("\nHardware Address Length: ")
	sb.WriteString(strconv.Itoa(int(msg.HardwareAddrLen)))

	sb.WriteString("\nHops: ")
	sb.WriteString(strconv.Itoa(int(msg.HopCount)))

	sb.WriteString("\nTransactionID: ")
	fmt.Fprintf(&sb, "%#x", msg.TransactionID)

	sb.WriteString("\nSeconds Elapsed: ")
	sb.WriteString(strconv.Itoa(int(msg.SecsElapsed)))

	sb.WriteString("\nFlags: ")
	fmt.Fprintf(&sb, "%#x", msg.Flags)

	sb.WriteString("\nClient IP Address: ")
	sb.WriteString(msg.ClientIPAddr.String())

	sb.WriteString("\nYour IP Address: ")
	sb.WriteString(msg.YourIPAddr.String())

	sb.WriteString("\nServer IP Address: ")
	sb.WriteString(msg.ServerIPAddr.String())

	sb.WriteString("\nGateway IP Address: ")
	sb.WriteString(msg.GatewayIPAddr.String())

	sb.WriteString("\nClient Mac Address: ")
	sb.WriteString(msg.ClientHardwareAddr.String())

	sb.WriteString("\nServer Hostname: ")
	fmt.Fprintf(&sb, "'%s'", msg.ServerHostname)

	sb.WriteString("\nBoot Filename: ")
	fmt.Fprintf(&sb, "'%s'", msg.ServerHostname)

	sb.WriteString("\n\n--Options--\n\n")
	sb.WriteString(msg.Options.String())

	return sb.String()
}

func NewEmptyMessage() Message {
	msg := Message{}
	msg.ClientIPAddr = net.IPv4zero
	msg.YourIPAddr = net.IPv4zero
	msg.ServerIPAddr = net.IPv4zero
	msg.GatewayIPAddr = net.IPv4zero
	msg.ClientHardwareAddr = make(net.HardwareAddr, 6)
	msg.Options = make(Options)

	return msg
}
