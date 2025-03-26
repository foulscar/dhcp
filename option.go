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
	Marshal() ([]byte, error)
	String() string
	IsValid() bool
}

// IsDefault checks if opt uses default Option handling.
// See OptionDataDefault
func (opt Option) IsDefault() bool {
	_, ok := opt.Data.(OptionDataDefault)
	return ok
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
// Will return false if OptionData is invalid or the OptionCode is not recognized
func (opt Option) IsValid() bool {
	if opt.Code.String() == "" || !opt.Data.IsValid() {
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
