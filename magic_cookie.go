package dhcp

import "slices"

var MagicCookie []byte = []byte{0x63, 0x82, 0x53, 0x63}

func HasMagicCookie(data []byte) bool {
	if slices.Equal(data[236:240], MagicCookie) {
		return true
	}
	return false
}
