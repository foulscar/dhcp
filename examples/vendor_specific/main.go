// This example shows how to incorporate custom (vendor-specific) options, using the dhcp package.
// This custom option will simply store a 'tag'.
// This 'tag' is only valid when it is 'cat' or 'dog'.
// This 'tag' is specific to this example and you could do whatever you want with the data
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

	ourOpt, err := NewOptionVendorABC("dog")
	if err != nil {
		fmt.Println("error creating ourOpt.", err)
		os.Exit(1)
	}

	msgA.AddOptions(ourOpt)

	if err := msgA.IsValid(); err != nil {
		fmt.Println("msgA is invalid.", err)
		os.Exit(1)
	}

	rawData, errA := dhcp.MarshalMessage(&msgA)
	if errA != nil {
		fmt.Println("error marshalling msgA.", errA)
		os.Exit(1)
	}

	msgB, errB := dhcp.UnmarshalMessage(rawData)
	if errB != nil {
		fmt.Println("error unmarshalling msgB.", errB)
		os.Exit(1)
	}

	fmt.Print("This is our original message (msgA):\n\n")
	fmt.Print(msgA, "\n\n")

	fmt.Print("This is our message after marshalling then unmarshalling (msgB):\n\n")
	fmt.Println(*msgB)
}
