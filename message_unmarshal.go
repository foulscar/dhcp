package dhcp

import (
	"encoding/binary"
	"net"
)

// UnmarshalMessage parses the dhcpv4 encoded message from data and returns a pointer to a Message.
// It will first check if it is a dhcp message but will not test full validity (use Message.IsValid())
func UnmarshalMessage(data []byte) (*Message, *ErrorExt) {
	mainErr := NewErrorExt("could not unmarshal message from data")

	if !IsEncodedMessage(data) {
		mainErr.Add("data does not contain an encoded dhcp message")
	}

	msg := &Message{}

	// MessageType
	switch BOOTPMessageType(data[0]) {
	case BOOTPMessageTypeRequest, BOOTPMessageTypeReply:
		msg.BOOTPMessageType = BOOTPMessageType(data[0])
	default:
		mainErr.Add("bootp message type is invalid")
	}

	// HardwareAddrType
	switch HardwareAddrType(data[1]) {
	case HardwareAddrTypeEthernet:
		msg.HardwareAddrType = HardwareAddrType(data[1])
	default:
		mainErr.Add("hardware type is invalid")
	}

	// HardwareAddrLen
	if uint8(data[2]) != msg.HardwareAddrType.ValidLength() {
		mainErr.Add("hardware address length does not match hardware type")
	}
	msg.HardwareAddrLen = uint8(data[2])

	// HopCount
	msg.HopCount = uint8(data[3])

	// TransactionID
	msg.TransactionID = binary.BigEndian.Uint32(data[4:8])

	// SecsElapsed
	msg.SecsElapsed = binary.BigEndian.Uint16(data[8:10])

	// Flags
	msg.Flags = Flags(binary.BigEndian.Uint16(data[10:12]))

	// ClientIPAddr
	msg.ClientIPAddr = net.IP(data[12:16])

	// YourIPAddr
	msg.YourIPAddr = net.IP(data[16:20])

	// ServerIPAddr
	msg.ServerIPAddr = net.IP(data[20:24])

	// GatewayIPAddr
	msg.GatewayIPAddr = net.IP(data[24:28])

	// ClientHardwareAddr
	temp := data[28:44]
	for _, b := range temp[6:] {
		if b != 0x00 {
			mainErr.Add("client hardware address extends 6 bytes")
		}
	}
	msg.ClientHardwareAddr = net.HardwareAddr(temp[:6])

	// ServerHostname
	msg.ServerHostname = string(data[44:108])

	// BootFilename
	msg.BootFilename = string(data[108:236])

	// Options
	opts, err := UnmarshalOptions(data[240:])
	mainErr.Add(err)
	msg.Options = opts

	if mainErr.HasReasons() {
		return nil, mainErr
	}

	return msg, nil
}
