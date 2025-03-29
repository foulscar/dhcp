package main

import (
	"fmt"
	"github.com/foulscar/dhcp"
)

type server struct {
	conn     *dhcp.Conn
	messages chan *dhcp.Message
}

func newServer(ifaceName string) (*server, error) {
	s := &server{}
	conn, err := dhcp.NewConn(ifaceName, 67, 68)
	if err != nil {
		return nil, fmt.Errorf("could not create server. %w", err)
	}

	s.conn = conn
	s.messages = make(chan *dhcp.Message)

	return s, nil
}

func (s *server) listenThenClose() {
        fmt.Println("Server Listening")
	buffer := make([]byte, 2048)
	for {
		n, err := s.conn.Read(buffer)
                if err != nil {
                        continue
                }

                if !dhcp.IsEncodedMessage(buffer[:n]) {
                        continue
                }

		msg, err := dhcp.UnmarshalMessage(buffer[:n])
		if err != nil {
			continue
		}

		s.messages <- msg
	}
}

func (s *server) stop() {
        s.conn.Close()
}
