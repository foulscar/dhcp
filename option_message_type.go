package dhcp

import "errors"

// OptionDataMessageType stores the DHCP Message Type of a Message
type OptionDataMessageType struct {
	Type OptionMessageTypeCode
}

// OptionMessageTypeCode represents a DHCP Message Type
type OptionMessageTypeCode uint8

// All recognized DHCP Message Type Codes
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

var optionMessageTypeCodeToString = map[OptionMessageTypeCode]string{
	OptionMessageTypeCodeDISCOVER: "DISCOVER",
	OptionMessageTypeCodeOFFER:    "OFFER",
	OptionMessageTypeCodeREQUEST:  "REQUEST",
	OptionMessageTypeCodeDECLINE:  "DECLINE",
	OptionMessageTypeCodeACK:      "ACK",
	OptionMessageTypeCodeNACK:     "NACK",
	OptionMessageTypeCodeRELEASE:  "RELEASE",
	OptionMessageTypeCodeINFORM:   "INFORM",
}

var optionMessageTypeCodeValidBOOTP = map[OptionMessageTypeCode]BOOTPMessageType{
	OptionMessageTypeCodeDISCOVER: BOOTPMessageTypeRequest,
	OptionMessageTypeCodeOFFER:    BOOTPMessageTypeReply,
	OptionMessageTypeCodeREQUEST:  BOOTPMessageTypeRequest,
	OptionMessageTypeCodeDECLINE:  BOOTPMessageTypeRequest,
	OptionMessageTypeCodeACK:      BOOTPMessageTypeReply,
	OptionMessageTypeCodeNACK:     BOOTPMessageTypeReply,
	OptionMessageTypeCodeRELEASE:  BOOTPMessageTypeRequest,
	OptionMessageTypeCodeINFORM:   BOOTPMessageTypeRequest,
}

// MatchesBOOTPMessageType returns true if msgType can be valid under the context of having
// bootpMsgType as the BOOTP Message Type of it's parent Message.
// Every OptionMessageTypeCode will be valid under either 'REQUEST' or 'REPLY'
func (msgType OptionMessageTypeCode) MatchesBOOTPMessageType(bootpMsgType BOOTPMessageType) bool {
	return optionMessageTypeCodeValidBOOTP[msgType] == bootpMsgType
}

// String returns a human-readable version of optD.Type
func (optD OptionDataMessageType) String() string {
	return optionMessageTypeCodeToString[optD.Type]
}

// IsValid checks if the underlying optD.Type exists as a recognized DHCP Message Type
func (optD OptionDataMessageType) IsValid() bool {
	if optD.String() == "" {
		return false
	}
	return true
}

// Marshal encodes optD as the value for a DHCP Message Type Option
func (optD OptionDataMessageType) Marshal() ([]byte, error) {
	if !optD.IsValid() {
		return nil, errors.New("option data contains an invalid message type")
	}
	return []byte{byte(optD.Type)}, nil
}

// UnmarshalOptionDataMessageType parses data as the value for a DHCP Message Type Option.
// Will return an error if len(data) > 1 or if data contains an unrecognized DHCP Message Type Code
func UnmarshalOptionDataMessageType(data []byte) (OptionData, error) {
	msgType := OptionMessageTypeCode(data[0])
	_, exists := optionMessageTypeCodeToString[msgType]
	if !exists || len(data) > 1 {
		return nil, errors.New("invalid message type")
	}

	return OptionDataMessageType{Type: msgType}, nil
}

// NewOptionMessageType is a helper function for constructing a DHCP Message Type Option.
// It will hold OptionDataMessageType as the Option's Data
func NewOptionMessageType(msgType OptionMessageTypeCode) (*Option, error) {
	_, exists := optionMessageTypeCodeToString[msgType]
	if !exists {
		return nil, errors.New("invalid message type")
	}

	return &Option{
		Code:      OptionCodeMessageType,
		Data:      OptionDataMessageType{Type: msgType},
		IsDefault: false,
	}, nil
}
