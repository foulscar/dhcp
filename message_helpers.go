package dhcp

import (
	"errors"
)

func (msg Message) GetDHCPMessageType() (OptionMessageTypeCode, error) {
	opt, exists := msg.Options[OptionCodeMessageType]
	if !exists {
		return OptionMessageTypeCode(0), errors.New("message does not contain the message type option")
	}

	optData := opt.Data.(OptionDataMessageType)
	if optData.String() == "" {
		return OptionMessageTypeCode(0), errors.New("message contains an invalid message type")
	}

	return optData.Type, nil
}
