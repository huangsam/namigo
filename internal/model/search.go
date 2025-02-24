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
	Name        string // Package name
	Description string // Package description
}

type PyPIPackage struct {
	Name        string // Package name
	Author      string // Package author
	Description string // Package description
}

type DNSRecord struct {
	FQDN   string   // Fully qualified domain name
	IPList []net.IP // Associated IP addresses
}

type EmailRecord struct {
	Addr           string // Email address
	HasValidSyntax bool   // Email address has valid syntax
	HasValidDomain bool   // Email address has valid domain
}
