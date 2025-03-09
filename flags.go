package dhcp

type Flags uint16

const (
	FlagsUnicast   Flags = 0x0
	FlagsBroadcast Flags = 0x8000
)

var FlagsToString = map[Flags]string{
	FlagsUnicast:   "Unicast",
	FlagsBroadcast: "Broadcast",
}

func (f Flags) String() string {
	return FlagsToString[f]
}
