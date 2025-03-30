// This example shows how to incorporate custom (vendor-specific) options, using the dhcp package.
// This custom option will simply store a 'tag'.
// This 'tag' is only valid when it is 'cat' or 'dog'.
// This 'tag' is specific to this example and you could do whatever you want with the data
package main

import (
	"errors"
	"fmt"

	"github.com/foulscar/dhcp"
)

const OptionCodeVendorABC dhcp.OptionCode = 222

// This struct implements dhcp.OptionData
type OptionDataVendorABC struct {
	Tag string
}

func (optD OptionDataVendorABC) String() string {
	switch optD.Tag {
	case "cat":
		return "Cat"
	case "dog":
		return "Dog"

	}

	return "Unknown Tag"
}

func (optD OptionDataVendorABC) IsValid() error {
	switch optD.Tag {
	case "cat":
		return nil
	case "dog":
		return nil
	}

	return errors.New("data contains invalid tag")
}

func (optD OptionDataVendorABC) Marshal() ([]byte, error) {
	return []byte(optD.Tag), nil
}

func UnmarshalOptionDataVendorABC(data []byte) (dhcp.OptionData, error) {
	return OptionDataVendorABC{Tag: string(data)}, nil
}

func NewOptionVendorABC(tag string) (*dhcp.Option, error) {
	data := OptionDataVendorABC{Tag: tag}
	if err := data.IsValid(); err != nil {
		return nil, fmt.Errorf("invalid tag: '%s'", tag)
	}

	return &dhcp.Option{
		Code:      OptionCodeVendorABC,
		Data:      data,
		IsDefault: false,
	}, nil
}
