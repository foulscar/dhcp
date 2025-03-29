// This example shows how to incorporate custom (vendor-specific) options, using the dhcp package.
// This custom option will simply store a 'tag'.
// This 'tag' is only valid when it is 'cat' or 'dog'.
// This 'tag' is specific to this example and you could do whatever you want with the data
package main

import (
	"github.com/foulscar/dhcp"
)

// Here we modify the global OptionCodeMapping.
// This changes the behavior of the dhcp package to incorporate
// our custom (vendor-specific) option
func modifyGlobalMapping() {
	optCodeMap := &dhcp.GlobalOptionCodeMapping

	optCodeMap.ToString[OptionCodeVendor_ABC] = "Vendor ABC Specific Option"
	optCodeMap.ToDataType[OptionCodeVendor_ABC] = OptionDataVendor_ABC{}
	optCodeMap.ToDataUnmarshaller[OptionCodeVendor_ABC] = UnmarshalOptionDataVendor_ABC
}
