package dhcp

// OptionCode represents the type of a DHCP Message Options Entry
type OptionCode uint8

// OptionDataUnmarshaller represents an OptionData constructor
type OptionDataUnmarshaller func([]byte) (OptionData, error)

// String returns the human-readable name represented by the OptionCode
func (code OptionCode) String() string {
	return optMap.ToString[code]
}

// All recognized OptionCodes
const (
	OptionCodePad                                        OptionCode = 0
	OptionCodeSubnetMask                                 OptionCode = 1
	OptionCodeTimeOffset                                 OptionCode = 2
	OptionCodeRouter                                     OptionCode = 3
	OptionCodeTimeServer                                 OptionCode = 4
	OptionCodeNameServer                                 OptionCode = 5
	OptionCodeDNS                                        OptionCode = 6
	OptionCodeLogServer                                  OptionCode = 7
	OptionCodeCookieServer                               OptionCode = 8
	OptionCodeLRPServer                                  OptionCode = 9
	OptionCodeImpressServer                              OptionCode = 10
	OptionCodeResourceLocationServer                     OptionCode = 11
	OptionCodeHostname                                   OptionCode = 12
	OptionCodeBootFileSize                               OptionCode = 13
	OptionCodeMeritDumpFile                              OptionCode = 14
	OptionCodeDomainName                                 OptionCode = 15
	OptionCodeSwapServer                                 OptionCode = 16
	OptionCodeRootPath                                   OptionCode = 17
	OptionCodeExtensionsPath                             OptionCode = 18
	OptionCodeIPForwarding                               OptionCode = 19
	OptionCodeNonLocalSourceRouting                      OptionCode = 20
	OptionCodePolicyFilter                               OptionCode = 21
	OptionCodeMaximumDatagramReassemblySize              OptionCode = 22
	OptionCodeDefaultIPTTL                               OptionCode = 23
	OptionCodePathMTUAgingTimeout                        OptionCode = 24
	OptionCodeMTUPlateauTable                            OptionCode = 25
	OptionCodeInterfaceMTU                               OptionCode = 26
	OptionCodeAllSubnetsLocal                            OptionCode = 27
	OptionCodeBroadcastAddr                              OptionCode = 28
	OptionCodePerformMaskDiscovery                       OptionCode = 29
	OptionCodeMaskSupplier                               OptionCode = 30
	OptionCodePerformRouterDiscovery                     OptionCode = 31
	OptionCodeRouterSolicitationAddress                  OptionCode = 32
	OptionCodeStaticRoute                                OptionCode = 33
	OptionCodeTrailerEncapsulation                       OptionCode = 34
	OptionCodeARPCacheTimeout                            OptionCode = 35
	OptionCodeEthernetEncapsulation                      OptionCode = 36
	OptionCodeTCPDefaultTTL                              OptionCode = 37
	OptionCodeTCPKeepAliveInterval                       OptionCode = 38
	OptionCodeTCPKeepAliveGarbage                        OptionCode = 39
	OptionCodeNetworkInfoServiceDomain                   OptionCode = 40
	OptionCodeNetworkInfoServer                          OptionCode = 41
	OptionCodeNetworkTimeProtocolServers                 OptionCode = 42
	OptionCodeVendorSpecificInfo                         OptionCode = 43
	OptionCodeNetBIOSOverTCPIPNameServer                 OptionCode = 44
	OptionCodeNetBIOSOverTCPIPDatagramDistributionServer OptionCode = 45
	OptionCodeNetBIOSOverTCPIPNodeType                   OptionCode = 46
	OptionCodeNetBIOSOverTCPIPScope                      OptionCode = 47
	OptionCodeXWindowSystemFontServer                    OptionCode = 48
	OptionCodeXWindowSystemDisplayManager                OptionCode = 49
	OptionCodeRequestedIPAddr                            OptionCode = 50
	OptionCodeIPAddrLeaseTime                            OptionCode = 51
	OptionCodeOverload                                   OptionCode = 52
	OptionCodeMessageType                                OptionCode = 53
	OptionCodeServerIdentifier                           OptionCode = 54
	OptionCodeParameterRequestList                       OptionCode = 55
	OptionCodeMessage                                    OptionCode = 56
	OptionCodeMaxMessageSize                             OptionCode = 57
	OptionCodeRenewalTime                                OptionCode = 58
	OptionCodeRebindingTime                              OptionCode = 59
	OptionCodeVendorClassIdentifier                      OptionCode = 60
	OptionCodeClientIdentifier                           OptionCode = 61
	OptionCodeNetworkInfoServicePlusDomain               OptionCode = 64
	OptionCodeNetworkInfoServicePlusServers              OptionCode = 65
	OptionCodeTFTPServerName                             OptionCode = 66
	OptionCodeBootfileName                               OptionCode = 67
	OptionCodeMobileIPHomeAgent                          OptionCode = 68
	OptionCodeSMTPServer                                 OptionCode = 69
	OptionCodePOP3Server                                 OptionCode = 70
	OptionCodeNNTPServer                                 OptionCode = 71
	OptionCodeDefaultWWWServer                           OptionCode = 72
	OptionCodeDefaultFingerServer                        OptionCode = 73
	OptionCodeDefaultIRCServer                           OptionCode = 74
	OptionCodeStreetTalkServer                           OptionCode = 75
	OptionCodeSTDAServer                                 OptionCode = 76
	OptionCodeEnd                                        OptionCode = 255
)

// OptionCodeToString holds the human-readable strings associated with OptionCodes.
// You can change these values before your program's main execution to affect how
// Options/OptionCodes are named when logging, debugging, etc.
var OptionCodeToString = map[OptionCode]string{
	OptionCodeSubnetMask:                                 "Subnet Mask",
	OptionCodeTimeOffset:                                 "Time Offset",
	OptionCodeRouter:                                     "Router",
	OptionCodeTimeServer:                                 "Time Server",
	OptionCodeNameServer:                                 "Name Server",
	OptionCodeDNS:                                        "Domain Name Server",
	OptionCodeLogServer:                                  "Log Server",
	OptionCodeCookieServer:                               "Cookie Server",
	OptionCodeLRPServer:                                  "LPR Server",
	OptionCodeImpressServer:                              "Impress Server",
	OptionCodeResourceLocationServer:                     "Resource Location Server",
	OptionCodeHostname:                                   "Host Name",
	OptionCodeBootFileSize:                               "Boot File Size",
	OptionCodeMeritDumpFile:                              "Merit Dump File",
	OptionCodeDomainName:                                 "Domain Name",
	OptionCodeSwapServer:                                 "Swap Server",
	OptionCodeRootPath:                                   "Root Path",
	OptionCodeExtensionsPath:                             "Extensions Path",
	OptionCodeIPForwarding:                               "IP Forwarding",
	OptionCodeNonLocalSourceRouting:                      "Non-Local Source Routing",
	OptionCodePolicyFilter:                               "Policy Filter",
	OptionCodeMaximumDatagramReassemblySize:              "Maximum Datagram Reassembly Size",
	OptionCodeDefaultIPTTL:                               "Default IP Time-to-Live",
	OptionCodePathMTUAgingTimeout:                        "Path MTU Aging Timeout",
	OptionCodeMTUPlateauTable:                            "Path MTU Plateau Table",
	OptionCodeInterfaceMTU:                               "Interface MTU",
	OptionCodeAllSubnetsLocal:                            "All Subnets are Local",
	OptionCodeBroadcastAddr:                              "Broadcast Address",
	OptionCodePerformMaskDiscovery:                       "Perform Mask Discovery",
	OptionCodeMaskSupplier:                               "Mask Supplier",
	OptionCodePerformRouterDiscovery:                     "Perform Router Discovery",
	OptionCodeRouterSolicitationAddress:                  "Router Solicitation Address",
	OptionCodeStaticRoute:                                "Static Route",
	OptionCodeTrailerEncapsulation:                       "Trailer Encapsulation",
	OptionCodeARPCacheTimeout:                            "ARP Cache Timeout",
	OptionCodeEthernetEncapsulation:                      "Ethernet Encapsulation",
	OptionCodeTCPDefaultTTL:                              "TCP Default TTL",
	OptionCodeTCPKeepAliveInterval:                       "TCP Keepalive Interval",
	OptionCodeTCPKeepAliveGarbage:                        "TCP Keepalive Garbage",
	OptionCodeNetworkInfoServiceDomain:                   "Network Information Service Domain",
	OptionCodeNetworkInfoServer:                          "Network Information Servers",
	OptionCodeNetworkTimeProtocolServers:                 "Network Time Protocol Servers",
	OptionCodeVendorSpecificInfo:                         "Vendor-Specific Information",
	OptionCodeNetBIOSOverTCPIPNameServer:                 "NetBIOS over TCP/IP Name Server",
	OptionCodeNetBIOSOverTCPIPDatagramDistributionServer: "NetBIOS over TCP/IP Datagram Distribution Server",
	OptionCodeNetBIOSOverTCPIPNodeType:                   "NetBIOS over TCP/IP Node Type",
	OptionCodeNetBIOSOverTCPIPScope:                      "NetBIOS over TCP/IP Scope",
	OptionCodeXWindowSystemFontServer:                    "X Window System Font Server",
	OptionCodeXWindowSystemDisplayManager:                "X Window System Display Manager",
	OptionCodeRequestedIPAddr:                            "Requested IP Address",
	OptionCodeIPAddrLeaseTime:                            "IP Address Lease Time",
	OptionCodeOverload:                                   "Overload",
	OptionCodeMessageType:                                "DHCP Message Type",
	OptionCodeServerIdentifier:                           "Server Identifier",
	OptionCodeParameterRequestList:                       "Parameter Request List",
	OptionCodeMessage:                                    "Message",
	OptionCodeMaxMessageSize:                             "Maximum DHCP Message Size",
	OptionCodeRenewalTime:                                "Renewal (T1) Time Value",
	OptionCodeRebindingTime:                              "Rebinding (T2) Time Value",
	OptionCodeVendorClassIdentifier:                      "Vendor Class Identifier",
	OptionCodeClientIdentifier:                           "Client Identifier",
	OptionCodeNetworkInfoServicePlusDomain:               "Network Information Service+ Domain",
	OptionCodeNetworkInfoServicePlusServers:              "Network Information Service+ Servers",
	OptionCodeTFTPServerName:                             "TFTP Server Name",
	OptionCodeBootfileName:                               "Bootfile Name",
	OptionCodeMobileIPHomeAgent:                          "Mobile IP Home Agent",
	OptionCodeSMTPServer:                                 "Simple Mail Transport Protocol Server",
	OptionCodePOP3Server:                                 "Post Office Protocol Server",
	OptionCodeNNTPServer:                                 "Network News Transport Protocol Server",
	OptionCodeDefaultWWWServer:                           "Default World Wide Web Server",
	OptionCodeDefaultFingerServer:                        "Default Finger Server",
	OptionCodeDefaultIRCServer:                           "Default Internet Relay Chat Server",
	OptionCodeStreetTalkServer:                           "StreetTalk Server",
	OptionCodeSTDAServer:                                 "StreetTalk Directory Assistance Server",
}

// OptionCodeToDataUnmarshaller holds OptionData constructors associated with their relevant OptionCode.
// You can change these values before your program's main execution to affect
// the behavior of parsed DHCP Messages
var OptionCodeToDataUnmarshaller = map[OptionCode]OptionDataUnmarshaller{
	OptionCodeSubnetMask:           UnmarshalOptionDataSubnetMask,
	OptionCodeRouter:               UnmarshalOptionDataRouter,
	OptionCodeMessageType:          UnmarshalOptionDataMessageType,
	OptionCodeParameterRequestList: UnmarshalOptionDataParameterRequestList,
}
