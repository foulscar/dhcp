# dhcp
A limited but capable dhcpv4 encoding/decoding package. Work in progress.
## TODO List
- [x] Basic message encoding/decoding
- [x] An interface for building options, each with their own behavior
- [x] A wrapper over net.Conn for receiving and broadcasting dhcp messages over a specified interface
- [ ] Implement all dhcp option codes
- [ ] Ensure package is RFC 2131 compliant
