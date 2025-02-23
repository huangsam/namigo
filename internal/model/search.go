package model

import "net"

const (
	NoAuthor      = "No author"      // Fallback for author
	NoDescription = "No description" // Fallback for description
)

type GoPackage struct {
	Name        string // Package name
	Path        string // Fully qualified package path
	Description string // Package description
}

type NPMPackage struct {
	Name         string // Package name
	Description  string // Package description
	IsExactMatch bool   // Is exact match?
}

type PyPIPackage struct {
	Name        string // Package name
	Description string // Package description
	Author      string // Package author
}

type DNSRecord struct {
	FQDN   string   // Fully qualified domain name
	IPList []net.IP // Associated IP addresses
}
