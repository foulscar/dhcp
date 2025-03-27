package dhcp

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Message represents a DHCPv4 packet header and options
type Message struct {
	// REQUEST or REPLY
	BOOTPMessageType BOOTPMessageType

	// Type of the Client Hardware Address (normally ethernet/mac)
	HardwareAddrType HardwareAddrType

	// Length of the Hardware Address
	HardwareAddrLen uint8

	// Number of hops made to reach a client/server (set by relay agents)
	HopCount uint8

	// Unique ID for the current message exchange
	TransactionID uint32

	// Number of seconds since a client has started the DHCP process
	SecsElapsed uint16

	// Used to indicate if a server can use unicast or if it must use broadcast
	Flags Flags

	// Client IPv4 Address if known
	ClientIPAddr net.IP

	// "Your" IPv4 Address (assigned to client)
	YourIPAddr net.IP

	// Server IPv4 Address
	ServerIPAddr net.IP

	// Relay Agent IPv4 Address
	GatewayIPAddr net.IP

	// Client Hardware Address
	ClientHardwareAddr net.HardwareAddr

	// Optional Server Hostname
	ServerHostname string

	// Optional Boot Filename (used for PXE Boot)
	BootFilename string

	// Option Fields mapped by their OptionCode
	Options Options
}

// String parses msg and returns a verbose, multi-line string of the entire Message and it's Options
func (msg Message) String() string {
	var sb strings.Builder

	sb.WriteString("BOOTP Message Type: ")
	sb.WriteString(msg.BOOTPMessageType.String())

	sb.WriteString("\nHardware Type: ")
	sb.WriteString(msg.HardwareAddrType.String())

	sb.WriteString("\nHardware Address Length: ")
	sb.WriteString(strconv.Itoa(int(msg.HardwareAddrLen)))

	sb.WriteString("\nHops: ")
	sb.WriteString(strconv.Itoa(int(msg.HopCount)))

	sb.WriteString("\nTransactionID: ")
	fmt.Fprintf(&sb, "%#x", msg.TransactionID)

	sb.WriteString("\nSeconds Elapsed: ")
	sb.WriteString(strconv.Itoa(int(msg.SecsElapsed)))

	sb.WriteString("\nFlags: ")
	fmt.Fprintf(&sb, "%s [%#x]", msg.Flags.String(), uint16(msg.Flags))

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

// NewMessage is a helper function that constructs a default Message that you can alter.
// This function uses pre-defined constructors, therefor if you are not
// using default mappings in GlobalOptionCodeMapping, you should not use this
func NewMessage() Message {
	msg := Message{}
	msg.BOOTPMessageType = BOOTPMessageTypeReply
	msg.HardwareAddrType = HardwareAddrTypeEthernet
	msg.HardwareAddrLen = HardwareAddrTypeEthernet.ValidLength()
	msg.Flags = FlagsBroadcast
	msg.ClientIPAddr = net.IPv4zero
	msg.YourIPAddr = net.IPv4zero
	msg.ServerIPAddr = net.IPv4zero
	msg.GatewayIPAddr = net.IPv4zero
	msg.ClientHardwareAddr = make(net.HardwareAddr, 6)
	msg.Options = make(Options)

	optMsgType, _ := NewOptionMessageType(OptionMessageTypeCodeNACK)

	msg.Options.Add(optMsgType)

	return msg
}
