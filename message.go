package dhcp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"slices"
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

type BOOTPMessageType uint8

const (
	BOOTPMessageTypeRequest BOOTPMessageType = 1
	BOOTPMessageTypeReply   BOOTPMessageType = 2
)

var BOOTPMessageTypeToString map[BOOTPMessageType]string = map[BOOTPMessageType]string{
	BOOTPMessageTypeRequest: "REQUEST",
	BOOTPMessageTypeReply:   "REPLY",
}

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

func (msg Message) String() string {
	var sb strings.Builder

	sb.WriteString("BOOTP Message Type: ")
	sb.WriteString(BOOTPMessageTypeToString[msg.BOOTPMessageType])

	sb.WriteString("\nHardware Type: ")

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

	sb.WriteString("\n\n--Options--\n")

	for _, opt := range msg.Options {
		fmt.Fprintf(&sb, "\n%s", OptionCodeToInfo[opt.Code].String)
		fmt.Fprintf(&sb, " [%s]: ", strconv.Itoa(int(opt.Code)))
		sb.WriteString(opt.Data.String())
	}
	fmt.Fprintf(&sb, "\nEND [%s]", strconv.Itoa(int(OptionCodeEnd)))

	return sb.String()
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
	data := make([]byte, 240)

	data[0] = byte(msg.BOOTPMessageType)
	data[1] = byte(msg.HardwareAddrType)
	data[2] = byte(msg.HardwareAddrLen)
	data[3] = byte(msg.HopCount)
	binary.BigEndian.PutUint32(data[4:8], msg.TransactionID)
	binary.BigEndian.PutUint16(data[8:10], msg.SecsElapsed)
	binary.BigEndian.PutUint16(data[10:12], msg.Flags)
	copy(data[12:16], []byte(msg.ClientIPAddr))
	copy(data[16:20], []byte(msg.YourIPAddr))
	copy(data[20:24], []byte(msg.ServerIPAddr))
	copy(data[24:28], []byte(msg.GatewayIPAddr))
	copy(data[28:44], []byte(msg.ClientHardwareAddr))
	copy(data[44:108], msg.ServerHostname)
	copy(data[108:236], msg.BootFilename)
	copy(data[236:240], MagicCookie)
	data = append(data, msg.Options.Unmarshal()...)

	paddingLen := 0
	if len(data) < 300 {
		paddingLen = 300 - len(data)
	}

	return append(data, make([]byte, paddingLen)...)
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
