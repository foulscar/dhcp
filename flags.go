package dhcp

// Flags represents DHCP Message Flags for a Message.
// This package currently only contains presets for Unicast and Broadcast
type Flags uint16

// Available Flags presets this package offers
const (
	FlagsUnicast   Flags = 0x0
	FlagsBroadcast Flags = 0x8000
)

var flagsToString = map[Flags]string{
	FlagsUnicast:   "Unicast",
	FlagsBroadcast: "Broadcast",
}

// String returns the addressing methods preferred from a client/server.
// "Unicast" or "Broadcast"
func (f Flags) String() string {
	return flagsToString[f]
}
