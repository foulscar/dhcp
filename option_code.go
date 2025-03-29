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
	OptionCodeLRP_Server                                 OptionCode = 9
	OptionCodeImpressServer                              OptionCode = 10
	OptionCodeResourceLocationServer                     OptionCode = 11
	OptionCodeHostname                                   OptionCode = 12
	OptionCodeBootFileSize                               OptionCode = 13
	OptionCodeMeritDumpFile                              OptionCode = 14
	OptionCodeDomainName                                 OptionCode = 15
	OptionCodeSwapServer                                 OptionCode = 16
	OptionCodeRootPath                                   OptionCode = 17
	OptionCodeExtensionsPath                             OptionCode = 18
	OptionCodeIP_Forwarding                              OptionCode = 19
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
	OptionCodeARP_CacheTimeout                           OptionCode = 35
	OptionCodeEthernetEncapsulation                      OptionCode = 36
	OptionCodeTCP_DefaultTTL                             OptionCode = 37
	OptionCodeTCP_KeepAliveInterval                      OptionCode = 38
	OptionCodeTCP_KeepAliveGarbage                       OptionCode = 39
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
	OptionCodeIP_AddrLeaseTime                           OptionCode = 51
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
	OptionCodeTFTP_ServerName                            OptionCode = 66
	OptionCodeBootfileName                               OptionCode = 67
	OptionCodeMobileIPHomeAgent                          OptionCode = 68
	OptionCodeSMTP_Server                                OptionCode = 69
	OptionCodePOP3_Server                                OptionCode = 70
	OptionCodeNNTP_Server                                OptionCode = 71
	OptionCodeWWW_Server                                 OptionCode = 72
	OptionCodeFingerServer                               OptionCode = 73
	OptionCodeIRC_Server                                 OptionCode = 74
	OptionCodeStreetTalkServer                           OptionCode = 75
	OptionCodeSTDA_Server                                OptionCode = 76
	OptionCodeUserClass                                  OptionCode = 77
	OptionCodeDirectoryAgent                             OptionCode = 78
	OptionCodeServiceScope                               OptionCode = 79
	OptionCodeRapidCommit                                OptionCode = 80
	OptionCodeClientFQDN                                 OptionCode = 81
	OptionCodeRelayAgentInfo                             OptionCode = 82
	OptionCodeISNS                                       OptionCode = 83
	OptionCodeNDS_Servers                                OptionCode = 85
	OptionCodeNDS_TreeName                               OptionCode = 86
	OptionCodeNDS_Context                                OptionCode = 87
	OptionCodeBCMCS_ControllerDomainNameList             OptionCode = 88
	OptionCodeBCMCS_ControllerIPv4Addr                   OptionCode = 89
	OptionCodeAuth                                       OptionCode = 90
	OptionCodeClientLastTransactionTime                  OptionCode = 91
	OptionCodeAssociatedIP                               OptionCode = 92
	OptionCodeClientSystem                               OptionCode = 93
	OptionCodeClientNDI                                  OptionCode = 94
	OptionCodeLDAP                                       OptionCode = 95
	OptionCodeUUID_GUID                                  OptionCode = 97
	OptionCodeUserAuth                                   OptionCode = 98
	OptionCodeGEOCONF_CIVIC                              OptionCode = 99
	OptionCodePCode                                      OptionCode = 100
	OptionCodeTCode                                      OptionCode = 101
	OptionCodeIPv6OnlyPreferred                          OptionCode = 108
	OptionCodeDHCP4O6_S46_SADDR                          OptionCode = 109
	OptionCodeNetInfoAddr                                OptionCode = 112
	OptionCodeNetInfoTag                                 OptionCode = 113
	OptionCodeDHCP_CaptivePortal                         OptionCode = 114
	OptionCodeAutoConfig                                 OptionCode = 116
	OptionCodeNameServiceSearch                          OptionCode = 117
	OptionCodeSubnetSelection                            OptionCode = 118
	OptionCodeDomainSearch                               OptionCode = 119
	OptionCodeSIP_Servers                                OptionCode = 120
	OptionCodeClasslessStaticRoute                       OptionCode = 121
	OptionCodeCableLabsClientConfig                      OptionCode = 122
	OptionCodeGeoConf                                    OptionCode = 123
	OptionCodeVendorIdentifyingVendorClass               OptionCode = 124
	OptionCodeVendorIfentifyingSpecificInfo              OptionCode = 125
	OptionCodePANA_Agent                                 OptionCode = 136
	OptionCodeLost                                       OptionCode = 137
	OptionCodeCAPWAP_AccessControllerAddrs               OptionCode = 138
	OptionCodeIPv4AddrMOS                                OptionCode = 139
	OptionCodeIPv4FQDN_MOS                               OptionCode = 140
	OptionCodeSIP_UserAgentConfDomains                   OptionCode = 141
	OptionCodeANDSF_IPv4Addr                             OptionCode = 142
	OptionCodeSZTP_REDIRECT                              OptionCode = 143
	OptionCodeGeospatialLocation                         OptionCode = 144
	OptionCodeForcerenewNonceCapable                     OptionCode = 145
	OptionCodeRDNSS_Selection                            OptionCode = 146
	OptionCodeDOTS_RI                                    OptionCode = 147
	OptionCodeDOTS_Addr                                  OptionCode = 148
	OptionCodeStatusCode                                 OptionCode = 151
	OptionCodeBaseTime                                   OptionCode = 152
	OptionCodeStartTimeOfState                           OptionCode = 153
	OptionCodeQueryStartTime                             OptionCode = 154
	OptionCodeQueryEndTime                               OptionCode = 155
	OptionCodeDHCP_State                                 OptionCode = 156
	OptionCodeDataSrc                                    OptionCode = 157
	OptionCodePCP_Server                                 OptionCode = 158
	OptionCodePortParams                                 OptionCode = 159
	OptionCodeMUD_URL                                    OptionCode = 161
	OptionCodeDNR                                        OptionCode = 162
	OptionCodePXE_LinuxMagicString                       OptionCode = 208
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
	OptionCodeLRP_Server:                                 "LPR Server",
	OptionCodeImpressServer:                              "Impress Server",
	OptionCodeResourceLocationServer:                     "Resource Location Server",
	OptionCodeHostname:                                   "Host Name",
	OptionCodeBootFileSize:                               "Boot File Size",
	OptionCodeMeritDumpFile:                              "Merit Dump File",
	OptionCodeDomainName:                                 "Domain Name",
	OptionCodeSwapServer:                                 "Swap Server",
	OptionCodeRootPath:                                   "Root Path",
	OptionCodeExtensionsPath:                             "Extensions Path",
	OptionCodeIP_Forwarding:                              "IP Forwarding",
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
	OptionCodeARP_CacheTimeout:                           "ARP Cache Timeout",
	OptionCodeEthernetEncapsulation:                      "Ethernet Encapsulation",
	OptionCodeTCP_DefaultTTL:                             "TCP Default TTL",
	OptionCodeTCP_KeepAliveInterval:                      "TCP Keepalive Interval",
	OptionCodeTCP_KeepAliveGarbage:                       "TCP Keepalive Garbage",
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
	OptionCodeIP_AddrLeaseTime:                           "IP Address Lease Time",
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
	OptionCodeTFTP_ServerName:                            "TFTP Server Name",
	OptionCodeBootfileName:                               "Bootfile Name",
	OptionCodeMobileIPHomeAgent:                          "Mobile IP Home Agent",
	OptionCodeSMTP_Server:                                "Simple Mail Transport Protocol Server",
	OptionCodePOP3_Server:                                "Post Office Protocol Server",
	OptionCodeNNTP_Server:                                "Network News Transport Protocol Server",
	OptionCodeWWW_Server:                                 "Default World Wide Web Server",
	OptionCodeFingerServer:                               "Default Finger Server",
	OptionCodeIRC_Server:                                 "Default Internet Relay Chat Server",
	OptionCodeStreetTalkServer:                           "StreetTalk Server",
	OptionCodeSTDA_Server:                                "StreetTalk Directory Assistance Server",
	OptionCodeUserClass:                                  "User Class",
	OptionCodeDirectoryAgent:                             "Directory Agent",
	OptionCodeServiceScope:                               "Service Scope",
	OptionCodeRapidCommit:                                "Rapid Commit",
	OptionCodeClientFQDN:                                 "Client Fully Qualified Domain Name",
	OptionCodeRelayAgentInfo:                             "Relay Agent Information",
	OptionCodeISNS:                                       "iSNS",
	OptionCodeNDS_Servers:                                "NDS Servers",
	OptionCodeNDS_TreeName:                               "NDS Tree Name",
	OptionCodeNDS_Context:                                "NDS Context",
	OptionCodeBCMCS_ControllerDomainNameList:             "BCMCS Controller Domain Name List",
	OptionCodeBCMCS_ControllerIPv4Addr:                   "BCMCS Controller IPv4 Address",
	OptionCodeAuth:                                       "Authentication",
	OptionCodeClientLastTransactionTime:                  "Client Last Transaction Time",
	OptionCodeAssociatedIP:                               "Associated IP",
	OptionCodeClientSystem:                               "Client System Architecture Type",
	OptionCodeClientNDI:                                  "Client Network Device Interface",
	OptionCodeLDAP:                                       "Lightweight Directory Access Protocol (LDAP)",
	OptionCodeUUID_GUID:                                  "UUID/GUID",
	OptionCodeUserAuth:                                   "User Authentication Protocol",
	OptionCodeGEOCONF_CIVIC:                              "GEOCONF Civic",
	OptionCodePCode:                                      "PCode",
	OptionCodeTCode:                                      "TCode",
	OptionCodeIPv6OnlyPreferred:                          "IPv6-Only Preferred",
	OptionCodeDHCP4O6_S46_SADDR:                          "DHCP 4o6 S46 SADDR",
	OptionCodeNetInfoAddr:                                "NetInfo Address",
	OptionCodeNetInfoTag:                                 "NetInfo Tag",
	OptionCodeDHCP_CaptivePortal:                         "DHCP Captive Portal",
	OptionCodeAutoConfig:                                 "Auto Configuration",
	OptionCodeNameServiceSearch:                          "Name Service Search",
	OptionCodeSubnetSelection:                            "Subnet Selection",
	OptionCodeDomainSearch:                               "Domain Search",
	OptionCodeSIP_Servers:                                "SIP Servers",
	OptionCodeClasslessStaticRoute:                       "Classless Static Route",
	OptionCodeCableLabsClientConfig:                      "CableLabs Client Configuration",
	OptionCodeGeoConf:                                    "Geospatial Configuration",
	OptionCodeVendorIdentifyingVendorClass:               "Vendor Identifying Vendor Class",
	OptionCodeVendorIfentifyingSpecificInfo:              "Vendor Identifying Specific Information",
	OptionCodePANA_Agent:                                 "PANA Agent",
	OptionCodeLost:                                       "Lost",
	OptionCodeCAPWAP_AccessControllerAddrs:               "CAPWAP Access Controller Addresses",
	OptionCodeIPv4AddrMOS:                                "IPv4 Address MOS",
	OptionCodeIPv4FQDN_MOS:                               "IPv4 FQDN MOS",
	OptionCodeSIP_UserAgentConfDomains:                   "SIP User Agent Configuration Domains",
	OptionCodeANDSF_IPv4Addr:                             "ANDSF IPv4 Address",
	OptionCodeSZTP_REDIRECT:                              "SZTP Redirect",
	OptionCodeGeospatialLocation:                         "Geospatial Location",
	OptionCodeForcerenewNonceCapable:                     "Forcerenew Nonce Capable",
	OptionCodeRDNSS_Selection:                            "RDNSS Selection",
	OptionCodeDOTS_RI:                                    "DOTS RI",
	OptionCodeDOTS_Addr:                                  "DOTS Address",
	OptionCodeStatusCode:                                 "Status Code",
	OptionCodeBaseTime:                                   "Base Time",
	OptionCodeStartTimeOfState:                           "Start Time of State",
	OptionCodeQueryStartTime:                             "Query Start Time",
	OptionCodeQueryEndTime:                               "Query End Time",
	OptionCodeDHCP_State:                                 "DHCP State",
	OptionCodeDataSrc:                                    "Data Source",
	OptionCodePXE_LinuxMagicString:                       "PXE Linux Magic String",
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
