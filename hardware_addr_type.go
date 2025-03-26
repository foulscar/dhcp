package dhcp

// HardwareAddrType represents the Hardware Address Type specified in a DHCP Message.
// This package currently only contains Ethernet
type HardwareAddrType uint8

const HardwareAddrTypeEthernet = 1

var hardwareAddrTypeToString = map[HardwareAddrType]string{
	HardwareAddrTypeEthernet: "Ethernet",
}

var hardwareAddrLengths = map[HardwareAddrType]uint8{
	HardwareAddrTypeEthernet: 6,
}

// String returns a human-readable version of HardwareAddrType
func (hw HardwareAddrType) String() string {
	return hardwareAddrTypeToString[hw]
}

// ValidLength returns the valid length a hardware address of HardwareAddrType type can be.
// This can also be useful for setting Message.HardwareAddrLen.
// Ex: Ethernet (mac addresses) have a valid length of 6 bytes
func (hw HardwareAddrType) ValidLength() uint8 {
	return hardwareAddrLengths[hw]
}
