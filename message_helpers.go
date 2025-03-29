package dhcp

import (
	"errors"
)

// GetMessageType returns the DHCP Message Type of the Message if the Option exists and is valid
func (msg Message) GetMessageType() (OptionMessageTypeCode, error) {
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

// SetMessageType changes or creates an Option for the specified DHCP Message Type
func (msg Message) SetMessageType(msgType OptionMessageTypeCode) error {
	_, exists := optionMessageTypeCodeToString[msgType]
	if !exists {
		return errors.New("invalid message type")
	}

	opt, err := NewOptionMessageType(msgType)
	if err != nil {
		return err
	}
	msg.Options.Add(opt)

	return nil
}

// AddOptions updates/adds the Options to msg.Options
func (msg Message) AddOptions(opts ...*Option) {
	msg.Options.Add(opts...)
}

// RemoveOption removes your OptionCode's entry from msg.Options
func (msg Message) RemoveOption(optCode OptionCode) {
	msg.Options.Remove(optCode)
}
