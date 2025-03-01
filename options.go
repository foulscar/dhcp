package dhcp

import (
	"fmt"
)

var MagicCookie []byte = []byte{0x63, 0x82, 0x53, 0x63}

type Options map[OptionCode]Option

type Option struct {
	Code OptionCode
	Data OptionData
}

type OptionCode uint8

type OptionData interface {
	Raw() []byte
}

func (opt Option) String() string {
	return OptionCodeToInfo[opt.Code].String
}

func (opt Option) Unmarshal() []byte {
	data := opt.Data.Raw()
	length := len(data)
	out := make([]byte, length+2)
	out[0] = byte(opt.Code)
	out[1] = byte(length)

	return out
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

		optData, err := optCodeInfo.DataBuilder(data[i+2 : i+2+optLen])
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
