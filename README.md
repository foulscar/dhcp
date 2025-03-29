# dhcp
DHCPv4 encoding/decoding package, with customizable behvaior
## What dhcp comes with builtin
- Basic encoding/decoding of messages and options
- Constants of almost all recognized DHCP Option Codes, with a mapping to a human-readable string
- Unique handlers for basic/common Options (You can replace them with your own)
- Easy binding to a specific interface for receiving/broadcasting dhcp messages
## Where dhcp makes customization easy
- Ability to give different Options unique handling/behavior
- Ability to override builtin option handlers and/or add vendor-specific options
- Modifying mappings changes the behvaior of the dhcp package (this was planned)
