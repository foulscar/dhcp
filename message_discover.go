package dhcp

import (
	"errors"
	"net"
)

// RequiredParamsDISCOVER contains a list of the minimum OptionCodes that must be specified
// in a Message's Parameter Request List Option for it to be valid as a DISCOVER Message
var requiredParamsDISCOVER = []OptionCode{
	OptionCodeSubnetMask,
	OptionCodeRouter,
	OptionCodeDNS,
	OptionCodeDomainName,
	OptionCodeBroadcastAddr,
}

// NewDISCOVER returns a Message, ready for use as a DISCOVER Message
func newDISCOVER(
	clientMacAddr net.HardwareAddr,
	paramsRequested ...OptionCode,
) (*Message, error) {
	if len(clientMacAddr) != 6 {
		return nil, errors.New("invalid client mac address")
	}

	msg := newEmptyMessage()
	msg.BOOTPMessageType = BOOTPMessageTypeRequest
	msg.SetMessageType(OptionMessageTypeCodeDISCOVER)

	params := make([]OptionCode, len(paramsRequested))
	copy(params, paramsRequested)
	for _, requiredOptC := range requiredParamsDISCOVER {
		contains := false
		for _, optC := range params {
			if optC == requiredOptC {
				contains = true
				break
			}
		}
		if !contains {
			params = append(params, requiredOptC)
		}
	}

	opt := NewOptionParameterRequestList(params...)
	if !opt.IsValid() {
		return nil, errors.New("invalid parameter request list")
	}

	return &msg, nil
}
