package dhcp

import (
	"errors"
	"strings"
)

// OptionDataParameterRequestList represents data for the DHCP Parameter Request List Option.
// It holds a list of parameters (OptionCodes) requested by a client.
type OptionDataParameterRequestList struct {
	List []OptionCode
}

// String returns "[options_here]" where options_here represents a comma-seperated list of all
// the options the parent Message is requesting as human-readable names.
// Ex: "[Subnet Mask, Hostname]"
func (optD OptionDataParameterRequestList) String() string {
	var sb strings.Builder

	sb.WriteString("[")
	for i, optC := range optD.List {
		if i != 0 {
			sb.WriteString(", ")
		}
		name := optC.String()
		if name == "" {
			name = "UNKNOWN"
		}
		sb.WriteString(name)
	}
	sb.WriteString("]")

	return sb.String()
}

// IsValid returns true as long as no requested OptionCodes are OptionCodePad or OptionCodeEnd
func (optD OptionDataParameterRequestList) IsValid() error {
	for _, code := range optD.List {
		if code == OptionCodePad || code == OptionCodeEnd {
			return errors.New("Paramter Request List contains Pad or End OptionCodes")
		}
	}
	return nil
}

// Marshal encodes optD as the value for a DHCP Parameters Request List Option
func (optD OptionDataParameterRequestList) Marshal() ([]byte, error) {
	if err := optD.IsValid(); err != nil {
		return nil, err
	}
	data := make([]byte, len(optD.List))
	for i, optC := range optD.List {
		data[i] = byte(optC)
	}

	return data, nil
}

// Add adds the given OptionCodes to optD.List
func (optD OptionDataParameterRequestList) Add(optCodes ...OptionCode) {
	optD.List = append(optD.List, optCodes...)
}

// UnmarshalOptionDataParameterRequestList parses data as the value for a DHCP Parameter Request List Option
func UnmarshalOptionDataParameterRequestList(data []byte) (OptionData, error) {
	list := make([]OptionCode, len(data))
	for i, b := range data {
		list[i] = OptionCode(b)
	}
	optData := OptionDataParameterRequestList{List: list}
	if err := optData.IsValid(); err != nil {
		return nil, err
	}

	return optData, nil
}

// NewOptionParameterRequestList is a helper function for constructing a Parameter Request List Option.
// It will hold OptionDataParameterRequestList as the Option's Data
func NewOptionParameterRequestList(optCodes ...OptionCode) *Option {
	opt := Option{
		Code: OptionCodeParameterRequestList,
		Data: OptionDataParameterRequestList{
			List: make([]OptionCode, 0),
		},
		IsDefault: false,
	}

	for _, optCode := range optCodes {
		opt.Data.(OptionDataParameterRequestList).Add(optCode)
	}

	return &opt
}
