package dhcp

import (
	"errors"
	"net"
)

// OptionDataSubnetMask represents data for the DHCP Subnet Mask Option
type OptionDataSubnetMask struct {
	Mask net.IPMask
}

// String returns the subnet mask as a human-readable string
func (optD OptionDataSubnetMask) String() string {
	return optD.Mask.String()
}

// IsValid returns true if optD.Mask represents a valid subnet mask
func (optD OptionDataSubnetMask) IsValid() error {
	if len(optD.Mask) != 4 {
		return errors.New("subnet mask has a len != 4")
	}
	ones, bits := optD.Mask.Size()

	if ones == 0 && bits == 0 {
		return errors.New("subnet mask is invalid")
	}

	return nil
}

// Marshal encodes optD as the value for a DHCP Subnet Mask Option
func (optD OptionDataSubnetMask) Marshal() ([]byte, error) {
	if err := optD.IsValid(); err != nil {
		return nil, err
	}

	return []byte(optD.Mask), nil
}

// UnmarshalOptionDataSubnetMask parses data as the value for a DHCP Subnet Mask Option
func UnmarshalOptionDataSubnetMask(data []byte) (OptionData, error) {
	optD := OptionDataSubnetMask{Mask: net.IPMask(data)}
	if err := optD.IsValid(); err != nil {
		return nil, err
	}

	return optD, nil
}

// NewOptionSubnetMask is a helper function for constructing a Subnet Mask Option.
// It will hold OptionDataSubnetMask as the Option's data
func NewOptionSubnetMask(mask net.IPMask) (*Option, error) {
	opt := &Option{
		Code:      OptionCodeSubnetMask,
		Data:      OptionDataSubnetMask{Mask: mask},
		IsDefault: false,
	}
	if err := opt.Data.IsValid(); err != nil {
		return nil, err
	}

	return opt, nil
}
