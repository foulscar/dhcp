package dhcp

// OptionCodeMapping holds maps used for parsing/handling Options.
// You can create your own if you would like to change the behavior of this package,
// see GlobalOptionCodeMapping
type OptionCodeMapping struct {
	ToString           map[OptionCode]string
	ToDataUnmarshaller map[OptionCode]OptionDataUnmarshaller
}

// GlobalOptionCodeMapping is used as the relevant OptionCodeMapping for all functions in this package.
// You can change it's values at the start of your program for changing the behavior of this package.
// Changing it's values whilst parsing/handling Messages, Options, etc. can be unpredictable and it is
// recommended that you handle this before your program's main execution.
// Leaving this unchanged will result in default behavior
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
