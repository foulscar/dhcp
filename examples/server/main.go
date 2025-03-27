package main

import (
	"fmt"
)

func main() {
	s, err := newServer("enp0s1")
	if err != nil {
		panic(err)
	}

	go s.listenThenClose()

	defer s.stop()

	for {
		msg := <-s.messages
		fmt.Println(msg.String())
	}
}
