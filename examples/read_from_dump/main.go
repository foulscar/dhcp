package main

import (
	"fmt"
	"github.com/foulscar/dhcp"
	"os"
        "reflect"
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

        fmt.Println(*msg)
        fmt.Println(reflect.TypeOf(msg.Options[6].Data))
        fmt.Println(reflect.TypeOf(dhcp.GlobalOptionCodeMapping.ToDataType[6]))

	valid, reason := msg.IsValid()
	if !valid {
                fmt.Println("message is invalid. reason:", reason)
                os.Exit(1)
	}

	fmt.Println(*msg)
}
