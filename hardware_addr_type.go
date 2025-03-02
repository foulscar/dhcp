package dhcp

type HardwareAddrType uint8

const HardwareAddrTypeEthernet = 1

var HardwareAddrTypeToString map[HardwareAddrType]string = map[HardwareAddrType]string{
	HardwareAddrTypeEthernet: "Ethernet",
}

var HardwareAddrLengths = map[HardwareAddrType]uint8{
	HardwareAddrTypeEthernet: 6,
}

func (hw HardwareAddrType) ValidLength() uint8 {
	return HardwareAddrLengths[hw]
}
