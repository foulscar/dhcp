package main

import (
	"fmt"
	"github.com/foulscar/dhcp"
)

type server struct {
	conn     *dhcp.Conn
	messages chan *dhcp.Message
	stopChan chan bool
}

func newServer(ifaceName string) (*server, error) {
	s := &server{}
	conn, err := dhcp.NewConn(ifaceName, 67, 68)
	if err != nil {
		return nil, fmt.Errorf("could not create server. %w", err)
	}

	s.conn = conn
	s.messages = make(chan *dhcp.Message)
	s.stopChan = make(chan bool)

	return s, nil
}

func (s *server) listenThenClose() {
	data := make([]byte, 2048)
	for {
		if <-s.stopChan {
			break
		}

		n, _ := s.conn.Read(data)
		if n < 240 {
			continue
		}

		msg, err := dhcp.UnmarshalMessage(data[:n])
		if err != nil {
			continue
		}

		s.messages <- msg
	}
	s.conn.Close()
}

func (s *server) stop() {
	s.stopChan <- true
}
