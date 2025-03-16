package dhcp

import "errors"

type OptionDataMessageType struct {
	Type OptionMessageTypeCode
}

type OptionMessageTypeCode uint8

const (
	OptionMessageTypeCodeDISCOVER OptionMessageTypeCode = 1
	OptionMessageTypeCodeOFFER    OptionMessageTypeCode = 2
	OptionMessageTypeCodeREQUEST  OptionMessageTypeCode = 3
	OptionMessageTypeCodeDECLINE  OptionMessageTypeCode = 4
	OptionMessageTypeCodeACK      OptionMessageTypeCode = 5
	OptionMessageTypeCodeNACK     OptionMessageTypeCode = 6
	OptionMessageTypeCodeRELEASE  OptionMessageTypeCode = 7
	OptionMessageTypeCodeINFORM   OptionMessageTypeCode = 8
)

var OptionMessageTypeCodeToString = map[OptionMessageTypeCode]string{
	OptionMessageTypeCodeDISCOVER: "DISCOVER",
	OptionMessageTypeCodeOFFER:    "OFFER",
	OptionMessageTypeCodeREQUEST:  "REQUEST",
	OptionMessageTypeCodeDECLINE:  "DECLINE",
	OptionMessageTypeCodeACK:      "ACK",
	OptionMessageTypeCodeNACK:     "NACK",
	OptionMessageTypeCodeRELEASE:  "RELEASE",
	OptionMessageTypeCodeINFORM:   "INFORM",
}

var OptionMessageTypeCodeValidBOOTP = map[OptionMessageTypeCode]BOOTPMessageType{
	OptionMessageTypeCodeDISCOVER: BOOTPMessageTypeRequest,
	OptionMessageTypeCodeOFFER:    BOOTPMessageTypeReply,
	OptionMessageTypeCodeREQUEST:  BOOTPMessageTypeRequest,
	OptionMessageTypeCodeDECLINE:  BOOTPMessageTypeRequest,
	OptionMessageTypeCodeACK:      BOOTPMessageTypeReply,
	OptionMessageTypeCodeNACK:     BOOTPMessageTypeReply,
	OptionMessageTypeCodeRELEASE:  BOOTPMessageTypeRequest,
	OptionMessageTypeCodeINFORM:   BOOTPMessageTypeRequest,
}

func (msgType OptionMessageTypeCode) MatchesBOOTPMessageType(bootpMsgType BOOTPMessageType) bool {
	return OptionMessageTypeCodeValidBOOTP[msgType] == bootpMsgType
}

func (optD OptionDataMessageType) String() string {
	return OptionMessageTypeCodeToString[optD.Type]
}

func (optD OptionDataMessageType) IsValid() bool {
	if optD.String() == "" {
		return false
	}
	return true
}

func (optD OptionDataMessageType) Marshal() ([]byte, error) {
	if !optD.IsValid() {
		return nil, errors.New("option data contains an invalid message type")
	}
	return []byte{byte(optD.Type)}, nil
}

func UnmarshalOptionDataMessageType(data []byte) (OptionData, error) {
	msgType := OptionMessageTypeCode(data[0])
	_, exists := OptionMessageTypeCodeToString[msgType]
	if !exists || len(data) > 1 {
		return nil, errors.New("invalid message type")
	}

	return OptionDataMessageType{Type: msgType}, nil
}

func NewOptionMessageType(msgType OptionMessageTypeCode) (Option, error) {
	_, exists := OptionMessageTypeCodeToString[msgType]
	if !exists {
		return Option{}, errors.New("invalid message type")
	}

	return Option{
		Code: OptionCodeMessageType,
		Data: OptionDataMessageType{Type: msgType},
	}, nil
}
