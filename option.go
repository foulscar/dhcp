package dhcp

import (
	"fmt"
	"strconv"
	"strings"
        "reflect"
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
	IsValid() bool
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
func (opt Option) IsValid() bool {
	if opt.Data == nil || reflect.TypeOf(opt.Data) != reflect.TypeOf(optMap.ToDataType[opt.Code]) || !opt.Data.IsValid() {
		return false
	}

	return true
}

// Marshal returns an encoded Options entry
func (opt Option) Marshal() ([]byte, error) {
	data, err := opt.Data.Marshal()
	if err != nil {
		return nil, fmt.Errorf("error marshalling option. %s", err)
	}
	length := len(data)
	out := make([]byte, 2)
	out[0] = byte(opt.Code)
	out[1] = byte(length)

	return append(out, data...), nil
}
