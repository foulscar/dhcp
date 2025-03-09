package dhcp

import (
	"errors"
)

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

func (msg Message) SetMessageType(msgType OptionMessageTypeCode) error {
	_, exists := OptionMessageTypeCodeToString[msgType]
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

func (msg Message) AddOption(opt Option) {
	msg.Options.Add(opt)
}

func (msg Message) UpdateOption(opt Option) {
	msg.Options.Update(opt)
}

func (msg Message) RemoveOption(optCode OptionCode) {
	msg.Options.Remove(optCode)
}
