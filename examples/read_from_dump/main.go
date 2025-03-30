package main

import (
	"fmt"
	"github.com/foulscar/dhcp"
	"os"
)

func main() {
	data, err := os.ReadFile("ack.raw")
	if err != nil {
		panic(err)
	}

	msg, msgErr := dhcp.UnmarshalMessage(data)
	if msgErr != nil {
		panic(msgErr)
	}

	if err := msg.IsValid(); err != nil {
		fmt.Println("message is invalid.", err)
		os.Exit(1)
	}

	fmt.Println(*msg)
}
