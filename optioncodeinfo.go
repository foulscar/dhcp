package dhcp

type OptionCodeInfo struct {
	String      string
	DataBuilder func([]byte) (OptionData, error)
}

const (
	OptionCodePad                 OptionCode = OptionCode(0)
	OptionCodeSubnetMask          OptionCode = OptionCode(1)
	OptionCodeRouter              OptionCode = OptionCode(3)
	OptionCodeTimeServer          OptionCode = OptionCode(4)
	OptionCodeNameServer          OptionCode = OptionCode(5)
	OptionCodeDNS                 OptionCode = OptionCode(6)
	OptionCodeLogServer           OptionCode = OptionCode(7)
	OptionCodeQuoteOfTheDayServer OptionCode = OptionCode(8)
	OptionCodeLRPServer           OptionCode = OptionCode(9)
	OptionCodeHostname            OptionCode = OptionCode(12)
	OptionCodeDomainname          OptionCode = OptionCode(15)
	OptionCodeRootPath            OptionCode = OptionCode(17)
	OptionCodeMessageType         OptionCode = OptionCode(53)

	OptionCodeEnd OptionCode = OptionCode(255)
)

var OptionCodeToInfo map[OptionCode]OptionCodeInfo = map[OptionCode]OptionCodeInfo{
	OptionCodeSubnetMask: OptionCodeInfo{
		String:      "Subnet Mask",
		DataBuilder: BuildOptionDataSubnetMask,
	},
	OptionCodeMessageType: OptionCodeInfo{
		String:      "DHCP Message Type",
		DataBuilder: BuildOptionDataMessageType,
	},
}
