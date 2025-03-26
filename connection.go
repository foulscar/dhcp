package dhcp

import (
	"fmt"
	"net"
	"strconv"
)

// DHCPConn is used for broadcasting/listening to dhcp packets on a specified interface.
type DHCPConn struct {
	listenConn net.PacketConn
	listenPort int
	sendConn   net.PacketConn
	sendPort   int
	iface      *net.Interface
	dstAddr    *net.UDPAddr
}

// NewDHCPConn returns a DHCPConn on a specified interface, with specified send/receive ports.
// You must have the permissions to bind to both
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
		dstAddr: &net.UDPAddr{
			IP:   net.IPv4bcast,
			Port: sendPort,
		},
	}, nil
}

// Read reads data from the binded listening port.
// You would marshal this data into a Message
func (dc *DHCPConn) Read(b []byte) (int, error) {
	n, _, err := dc.listenConn.ReadFrom(b)
	return n, err
}

// Write broadcasts a packet to the binded broadcasting port.
// You would marshal a Message first
func (dc *DHCPConn) Write(b []byte) (int, error) {
	return dc.sendConn.WriteTo(b, dc.dstAddr)
}

// Close closes both connections for the broadcasting/listening ports
func (dc *DHCPConn) Close() error {
	lErr := dc.listenConn.Close()
	sErr := dc.sendConn.Close()
	if lErr != nil || sErr != nil {
		return fmt.Errorf("listenConn: %w, sendConn: %w", lErr, sErr)
	}
	return nil
}

/* Will implement net.Conn soon
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
*/
