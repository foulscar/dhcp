package dhcp

import (
	"errors"
	"net"
	"strings"
)

// OptionDataRouter represents data for the DHCP Router Option.
// It holds a list of IP addresses for routers on a client's subnet.
// This list should be in order of preference all entries should be unique
type OptionDataRouter struct {
	Routers []net.IP
}

// String returns "[routers_here]" where routers_here represents a
// comma-seprated list of all optD.Routers
func (optD OptionDataRouter) String() string {
	sb := strings.Builder{}

	sb.WriteString("[")
	for i, router := range optD.Routers {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(router.String())
	}
	sb.WriteString("]")

	return sb.String()
}

// IsValid checks if all entries in optD.Routers are unique and valid ipv4 addresses.
// There must be atleast one entry
func (optD OptionDataRouter) IsValid() error {
	if len(optD.Routers) < 1 {
		return errors.New("len of routers list is less than 1")
	}
	for i, router := range optD.Routers {
		if len(router) != 4 {
			return errors.New("routers list contains a router that does not have a len of 4")
		}
		if router.Equal(net.IPv4zero) {
			return errors.New("routers list contains a router with an addr of 0.0.0.0")
		}
		if router[0] == 0x00 {
			return errors.New("routers list contains a router with an addr that starts with 0 (0.*.*.*)")
		}

		for j, routerB := range optD.Routers {
			if i == j {
				continue
			}
			if router.Equal(routerB) {
				return errors.New("routers list contains duplicate router addrs")
			}
		}
	}

	return nil
}

// Marshal encodes optD as the value to a DHCP Router Option
func (optD OptionDataRouter) Marshal() ([]byte, error) {
	if err := optD.IsValid(); err != nil {
		return nil, err
	}

	data := make([]byte, 0)
	for _, router := range optD.Routers {
		data = append(data, []byte(router)...)
	}

	return data, nil
}

// UnmarshalOptionDataRouter parses data as the encoded value for a DHCP Router Option
func UnmarshalOptionDataRouter(data []byte) (OptionData, error) {
	if len(data) < 4 || len(data)%4 != 0 {
		return nil, errors.New("length of data must be atleast 4 and be a multiple of 4")
	}

	numEntries := len(data) / 4
	routers := make([]net.IP, 0)
	for i := 0; i < numEntries; i++ {
		routers = append(routers, net.IP(data[i*4:i*4+4]))
	}

	optD := OptionDataRouter{Routers: routers}
	if err := optD.IsValid(); err != nil {
		return nil, err
	}

	return optD, nil
}

// NewOptionRouter is a helper function for constructing a DHCP Router Option.
// It will hold OptionDataRouter as the Option's data
func NewOptionRouter(routers ...net.IP) (*Option, error) {
	optD := OptionDataRouter{Routers: routers}
	if err := optD.IsValid(); err != nil {
		return nil, errors.New("one of the addresses is invalid")
	}

	opt := Option{
		Code:      OptionCodeRouter,
		Data:      optD,
		IsDefault: false,
	}

	return &opt, nil
}
