package dhcp

import (
	"encoding/binary"
	"errors"
	"net"
	"slices"
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

type BOOTPMessageType uint8

const (
	BOOTPMessageTypeRequest BOOTPMessageType = 1
	BOOTPMessageTypeReply   BOOTPMessageType = 2
)

type HardwareAddrType uint8

const HardwareAddrTypeEthernet = 1

var HardwareAddrLengths = map[HardwareAddrType]uint8{
	HardwareAddrTypeEthernet: 6,
}

func (hw HardwareAddrType) ValidLength() uint8 {
	return HardwareAddrLengths[hw]
}

func HasMagicCookie(data []byte) bool {
	if slices.Equal(data[236:240], MagicCookie) {
		return true
	}
	return false
}

func IsEncodedMessage(data []byte) bool {
	if len(data) < 300 {
		return false
	}
	if !HasMagicCookie(data) {
		return false
	}

	return true
}

func MarshalMessage(data []byte) (*Message, error) {
	if !IsEncodedMessage(data) {
		return nil, errors.New("data does not contain an encoded dhcp message")
	}

	msg := &Message{}

	// MessageType
	switch BOOTPMessageType(data[0]) {
	case BOOTPMessageTypeRequest, BOOTPMessageTypeReply:
		msg.BOOTPMessageType = BOOTPMessageType(data[0])
	default:
		return nil, errors.New("bootp message type is invalid")
	}

	// HardwareAddrType
	switch HardwareAddrType(data[1]) {
	case HardwareAddrTypeEthernet:
		msg.HardwareAddrType = HardwareAddrType(data[1])
	default:
		return nil, errors.New("hardware type is invalid")
	}

	// HardwareAddrLen
	if uint8(data[2]) != msg.HardwareAddrType.ValidLength() {
		return nil, errors.New("hardware address length does not match hardware type")
	}
	msg.HardwareAddrLen = uint8(data[2])

	// HopCount
	msg.HopCount = uint8(data[3])

	// TransactionID
	msg.TransactionID = binary.BigEndian.Uint32(data[4:8])

	// SecsElapsed
	msg.SecsElapsed = binary.BigEndian.Uint16(data[8:10])

	// Flags
	msg.Flags = binary.BigEndian.Uint16(data[10:12])

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
			return nil, errors.New("client hardware address extends 6 bytes")
		}
	}
	msg.ClientHardwareAddr = net.HardwareAddr(temp[:6])

	// ServerHostname
	msg.ServerHostname = string(data[44:108])

	// BootFilename
	msg.BootFilename = string(data[108:236])

	// Options
	msg.Options, _ = MarshalOptions(data[240:])

	return msg, nil
}

func (msg *Message) Unmarshal() []byte {
        var data []byte

        data = append(data,
                byte(msg.BOOTPMessageType),
                byte(msg.HardwareAddrType),
                byte(msg.HardwareAddrLen),
                byte(msg.HopCount),
        )
        binary.BigEndian.AppendUint32(data, msg.TransactionID)
        binary.BigEndian.AppendUint16(data, msg.SecsElapsed)
        binary.BigEndian.AppendUint16(data, msg.Flags)
        data = append(data, []byte(msg.ClientIPAddr)...)
        data = append(data, []byte(msg.YourIPAddr)...)
        data = append(data, []byte(msg.ServerIPAddr)...)
        data = append(data, []byte(msg.GatewayIPAddr)...)
        data = append(data, []byte(msg.ClientHardwareAddr)...)
        sHostname := [64]byte{}
        copy(sHostname[:], msg.ServerHostname)
        data = append(data, sHostname[:]...)
        bootFilename := [128]byte{}
        copy(bootFilename[:], msg.BootFilename)
        data = append(data, bootFilename[:]...)
        data = append(data, MagicCookie...)
        data = append(data, msg.Options.Unmarshal()...) 

        return data
}
