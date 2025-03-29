# dhcp
DHCPv4 encoding/decoding package, with the ability to add user-defined behavior
## What dhcp comes with builtin
- Basic encoding/decoding of messages and options
- Easy binding to a specific interface for receiving/broadcasting dhcp messages
- Constants of almost all recognized DHCP Option Codes, with a mapping to a human-readable string
- Unique handlers for basic/common Options (You can replace them with your own)
## Where dhcp makes customization easy
- Ability to give different Options user-defined handling/behavior, using global mappings
- Ability to override builtin option handlers and/or add vendor-specific options
- Modifying mappings changes behavior across almost the entire package
## Usage
To use this package, find the latest stable tag and run
```bash
go get github.com/foulscar/dhcp@LATEST_STABLE_TAG_HERE
```
## Documentation
[Go Docs](https://pkg.go.dev/github.com/foulscar/dhcp)
## Examples
Below are examples showcasing various usecases

- [Basic Listening and Decoding](examples/dummy_listener)
- [Implementing Custom Options/Handlers](examples/vendor_specific)
