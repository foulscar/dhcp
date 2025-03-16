package dhcp

type OptionCodeMapping struct {
	ToString           map[OptionCode]string
	ToDataUnmarshaller map[OptionCode]OptionDataUnmarshaller
}

// You can define custom behavior here
var GlobalOptionCodeMapping = OptionCodeMapping{
	ToString:           OptionCodeToString,
	ToDataUnmarshaller: OptionCodeToDataUnmarshaller,
}

func (m OptionCodeMapping) GetString(code OptionCode) string { return m.ToString[code] }

func (m OptionCodeMapping) GetDataUnmarshaller(code OptionCode) OptionDataUnmarshaller {
	if code == OptionCodePad || code == OptionCodeEnd {
		return nil
	}
	if m.ToDataUnmarshaller[code] == nil {
		return UnmarshalOptionDefault
	}
	return m.ToDataUnmarshaller[code]
}

var optMap = &GlobalOptionCodeMapping
