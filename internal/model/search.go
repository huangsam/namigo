package model

import "net"

const (
	NoAuthor      = "No author"      // Fallback for author
	NoDescription = "No description" // Fallback for description
)

type GoPackageResult struct {
	Name        string // Package name
	Path        string // Fully qualified package path
	Description string // Package description
}

type NPMPackageResult struct {
	Name         string // Package name
	Description  string // Package description
	IsExactMatch bool   // Is exact match?
}

type PyPIPackageResult struct {
	Name        string // Package name
	Description string // Package description
	Author      string // Package author
}

type DNSResult struct {
	FQDN   string   // Fully qualified domain name
	IPList []net.IP // Associated IP addresses
}
