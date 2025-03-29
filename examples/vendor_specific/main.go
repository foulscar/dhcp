// This example shows how to incorporate custom (vendor-specific) options, using the dhcp package.
// This custom option will simply store a 'tag'.
// This 'tag' is only valid when it is 'cat' or 'dog'.
// This 'tag' is specific to this example and you could do whatever you want with the dataackage main
package main

import (
	"fmt"
	"github.com/foulscar/dhcp"
	"os"
)

func main() {
	modifyGlobalMapping()

	msgA := dhcp.NewMessage()
	msgA.SetMessageType(dhcp.OptionMessageTypeCodeDISCOVER)
	msgA.BOOTPMessageType = dhcp.BOOTPMessageTypeRequest

	ourOpt, err := NewOptionVendor_ABC("dog")
	if err != nil {
		fmt.Println("error creating ourOpt.", err)
		os.Exit(1)
	}

	msgA.AddOptions(ourOpt)

	valid, reason := msgA.IsValid()
	if !valid {
		fmt.Println("msgA is invalid.", reason)
		os.Exit(1)
	}

	rawData, err := dhcp.MarshalMessage(&msgA)
	if err != nil {
		fmt.Println("error marshalling msgA.", err)
		os.Exit(1)
	}

	msgB, err := dhcp.UnmarshalMessage(rawData)
	if err != nil {
		fmt.Println("error unmarshalling msgB.", err)
		os.Exit(1)
	}

	fmt.Println("This is our original message (msgA):\n")
	fmt.Println(msgA)

	fmt.Println("\nThis is our message after marshalling then unmarshalling (msgB):\n")
	fmt.Println(*msgB)
}
