package dhcp

var MagicCookie []byte = []byte{0x63, 0x82, 0x53, 0x63}

type OptionCode byte

type OptionField struct {
	OptionCode OptionCode
	Len        uint8
	Val        []byte
}

type Options []OptionField
