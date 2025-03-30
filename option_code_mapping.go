package dhcp

import "reflect"

// OptionCodeMapping holds maps used for parsing/handling Options.
// You can create your own if you would like to change the behavior of this package,
// see GlobalOptionCodeMapping
type OptionCodeMapping struct {
	ToString                map[OptionCode]string
	ToDataType              map[OptionCode]reflect.Type
	ToDataUnmarshaller      map[OptionCode]OptionDataUnmarshaller
	DefaultDataType         reflect.Type
	DefaultDataUnmarshaller OptionDataUnmarshaller
}

// GlobalOptionCodeMapping is used as the relevant OptionCodeMapping for all functions in this package.
// You can change it's values at the start of your program for changing the behavior of this package.
// Changing it's values whilst parsing/handling Messages, Options, etc. can be unpredictable and it is
// recommended that you handle this before your program's main execution.
// Leaving this unchanged will result in default behavior
var GlobalOptionCodeMapping = OptionCodeMapping{
	ToString:                OptionCodeToString,
	ToDataType:              OptionCodeToDataType,
	ToDataUnmarshaller:      OptionCodeToDataUnmarshaller,
	DefaultDataType:         reflect.TypeOf(OptionDataDefault{}),
	DefaultDataUnmarshaller: UnmarshalOptionDefault,
}

// GetString returns the human-readable name represented by the OptionCode.
// It will fetch this value from optCodeMap
func (optCodeMap OptionCodeMapping) GetString(code OptionCode) string {
	return optCodeMap.ToString[code]
}

// GetDataType returns the the OptionData implementation associated with the OptionCode and
// a boolean that is true if returning the default OptionData implementation. If the OptionCode does
// not have a mapping, the default OptionData implementation will be returned.
// It will fetch this value from optCodeMap
func (optCodeMap OptionCodeMapping) GetDataType(code OptionCode) (dataType reflect.Type, isDefault bool) {
	if code == OptionCodePad || code == OptionCodeEnd {
		return nil, false
	}
	if optCodeMap.ToDataType[code] == nil {
		return optCodeMap.DefaultDataType, true
	}
	return optCodeMap.ToDataType[code], false
}

// GetDataUnmarshaller returns the the OptionDataUnmarshaller associated with the OptionCode and
// a boolean that is true if returning the default OptionDataUnmarshaller. If the OptionCode does
// not have a mapping, the default OptionDataUnmarshaller will be returned.
// It will fetch this value from optCodeMap
func (optCodeMap OptionCodeMapping) GetDataUnmarshaller(code OptionCode) (optDUnmarshaller OptionDataUnmarshaller, isDefault bool) {
	if code == OptionCodePad || code == OptionCodeEnd {
		return nil, false
	}
	if optCodeMap.ToDataUnmarshaller[code] == nil {
		return optCodeMap.DefaultDataUnmarshaller, true
	}
	return optCodeMap.ToDataUnmarshaller[code], false
}

var optMap = &GlobalOptionCodeMapping
