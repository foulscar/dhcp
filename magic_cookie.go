package dhcp

import "slices"

// MagicCookie is a 4-byte slice used for encoding DHCP Messages.
// All DHCP Messages must contain this to be considered a DHCP Message
var MagicCookie []byte = []byte{0x63, 0x82, 0x53, 0x63}

// HasMagicCookie checks if an encoded, potential DHCP Message contains the DHCP Magic Cookie.
// Useful for checking for DHCP Messages before testing validity
func HasMagicCookie(data []byte) bool {
	if len(data) < 240 {
		return false
	}
	if slices.Equal(data[236:240], MagicCookie) {
		return true
	}

	return false
}
