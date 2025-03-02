package dhcp

import (
	"errors"
	"net"
)

type OptionDataSubnetMask struct {
	Mask net.IPMask
}

func (optD OptionDataSubnetMask) Raw() []byte {
	return []byte(optD.Mask)
}

func MarshalOptionDataSubnetMask(data []byte) (OptionData, error) {
	if len(data) != 4 {
		return nil, errors.New("Data does not represent an ipv4 subnet mask")
	}

	optD := OptionDataSubnetMask{Mask: net.IPMask(data)}
	return optD, nil
}
