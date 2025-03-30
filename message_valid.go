package dhcp

import ()

// IsEncodedMessage checks if data represents an encoded Message.
// It will not test validity, see Message.IsValid()
func IsEncodedMessage(data []byte) bool {
	if len(data) < 300 {
		return false
	}
	if !HasMagicCookie(data) {
		return false
	}

	return true
}

// IsValid checks if msg is a valid Message
func (msg Message) IsValid() *ErrorExt {
	mainErr := NewErrorExt("message is invalid")

	switch msg.BOOTPMessageType {
	case BOOTPMessageTypeRequest, BOOTPMessageTypeReply:
	default:
		mainErr.Add("invalid bootp message type")
	}

	_, exists := hardwareAddrTypeToString[msg.HardwareAddrType]
	if !exists {
		mainErr.Add("invalid hardware type")
	}

	if msg.HardwareAddrLen != msg.HardwareAddrType.ValidLength() {
		mainErr.Add("hardware address length does not match hardware type")
	}

	_, exists = flagsToString[msg.Flags]
	if !exists {
		mainErr.Add("invalid flags")
	}

	switch msg.HardwareAddrType {
	case HardwareAddrTypeEthernet:
		if len(msg.ClientHardwareAddr) != 6 {
			mainErr.Add("client hardware address is invalid")
		}
	}

	err := msg.Options.IsValid()
	if err != nil {
		mainErr.Add(err)
	}

	msgTypeCode, err := msg.GetMessageType()
	if err != nil {
		mainErr.Add("option message type is required and must be valid")
	}

	if !msgTypeCode.MatchesBOOTPMessageType(msg.BOOTPMessageType) {
		mainErr.Add("bootp message type does not match message type option")
	}

	if mainErr.HasReasons() {
		return mainErr
	}

	return nil
}
