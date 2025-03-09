package dhcp

import (
	"errors"
	"net"
)

type OptionDataSubnetMask struct {
	Mask net.IPMask
}

func (optD OptionDataSubnetMask) String() string {
	return optD.Mask.String()
}

func (optD OptionDataSubnetMask) IsValid() bool {
	if len(optD.Mask) != 4 {
		return false
	}
	ones, bits := optD.Mask.Size()

	return !(ones == 0 && bits == 0)
}

func (optD OptionDataSubnetMask) Unmarshal() ([]byte, error) {
	if !optD.IsValid() {
		return nil, errors.New("option data is invalid")
	}

	return []byte(optD.Mask), nil
}

func MarshalOptionDataSubnetMask(data []byte) (OptionData, error) {
	optD := OptionDataSubnetMask{Mask: net.IPMask(data)}
	if !optD.IsValid() {
		return nil, errors.New("data does not represent an ipv4 subnet mask")
	}

	return optD, nil
}

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
