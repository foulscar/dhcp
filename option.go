package dhcp

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Option represents a DHCP Message Options Entry
type Option struct {
	Code      OptionCode
	Data      OptionData
	IsDefault bool
}

// OptionData should hold the value of the represented Option Type
type OptionData interface {
	Marshal() ([]byte, error)
	String() string
	IsValid() error
}

// String returns a verbose, human-readable string from opt
func (opt Option) String() string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "%s", optMap.GetString(opt.Code))
	fmt.Fprintf(&sb, " [%s]: ", strconv.Itoa(int(opt.Code)))
	sb.WriteString(opt.Data.String())

	return sb.String()
}

// IsValid checks if opt is a valid Option.
// Will return false if OptionData is nil/invalid or if
// the type of opt.Data != GlobalOptionCodeMapping.ToDataType[opt.Code]
func (opt Option) IsValid() *ErrorExt {
	dataType := reflect.TypeOf(opt.Data)
	errPrefix := "Option with OptionCode '" + strconv.Itoa(int(opt.Code)) + "'"

	if opt.Data == nil {
		return NewErrorExt(errPrefix + " has nil data")
	}
	if dataType != optMap.DefaultDataType && dataType != optMap.ToDataType[opt.Code] {
		return NewErrorExt(errPrefix + "is not using it's assigned OptionData type")
	}
	if err := opt.Data.IsValid(); err != nil {
		return NewErrorExt(errPrefix+" has invalid data", err)
	}

	return nil
}

// Marshal returns an encoded Options entry
func (opt Option) Marshal() ([]byte, *ErrorExt) {
	data, err := opt.Data.Marshal()
	if err != nil {
		return nil, NewErrorExt("error marshalling option", err)
	}
	length := len(data)
	out := make([]byte, 2)
	out[0] = byte(opt.Code)
	out[1] = byte(length)

	return append(out, data...), nil
}
