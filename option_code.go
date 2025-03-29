package dhcp

// OptionCode represents the type of a DHCP Message Options Entry
type OptionCode uint8

// OptionDataUnmarshaller represents an OptionData constructor
type OptionDataUnmarshaller func([]byte) (OptionData, error)

// String returns the human-readable name represented by the OptionCode
func (code OptionCode) String() string {
	return optMap.ToString[code]
}

// Almost all recognized OptionCodes.
// I have not included OptionCodes that have multiple recognized uses.
// I gathered these from 'https://www.iana.org/assignments/bootp-dhcp-parameters/bootp-dhcp-parameters.xhtml'.
// You can define your own within your own package then modify GlobalOptionCodeMapping
// to contain strings, data types, and unmarshallers
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
	OptionCodeWWWServer                                  OptionCode = 72
	OptionCodeFingerServer                               OptionCode = 73
	OptionCodeIRCServer                                  OptionCode = 74
	OptionCodeStreetTalkServer                           OptionCode = 75
	OptionCodeSTDAServer                                 OptionCode = 76
	OptionCodeUserClass                                  OptionCode = 77
	OptionCodeDirectoryAgent                             OptionCode = 78
	OptionCodeServiceScope                               OptionCode = 79
	OptionCodeRapidCommit                                OptionCode = 80
	OptionCodeClientFQDN                                 OptionCode = 81
	OptionCodeRelayAgentInfo                             OptionCode = 82
	OptionCodeISNS                                       OptionCode = 83
	OptionCodeNDSServers                                 OptionCode = 85
	OptionCodeNDSTreeName                                OptionCode = 86
	OptionCodeNDSContext                                 OptionCode = 87
	OptionCodeBCMCSControllerDomainNameList              OptionCode = 88
	OptionCodeBCMCSControllerIPv4Addr                    OptionCode = 89
	OptionCodeAuth                                       OptionCode = 90
	OptionCodeClientLastTransactionTime                  OptionCode = 91
	OptionCodeAssociatedIP                               OptionCode = 92
	OptionCodeClientSystem                               OptionCode = 93
	OptionCodeClientNDI                                  OptionCode = 94
	OptionCodeLDAP                                       OptionCode = 95
	OptionCodeUUIDGUID                                   OptionCode = 97
	OptionCodeUserAuth                                   OptionCode = 98
	OptionCodeGEOCONFCIVIC                               OptionCode = 99
	OptionCodePCode                                      OptionCode = 100
	OptionCodeTCode                                      OptionCode = 101
	OptionCodeIPv6OnlyPreferred                          OptionCode = 108
	OptionCodeDHCP4O6S46SADDR                            OptionCode = 109
	OptionCodeNetInfoAddr                                OptionCode = 112
	OptionCodeNetInfoTag                                 OptionCode = 113
	OptionCodeDHCPCaptivePortal                          OptionCode = 114
	OptionCodeAutoConfig                                 OptionCode = 116
	OptionCodeNameServiceSearch                          OptionCode = 117
	OptionCodeSubnetSelection                            OptionCode = 118
	OptionCodeDomainSearch                               OptionCode = 119
	OptionCodeSIPServers                                 OptionCode = 120
	OptionCodeClasslessStaticRoute                       OptionCode = 121
	OptionCodeCableLabsClientConfig                      OptionCode = 122
	OptionCodeGeoConf                                    OptionCode = 123
	OptionCodeVendorIdentifyingVendorClass               OptionCode = 124
	OptionCodeVendorIfentifyingSpecificInfo              OptionCode = 125
	OptionCodePANAAgent                                  OptionCode = 136
	OptionCodeLost                                       OptionCode = 137
	OptionCodeCAPWAPAccessControllerAddrs                OptionCode = 138
	OptionCodeIPv4AddrMOS                                OptionCode = 139
	OptionCodeIPv4FQDNMOS                                OptionCode = 140
	OptionCodeSIPUserAgentConfDomains                    OptionCode = 141
	OptionCodeANDSFIPv4Addr                              OptionCode = 142
	OptionCodeSZTPREDIRECT                               OptionCode = 143
	OptionCodeGeospatialLocation                         OptionCode = 144
	OptionCodeForcerenewNonceCapable                     OptionCode = 145
	OptionCodeRDNSSSelection                             OptionCode = 146
	OptionCodeDOTSRI                                     OptionCode = 147
	OptionCodeDOTSAddr                                   OptionCode = 148
	OptionCodeStatusCode                                 OptionCode = 151
	OptionCodeBaseTime                                   OptionCode = 152
	OptionCodeStartTimeOfState                           OptionCode = 153
	OptionCodeQueryStartTime                             OptionCode = 154
	OptionCodeQueryEndTime                               OptionCode = 155
	OptionCodeDHCPState                                  OptionCode = 156
	OptionCodeDataSrc                                    OptionCode = 157
	OptionCodePCPServer                                  OptionCode = 158
	OptionCodePortParams                                 OptionCode = 159
	OptionCodeMUDURL                                     OptionCode = 161
	OptionCodeDNR                                        OptionCode = 162
	OptionCodePXELinuxMagicString                        OptionCode = 208
	OptionCodeConfigFile                                 OptionCode = 209
	OptionCodePathPrefix                                 OptionCode = 210
	OptionCodeRebootTime                                 OptionCode = 211
	OptionCodeAccessDomain                               OptionCode = 213
	OptionCodeSubnetAllocation                           OptionCode = 220
	OptionCodeVirtualSubnetSelection                     OptionCode = 221
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
	OptionCodeWWWServer:                                  "Default World Wide Web Server",
	OptionCodeFingerServer:                               "Default Finger Server",
	OptionCodeIRCServer:                                  "Default Internet Relay Chat Server",
	OptionCodeStreetTalkServer:                           "StreetTalk Server",
	OptionCodeSTDAServer:                                 "StreetTalk Directory Assistance Server",
	OptionCodeUserClass:                                  "User Class",
	OptionCodeDirectoryAgent:                             "Directory Agent",
	OptionCodeServiceScope:                               "Service Scope",
	OptionCodeRapidCommit:                                "Rapid Commit",
	OptionCodeClientFQDN:                                 "Client Fully Qualified Domain Name",
	OptionCodeRelayAgentInfo:                             "Relay Agent Information",
	OptionCodeISNS:                                       "iSNS",
	OptionCodeNDSServers:                                 "NDS Servers",
	OptionCodeNDSTreeName:                                "NDS Tree Name",
	OptionCodeNDSContext:                                 "NDS Context",
	OptionCodeBCMCSControllerDomainNameList:              "BCMCS Controller Domain Name List",
	OptionCodeBCMCSControllerIPv4Addr:                    "BCMCS Controller IPv4 Address",
	OptionCodeAuth:                                       "Authentication",
	OptionCodeClientLastTransactionTime:                  "Client Last Transaction Time",
	OptionCodeAssociatedIP:                               "Associated IP",
	OptionCodeClientSystem:                               "Client System Architecture Type",
	OptionCodeClientNDI:                                  "Client Network Device Interface",
	OptionCodeLDAP:                                       "Lightweight Directory Access Protocol (LDAP)",
	OptionCodeUUIDGUID:                                   "UUID/GUID",
	OptionCodeUserAuth:                                   "User Authentication Protocol",
	OptionCodeGEOCONFCIVIC:                               "GEOCONF Civic",
	OptionCodePCode:                                      "PCode",
	OptionCodeTCode:                                      "TCode",
	OptionCodeIPv6OnlyPreferred:                          "IPv6-Only Preferred",
	OptionCodeDHCP4O6S46SADDR:                            "DHCP 4o6 S46 SADDR",
	OptionCodeNetInfoAddr:                                "NetInfo Address",
	OptionCodeNetInfoTag:                                 "NetInfo Tag",
	OptionCodeDHCPCaptivePortal:                          "DHCP Captive Portal",
	OptionCodeAutoConfig:                                 "Auto Configuration",
	OptionCodeNameServiceSearch:                          "Name Service Search",
	OptionCodeSubnetSelection:                            "Subnet Selection",
	OptionCodeDomainSearch:                               "Domain Search",
	OptionCodeSIPServers:                                 "SIP Servers",
	OptionCodeClasslessStaticRoute:                       "Classless Static Route",
	OptionCodeCableLabsClientConfig:                      "CableLabs Client Configuration",
	OptionCodeGeoConf:                                    "Geospatial Configuration",
	OptionCodeVendorIdentifyingVendorClass:               "Vendor Identifying Vendor Class",
	OptionCodeVendorIfentifyingSpecificInfo:              "Vendor Identifying Specific Information",
	OptionCodePANAAgent:                                  "PANA Agent",
	OptionCodeLost:                                       "Lost",
	OptionCodeCAPWAPAccessControllerAddrs:                "CAPWAP Access Controller Addresses",
	OptionCodeIPv4AddrMOS:                                "IPv4 Address MOS",
	OptionCodeIPv4FQDNMOS:                                "IPv4 FQDN MOS",
	OptionCodeSIPUserAgentConfDomains:                    "SIP User Agent Configuration Domains",
	OptionCodeANDSFIPv4Addr:                              "ANDSF IPv4 Address",
	OptionCodeSZTPREDIRECT:                               "SZTP Redirect",
	OptionCodeGeospatialLocation:                         "Geospatial Location",
	OptionCodeForcerenewNonceCapable:                     "Forcerenew Nonce Capable",
	OptionCodeRDNSSSelection:                             "RDNSS Selection",
	OptionCodeDOTSRI:                                     "DOTS RI",
	OptionCodeDOTSAddr:                                   "DOTS Address",
	OptionCodeStatusCode:                                 "Status Code",
	OptionCodeBaseTime:                                   "Base Time",
	OptionCodeStartTimeOfState:                           "Start Time of State",
	OptionCodeQueryStartTime:                             "Query Start Time",
	OptionCodeQueryEndTime:                               "Query End Time",
	OptionCodeDHCPState:                                  "DHCP State",
	OptionCodeDataSrc:                                    "Data Source",
	OptionCodePCPServer:                                  "PCP Server",
	OptionCodePortParams:                                 "Port Parameters",
	OptionCodeMUDURL:                                     "Manufacturer Usage Descriptions",
	OptionCodeDNR:                                        "Encrypted DNS Server",
	OptionCodePXELinuxMagicString:                        "PXE Linux Magic String",
	OptionCodeConfigFile:                                 "Config File",
	OptionCodePathPrefix:                                 "Path Prefix",
	OptionCodeRebootTime:                                 "Reboot Time",
	OptionCodeAccessDomain:                               "Access Domain",
	OptionCodeSubnetAllocation:                           "Subnet Allocation",
	OptionCodeVirtualSubnetSelection:                     "Virtual Subnet Selection",
}

// OptionCodeToDataType maps OptionData implementations to their relevant OptionCode.
// You can change these values before your program's main execution to affect
// the behavior of this package
var OptionCodeToDataType = map[OptionCode]OptionData{
	OptionCodeSubnetMask:           OptionDataSubnetMask{},
	OptionCodeRouter:               OptionDataRouter{},
	OptionCodeMessageType:          OptionDataMessageType{},
	OptionCodeParameterRequestList: OptionDataParameterRequestList{},
}

// OptionCodeToDataUnmarshaller maps OptionData constructors to their relevant OptionCode.
// You can change these values before your program's main execution to affect
// the behavior of this package
var OptionCodeToDataUnmarshaller = map[OptionCode]OptionDataUnmarshaller{
	OptionCodeSubnetMask:           UnmarshalOptionDataSubnetMask,
	OptionCodeRouter:               UnmarshalOptionDataRouter,
	OptionCodeMessageType:          UnmarshalOptionDataMessageType,
	OptionCodeParameterRequestList: UnmarshalOptionDataParameterRequestList,
}
