package dhcp

import (
	"strconv"
	"strings"
)

// Options represents a DHCP Message's Options
type Options map[OptionCode]Option

// String returns a verbose, multi-line string of all the Option Entries
func (opts Options) String() string {
	var sb strings.Builder

	i := 0
	for _, opt := range opts {
		if i != 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(opt.String())
		i++
	}
	if i != 0 {
		sb.WriteString("\n")
	}
	sb.WriteString("END [255]")

	return sb.String()
}

// Add updates/creates an entry in opts for each Option given
func (opts Options) Add(newOpts ...*Option) {
	for _, opt := range newOpts {
		if opt == nil {
			continue
		}
		opts[opt.Code] = *opt
	}
}

// Remove removes the Option associated with optCode from opts
func (opts Options) Remove(optCode OptionCode) {
	delete(opts, optCode)
}

// Contains returns true if there is an Option entry associated with code in opts
func (opts Options) Contains(code OptionCode) bool {
	_, exists := opts[code]
	return exists
}

// IsValid returns true if all Option Entries in opts are valid
func (opts Options) IsValid() *ErrorExt {
	mainErr := NewErrorExt("Options with " + strconv.Itoa(len(opts)) + " options is invalid")

	for _, opt := range opts {
		err := opt.IsValid()
		if err == nil {
			continue
		}
		mainErr.Add(err)
	}

	if !mainErr.HasReasons() {
		return nil
	}

	return mainErr
}

// GetDefaults returns a slice of all OptionCodes present in opts that are using default data handling.
// OptionDataDefault or a user-defined OptionData for defaults
func (opts Options) GetDefaults() []OptionCode {
	var defCodes []OptionCode
	for code, opt := range opts {
		if opt.IsDefault {
			defCodes = append(defCodes, code)
		}
	}

	return defCodes
}

// GetNonDefaults returns a slice of all OptionCodes present in opts that are NOT using default data handling
func (opts Options) GetNonDefaults() []OptionCode {
	var codes []OptionCode
	for code, opt := range opts {
		if !opt.IsDefault {
			codes = append(codes, code)
		}
	}

	return codes
}

// Marshal encodes opts as an Options field for a DHCP Message
func (opts Options) Marshal() ([]byte, *ErrorExt) {
	mainErr := NewErrorExt("could not marshal Options")

	var data []byte
	for _, opt := range opts {
		optErr := opt.IsValid()
		if optErr != nil {
			mainErr.Add(optErr)
			continue
		}

		optData, optDataErr := opt.Data.Marshal()
		if optDataErr != nil {
			mainErr.Add(NewErrorExt("Option with OptionCode '"+strconv.Itoa(int(opt.Code))+"' has invalid data", optDataErr))
			continue
		}

		if mainErr.HasReasons() {
			continue
		}

		data = append(data, byte(opt.Code))
		data = append(data, byte(len(optData)))
		data = append(data, optData...)
	}

	if mainErr.HasReasons() {
		return nil, mainErr
	}

	data = append(data, byte(OptionCodeEnd))

	return data, nil
}

// UnmarshalOptions parses data as an Options field from a DHCP Message
func UnmarshalOptions(data []byte) (Options, *ErrorExt) {
	mainErr := NewErrorExt("data contained invalid Options")
	opts := make(Options)

	i := 0
	for i < len(data)-1 {
		optCode := OptionCode(data[i])
		if optCode == OptionCodeEnd {
			break
		}
		if optCode == OptionCodePad {
			i++
			continue
		}

		optLen := int(data[i+1])
		if len(data) < i+2+optLen {
			break
		}

		optDataUnmarshaller, isDefault := optMap.GetDataUnmarshaller(optCode)
		optData, err := optDataUnmarshaller(data[i+2 : i+2+optLen])
		if err != nil {
			mainErr.Add(err)
			i += 2 + optLen
			continue
		}

		opt := Option{
			Code:      optCode,
			Data:      optData,
			IsDefault: isDefault,
		}

		opts[optCode] = opt
		i += 2 + optLen
	}

	if !mainErr.HasReasons() {
		return opts, nil
	}

	return opts, mainErr
}
