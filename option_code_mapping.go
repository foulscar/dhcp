package dhcp

type OptionCodeMapping struct {
	ToString           map[OptionCode]string
	ToDataUnmarshaller map[OptionCode]OptionDataUnmarshaller
}

var GlobalOptionCodeMapping = OptionCodeMapping{
	ToString:           OptionCodeToString,
	ToDataUnmarshaller: OptionCodeToDataUnmarshaller,
}

// GetString returns the human-readable name represented by the OptionCode.
// It will fetch this value from optCodeMap
func (optCodeMap OptionCodeMapping) GetString(code OptionCode) string {
        return optCodeMap.ToString[code]
}

// GetDataUnmarshaller returns the the OptionDataUnmarshaller associated with the OptionCode.
// It will fetch this value from optCodeMap
func (optCodeMap OptionCodeMapping) GetDataUnmarshaller(code OptionCode) OptionDataUnmarshaller {
	if code == OptionCodePad || code == OptionCodeEnd {
		return nil
	}
	if optCodeMap.ToDataUnmarshaller[code] == nil {
		return UnmarshalOptionDefault
	}
	return optCodeMap.ToDataUnmarshaller[code]
}

var optMap = &GlobalOptionCodeMapping
