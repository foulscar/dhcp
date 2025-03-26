package dhcp

import (
	"fmt"
	"net"
	"strconv"
)

// Conn is used for broadcasting/listening to dhcp packets on a specified interface.
// Conn does not yet implement net.Conn, but I have plans for this
type Conn struct {
	listenConn net.PacketConn
	listenPort int
	sendConn   net.PacketConn
	sendPort   int
	iface      *net.Interface
	dstAddr    *net.UDPAddr
}

// NewConn returns a Conn on a specified interface, with specified send/receive ports.
// You must have the permissions to bind to both
func NewConn(ifaceName string, listenPort, sendPort int) (*Conn, error) {
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

	return &Conn{
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
func (c *Conn) Read(b []byte) (int, error) {
	n, _, err := c.listenConn.ReadFrom(b)
	return n, err
}

// Write broadcasts a packet to the binded broadcasting port.
// You would marshal a Message first
func (c *Conn) Write(b []byte) (int, error) {
	return c.sendConn.WriteTo(b, c.dstAddr)
}

// Close closes both connections for the broadcasting/listening ports
func (c *Conn) Close() error {
	lErr := c.listenConn.Close()
	sErr := c.sendConn.Close()
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
