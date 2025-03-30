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

	msg, err := dhcp.UnmarshalMessage(data)
	if err != nil {
		panic(err)
	}

	valid, reason := msg.IsValid()
	if !valid {
		fmt.Println("message is invalid. reason:", reason)
		os.Exit(1)
	}

	fmt.Println(*msg)
}
