package dhcp

// BOOTPMessageType represents the BOOTP Message Type of a message.
// REQUEST or REPLY
type BOOTPMessageType uint8

const (
	BOOTPMessageTypeRequest BOOTPMessageType = 1
	BOOTPMessageTypeReply   BOOTPMessageType = 2
)

var bootpMessageTypeToString = map[BOOTPMessageType]string{
	BOOTPMessageTypeRequest: "REQUEST",
	BOOTPMessageTypeReply:   "REPLY",
}

// String returns the name of the BOOTP Message Type associated with msgType.
// "REQUEST" or "REPLY"
func (msgType BOOTPMessageType) String() string {
        return bootpMessageTypeToString[msgType]
}
