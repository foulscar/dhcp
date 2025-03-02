package dhcp

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// Implements net.Conn
type DHCPConn struct {
	listenConn net.PacketConn
	listenPort int
	sendConn   net.PacketConn
	sendPort   int
	iface      *net.Interface
}

func NewDHCPConn(ifaceName string, listenPort, sendPort int) (*DHCPConn, error) {
	iface, err := net.InterfaceByName(ifaceName)
	if err != nil {
		return nil, fmt.Errorf("error finding interface %s: %w", ifaceName, err)
	}

	listenConn, err := net.ListenPacket("udp4", ":"+strconv.Itoa(listenPort))
	if err != nil {
		return nil, fmt.Errorf("failed to listen on UDP "+strconv.Itoa(listenPort)+": %w", err)
	}

	sendConn, err := net.ListenPacket("udp4", "0.0.0.0:0")
	if err != nil {
		listenConn.Close()
		return nil, fmt.Errorf("failed to create UDP send socket: %w", err)
	}

	return &DHCPConn{
		listenConn: listenConn,
		listenPort: listenPort,
		sendConn:   sendConn,
		sendPort:   sendPort,
		iface:      iface,
	}, nil
}

func (dc *DHCPConn) Read(b []byte) (int, error) {
	n, _, err := dc.listenConn.ReadFrom(b)
	return n, err
}

func (dc *DHCPConn) Write(b []byte) (int, error) {
	dstAddr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: dc.sendPort,
	}
	return dc.sendConn.WriteTo(b, dstAddr)
}

func (dc *DHCPConn) Close() error {
	lErr := dc.listenConn.Close()
	sErr := dc.sendConn.Close()
	if lErr != nil || sErr != nil {
		return fmt.Errorf("listenConn: %w, sendConn: %w", lErr, sErr)
	}
	return nil
}

func (dc *DHCPConn) SetDeadline(t time.Time) error {
	lErr := dc.listenConn.SetDeadline(t)
	sErr := dc.sendConn.SetDeadline(t)
	if lErr != nil || sErr != nil {
		return fmt.Errorf("listenConn: %w, sendConn: %w", lErr, sErr)
	}
	return nil
}

func (dc *DHCPConn) SetReadDeadline(t time.Time) error {
	return dc.listenConn.SetReadDeadline(t)
}

func (dc *DHCPConn) SetWriteDeadline(t time.Time) error {
	return dc.sendConn.SetWriteDeadline(t)
}
