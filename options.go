package dhcp

import (
	"fmt"
	"strings"
)

type Options map[OptionCode]Option

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

func (opts Options) Add(opt Option) {
	opts[opt.Code] = opt
}

func (opts Options) Update(opt Option) {
	opts[opt.Code] = opt
}

func (opts Options) Remove(optCode OptionCode) {
	delete(opts, optCode)
}

func (opts Options) Contains(code OptionCode) bool {
	_, exists := opts[code]
	return exists
}

func (opts Options) IsValid() bool {
	for _, opt := range opts {
		if !opt.IsValid() {
			return false
		}
	}

	return true
}

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
		if len(data) <= i+2+optLen {
			break
		}

		optData, err := optMap.GetDataUnmarshaller(optCode)(data[i+2 : i+2+optLen])
		if err != nil {
			errs = append(errs, err)
			i += 1 + optLen
			continue
		}

		opt := Option{
			Code: optCode,
			Data: optData,
		}

		opts[optCode] = opt
		i += 1 + optLen
	}

	return opts, errs
}
