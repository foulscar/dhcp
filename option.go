package dhcp

import (
	"fmt"
	"strconv"
	"strings"
)

type Option struct {
	Code OptionCode
	Data OptionData
}

type OptionData interface {
	Unmarshal() ([]byte, error)
	String() string
	IsValid() bool
}

func (opt Option) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "%s", OptionCodeToString[opt.Code])
	fmt.Fprintf(&sb, " [%s]: ", strconv.Itoa(int(opt.Code)))
	sb.WriteString(opt.Data.String())

	return sb.String()
}

func (opt Option) IsValid() bool {
	if opt.Code.String() == "" || !opt.Data.IsValid() {
		return false
	}

	return true
}

func (opt Option) Unmarshal() ([]byte, error) {
	data, err := opt.Data.Unmarshal()
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling option. %s", err)
	}
	length := len(data)
	out := make([]byte, 2)
	out[0] = byte(opt.Code)
	out[1] = byte(length)

	return append(out, data...), nil
}
