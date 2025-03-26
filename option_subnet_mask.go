package dhcp

import (
	"errors"
	"net"
)

// OptionDataSubnetMask represents data for the Subnet Mask Option
type OptionDataSubnetMask struct {
	Mask net.IPMask
}

// String returns the subnet mask as a human-readable string
func (optD OptionDataSubnetMask) String() string {
	return optD.Mask.String()
}

// IsValid returns true if optD.Mask represents a valid subnet mask
func (optD OptionDataSubnetMask) IsValid() bool {
	if len(optD.Mask) != 4 {
		return false
	}
	ones, bits := optD.Mask.Size()

	return !(ones == 0 && bits == 0)
}

// Marshal encodes optD as the value for a Subnet Mask Option
func (optD OptionDataSubnetMask) Marshal() ([]byte, error) {
	if !optD.IsValid() {
		return nil, errors.New("option data is invalid")
	}

	return []byte(optD.Mask), nil
}

// UnmarshalOptionDataSubnetMask parses data as the value for a Subnet Mask Option
func UnmarshalOptionDataSubnetMask(data []byte) (OptionData, error) {
	optD := OptionDataSubnetMask{Mask: net.IPMask(data)}
	if !optD.IsValid() {
		return nil, errors.New("data does not represent an ipv4 subnet mask")
	}

	return optD, nil
}

// NewOptionSubnetMask is a helper function for constructing a Subnet Mask Option.
// It will hold OptionDataSubnetMask as the Option's data
func NewOptionSubnetMask(mask net.IPMask) (*Option, error) {
	opt := &Option{
		Code: OptionCodeSubnetMask,
		Data: OptionDataSubnetMask{Mask: mask},
	}
	if !opt.Data.IsValid() {
		return nil, errors.New("mask is invalid")
	}

	return opt, nil
}
