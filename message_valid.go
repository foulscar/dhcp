package dhcp

import (
	"net"
)

func IsEncodedMessage(data []byte) bool {
	if len(data) < 300 {
		return false
	}
	if !HasMagicCookie(data) {
		return false
	}

	return true
}

func (msg Message) IsValid() (valid bool, reason string) {
	switch msg.BOOTPMessageType {
	case BOOTPMessageTypeRequest, BOOTPMessageTypeReply:
	default:
		return false, "invalid bootp message type"
	}

	_, exists := HardwareAddrTypeToString[msg.HardwareAddrType]
	if !exists {
		return false, "invalid hardware type"
	}

	if msg.HardwareAddrLen != msg.HardwareAddrType.ValidLength() {
		return false, "hardware address length does not match hardware type"
	}

	_, exists = FlagsToString[msg.Flags]
	if !exists {
		return false, "invalid flags"
	}

	switch msg.HardwareAddrType {
	case HardwareAddrTypeEthernet:
		if len(msg.ClientHardwareAddr) != 6 {
			return false, "client hardware address is invalid"
		}
	}

	if !msg.Options.IsValid() {
		return false, "options are invalid"
	}

	msgTypeCode, err := msg.GetMessageType()
	if err != nil {
		return false, "option message type is required and must be valid"
	}

	if !msgTypeCode.MatchesBOOTPMessageType(msg.BOOTPMessageType) {
		return false, "bootp message type does not match message type option"
	}

	return true, "ok"
}
