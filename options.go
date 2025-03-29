package dhcp

import (
	"fmt"
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
func (opts Options) IsValid() bool {
	for _, opt := range opts {
		if !opt.IsValid() {
			return false
		}
	}

	return true
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
func (opts Options) Marshal() ([]byte, error) {
	var data []byte
	for _, opt := range opts {
		valid := opt.IsValid()
		optData, err := opt.Data.Marshal()
		if !valid || err != nil {
			return nil, fmt.Errorf("could not marshal options. option with code '%d' is invalid", int(opt.Code))
		}
		data = append(data, optData...)
	}
	data = append(data, byte(OptionCodeEnd))

	return data, nil
}

// UnmarshalOptions parses data as an Options field from a DHCP Message
func UnmarshalOptions(data []byte) (Options, []error) {
	opts := make(Options)
	var errs []error

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
			errs = append(errs, err)
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

	return opts, errs
}
