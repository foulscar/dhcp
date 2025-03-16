package dhcp

type OptionCodeMapping struct {
	ToString           map[OptionCode]string
	ToDataUnmarshaller map[OptionCode]OptionDataUnmarshaller
}

var GlobalOptionCodeMapping = OptionCodeMapping{
	ToString:           OptionCodeToString,
	ToDataUnmarshaller: OptionCodeToDataUnmarshaller,
}

var optMap = &GlobalOptionCodeMapping
