package dhcp

import (
	"errors"
	"strings"
)

type OptionDataParameterRequestList struct {
	List []OptionCode
}

func (optD OptionDataParameterRequestList) String() string {
	var sb strings.Builder

	sb.WriteString("[")
	for i, optC := range optD.List {
		if i != 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(OptionCodeToString[optC])
	}
	sb.WriteString("]")

	return sb.String()
}

func (optD OptionDataParameterRequestList) IsValid() bool {
	for _, code := range optD.List {
		if code == OptionCodePad || code == OptionCodeEnd {
			return false
		}
	}
	return true
}

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

func (optD OptionDataParameterRequestList) Add(optC OptionCode) {
	optD.List = append(optD.List, optC)
}

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
