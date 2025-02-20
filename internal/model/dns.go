package model

import "net"

type DNSResult struct {
	FQDN   string   // Fully qualified domain name
	IPList []net.IP // Associated IP addresses
}
