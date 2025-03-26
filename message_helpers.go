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
	msg.Options.Update(opt)

	return nil
}

// AddOption adds the Option to msg.Options
func (msg Message) AddOption(opt Option) {
	msg.Options.Add(opt)
}

// UpdateOption swaps the existing Option entry associated with your Option's OptionCode with your Option.
// It will create an entry in msg.Options if one doesn't already exist
func (msg Message) UpdateOption(opt Option) {
	msg.Options.Update(opt)
}

// RemoveOption removes your OptionCode's entry from msg.Options
func (msg Message) RemoveOption(optCode OptionCode) {
	msg.Options.Remove(optCode)
}
