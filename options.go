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

func (opts Options) Unmarshal() []byte {
	var data []byte
	for _, opt := range opts {
		data = append(data, opt.Unmarshal()...)
	}
	data = append(data, byte(OptionCodeEnd))

	return data
}

func MarshalOptions(data []byte) (Options, []error) {
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

		optCodeInfo, exists := OptionCodeToInfo[optCode]
		if !exists {
			errs = append(errs, fmt.Errorf("OptionCode '(optc %d)' not recognized", int(optCode)))
			i += 1 + optLen
			continue
		}

		optData, err := optCodeInfo.DataMarshaler(data[i+2 : i+2+optLen])
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
