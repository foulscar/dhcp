package dhcp

import (
	"encoding/binary"
)

func (msg *Message) Unmarshal() []byte {
	data := make([]byte, 240)

	data[0] = byte(msg.BOOTPMessageType)
	data[1] = byte(msg.HardwareAddrType)
	data[2] = byte(msg.HardwareAddrLen)
	data[3] = byte(msg.HopCount)
	binary.BigEndian.PutUint32(data[4:8], msg.TransactionID)
	binary.BigEndian.PutUint16(data[8:10], msg.SecsElapsed)
	binary.BigEndian.PutUint16(data[10:12], msg.Flags)
	copy(data[12:16], []byte(msg.ClientIPAddr))
	copy(data[16:20], []byte(msg.YourIPAddr))
	copy(data[20:24], []byte(msg.ServerIPAddr))
	copy(data[24:28], []byte(msg.GatewayIPAddr))
	copy(data[28:44], []byte(msg.ClientHardwareAddr))
	copy(data[44:108], msg.ServerHostname)
	copy(data[108:236], msg.BootFilename)
	copy(data[236:240], MagicCookie)
	data = append(data, msg.Options.Unmarshal()...)

	paddingLen := 0
	if len(data) < 300 {
		paddingLen = 300 - len(data)
	}

	return append(data, make([]byte, paddingLen)...)
}
