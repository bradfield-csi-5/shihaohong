package main

const (
	SERVER_PORT = 53
	SERVER_TYPE = "udp"
)

const (
	IPv4len                        = 4
	DNSHeaderLength                = 12
	DNSStandardQueryHeaderFlags    = 0x0120 // Wireshark val, TODO: decompose it
	DNSStandardResponseHeaderFlags = 0x8180 // Wireshark val, TODO: decompose it
	DNSMessageMaxLength            = 512    // RFC 1035, 2.3.4
)
