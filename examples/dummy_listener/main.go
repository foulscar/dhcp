package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting Server")
	s, err := newServer("enp0s2")
	if err != nil {
		panic(err)
	}

	defer s.stop()
	go s.listenThenClose()

	for {
		msg := <-s.messages

		fmt.Print("\n\nReceived DHCP Message\n\n")
		fmt.Println(msg)
	}
}
