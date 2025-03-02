package dhcp

type BOOTPMessageType uint8

const (
	BOOTPMessageTypeRequest BOOTPMessageType = 1
	BOOTPMessageTypeReply   BOOTPMessageType = 2
)

var BOOTPMessageTypeToString map[BOOTPMessageType]string = map[BOOTPMessageType]string{
	BOOTPMessageTypeRequest: "REQUEST",
	BOOTPMessageTypeReply:   "REPLY",
}
