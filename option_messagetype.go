package dhcp

import "errors"

type OptionDataMessageType struct {
	Type OptionMessageTypeCode
}

type OptionMessageTypeCode uint8

const (
	OptionMessageTypeCodeDISCOVER OptionMessageTypeCode = OptionMessageTypeCode(1)
	OptionMessageTypeCodeOFFER    OptionMessageTypeCode = OptionMessageTypeCode(2)
	OptionMessageTypeCodeREQUEST  OptionMessageTypeCode = OptionMessageTypeCode(3)
	OptionMessageTypeCodeDECLINE  OptionMessageTypeCode = OptionMessageTypeCode(4)
	OptionMessageTypeCodeACK      OptionMessageTypeCode = OptionMessageTypeCode(5)
	OptionMessageTypeCodeNACK     OptionMessageTypeCode = OptionMessageTypeCode(6)
	OptionMessageTypeCodeRELEASE  OptionMessageTypeCode = OptionMessageTypeCode(7)
	OptionMessageTypeCodeINFORM   OptionMessageTypeCode = OptionMessageTypeCode(8)
)

var OptionMessageTypeCodeToString map[OptionMessageTypeCode]string = map[OptionMessageTypeCode]string{
	OptionMessageTypeCodeDISCOVER: "DISCOVER",
	OptionMessageTypeCodeOFFER:    "OFFER",
	OptionMessageTypeCodeREQUEST:  "REQUEST",
	OptionMessageTypeCodeDECLINE:  "DECLINE",
	OptionMessageTypeCodeACK:      "ACK",
	OptionMessageTypeCodeNACK:     "NACK",
	OptionMessageTypeCodeRELEASE:  "RELEASE",
	OptionMessageTypeCodeINFORM:   "INFORM",
}

func (optD OptionDataMessageType) Raw() []byte {
	return []byte{byte(optD.Type)}
}

func BuildOptionDataMessageType(data []byte) (OptionData, error) {
	msgType := OptionMessageTypeCode(data[0])
	_, exists := OptionMessageTypeCodeToString[msgType]
	if !exists || len(data) > 1 {
		return nil, errors.New("invalid message type")
	}

	return OptionDataMessageType{Type: msgType}, nil
}
