package dhcp

type OptionCode uint8

type OptionDataMarshaller func([]byte) (OptionData, error)

func (code OptionCode) String() string {
	return OptionCodeToString[code]
}

const (
	OptionCodePad OptionCode = OptionCode(0)

	OptionCodeSubnetMask           OptionCode = OptionCode(1)
	OptionCodeRouter               OptionCode = OptionCode(3)
	OptionCodeTimeServer           OptionCode = OptionCode(4)
	OptionCodeNameServer           OptionCode = OptionCode(5)
	OptionCodeDNS                  OptionCode = OptionCode(6)
	OptionCodeLogServer            OptionCode = OptionCode(7)
	OptionCodeQuoteOfTheDayServer  OptionCode = OptionCode(8)
	OptionCodeLRPServer            OptionCode = OptionCode(9)
	OptionCodeHostname             OptionCode = OptionCode(12)
	OptionCodeDomainname           OptionCode = OptionCode(15)
	OptionCodeRootPath             OptionCode = OptionCode(17)
	OptionCodeMessageType          OptionCode = OptionCode(53)
	OptionCodeParameterRequestList OptionCode = OptionCode(55)

	OptionCodeEnd OptionCode = OptionCode(255)
)

var OptionCodeToString = map[OptionCode]string{
	OptionCodeSubnetMask:           "Subnet Mask",
	OptionCodeMessageType:          "DHCP Message Type",
	OptionCodeParameterRequestList: "Parameter Request List",
}

var OptionCodeToDataMarshaller = map[OptionCode]OptionDataMarshaller{
	OptionCodeSubnetMask:           MarshalOptionDataSubnetMask,
	OptionCodeMessageType:          MarshalOptionDataMessageType,
	OptionCodeParameterRequestList: MarshalOptionDataParameterRequestList,
}
