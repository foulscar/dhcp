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
	Raw() []byte
	String() string
}

func (opt Option) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "%s", OptionCodeToInfo[opt.Code].String)
	fmt.Fprintf(&sb, " [%s]: ", strconv.Itoa(int(opt.Code)))
	sb.WriteString(opt.Data.String())

	return sb.String()
}

func (opt Option) Unmarshal() []byte {
	data := opt.Data.Raw()
	length := len(data)
	out := make([]byte, 2)
	out[0] = byte(opt.Code)
	out[1] = byte(length)

	return append(out, data...)
}
