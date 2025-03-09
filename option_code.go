package dhcp

type OptionCode uint8

type OptionDataMarshaller func([]byte) (OptionData, error)

func (code OptionCode) String() string {
	return OptionCodeToString[code]
}

const (
	OptionCodePad OptionCode = 0

	OptionCodeSubnetMask           OptionCode = 1
	OptionCodeRouter               OptionCode = 3
	OptionCodeTimeServer           OptionCode = 4
	OptionCodeNameServer           OptionCode = 5
	OptionCodeDNS                  OptionCode = 6
	OptionCodeLogServer            OptionCode = 7
	OptionCodeQuoteOfTheDayServer  OptionCode = 8
	OptionCodeLRPServer            OptionCode = 9
	OptionCodeHostname             OptionCode = 12
	OptionCodeDomainname           OptionCode = 15
	OptionCodeRootPath             OptionCode = 17
	OptionCodeMessageType          OptionCode = 53
	OptionCodeParameterRequestList OptionCode = 55

	OptionCodeEnd OptionCode = 255
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
