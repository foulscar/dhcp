package dhcp

import (
	"errors"
	"net"
)

var RequiredParamsDISCOVER = []OptionCode{
	OptionCodeSubnetMask,
	OptionCodeRouter,
	OptionCodeDNS,
	OptionCodeDomainName,
	OptionCodeBroadcastAddr,
}

func NewDISCOVER(
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
	for _, requiredOptC := range RequiredParamsDISCOVER {
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
