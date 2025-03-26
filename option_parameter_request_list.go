package dhcp

import (
	"errors"
	"strings"
)

// OptionDataParameterRequestList represents data for the Parameter Request List Option
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
		sb.WriteString(optC.String())
	}
	sb.WriteString("]")

	return sb.String()
}

// IsValid returns true as long as no requested OptionCodes are OptionCodePad or OptionCodeEnd
func (optD OptionDataParameterRequestList) IsValid() bool {
	for _, code := range optD.List {
		if code == OptionCodePad || code == OptionCodeEnd {
			return false
		}
	}
	return true
}

// Marshal encodes optD as the value for a Parameters Request List Option 
func (optD OptionDataParameterRequestList) Marshal() ([]byte, error) {
	if !optD.IsValid() {
		return nil, errors.New("option data is invalid")
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

// UnmarshalOptionDataParameterRequestList parses data as the value for a Parameter Request List Option
func UnmarshalOptionDataParameterRequestList(data []byte) (OptionData, error) {
	list := make([]OptionCode, len(data))
	for i, b := range data {
		list[i] = OptionCode(b)
	}
	optData := OptionDataParameterRequestList{List: list}
	if !optData.IsValid() {
		return nil, errors.New("data contains an invalid option code")
	}

	return optData, nil
}

// NewOptionParameterRequestList is a helper function for constructing a Parameter Request List Option.
// It will hold OptionDataParameterRequestList as the Option's Data
func NewOptionParameterRequestList(optCodes ...OptionCode) Option {
	opt := Option{
		Code: OptionCodeParameterRequestList,
		Data: OptionDataParameterRequestList{
			List: make([]OptionCode, 0),
		},
	}

	for _, optCode := range optCodes {
		opt.Data.(OptionDataParameterRequestList).Add(optCode)
	}

	return opt
}
