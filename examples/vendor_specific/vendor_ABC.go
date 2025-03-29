// This example shows how to incorporate custom (vendor-specific) options, using the dhcp package.
// This custom option will simply store a 'tag'.
// This 'tag' is only valid when it is 'cat' or 'dog'.
// This 'tag' is specific to this example and you could do whatever you want with the data
package main

import (
        "fmt"

        "github.com/foulscar/dhcp"
)

const OptionCodeVendor_ABC dhcp.OptionCode = 222

// This struct implements dhcp.OptionData
type OptionDataVendor_ABC struct {
        Tag string
}

func (optD OptionDataVendor_ABC) String() string {
        switch (optD.Tag) {
        case "cat":
                return "Cat"
        case "dog":
                return "Dog"

        }

        return "Unknown Tag"
}

func (optD OptionDataVendor_ABC) IsValid() bool {
        switch (optD.Tag) {
        case "cat":
                return true
        case "dog":
                return true
        }

        return false
}

func (optD OptionDataVendor_ABC) Marshal() ([]byte, error) {
        return []byte(optD.Tag), nil
}

func UnmarshalOptionDataVendor_ABC(data []byte) (dhcp.OptionData, error) {
        return OptionDataVendor_ABC{ Tag: string(data) }, nil
}

func NewOptionVendor_ABC(tag string) (*dhcp.Option, error) {
        data := OptionDataVendor_ABC{ Tag: tag }
        if !data.IsValid() {
                return nil, fmt.Errorf("invalid tag: '%s'", tag)
        }

        return &dhcp.Option{
                Code: OptionCodeVendor_ABC,
                Data: data,
                IsDefault: false,
        }, nil
}
