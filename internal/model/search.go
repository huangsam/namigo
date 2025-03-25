package model

import "net"

const (
	NoAuthor      = "No author"      // Fallback for author
	NoDescription = "No description" // Fallback for description
)

type SearchRecordKey int

const (
	UnknownKey SearchRecordKey = iota
	GoKey
	NPMKey
	PyPIKey
	DNSKey
	EmailKey
)

func (k SearchRecordKey) String() string {
	switch k {
	case GoKey:
		return "Golang"
	case NPMKey:
		return "NPM"
	case PyPIKey:
		return "PyPI"
	case DNSKey:
		return "DNS"
	case EmailKey:
		return "Email"
	default:
		return "Unknown"
	}
}

// SearchRecord is an union type for all search Record values
type SearchRecord interface{ Record() }

type GoPackage struct {
	Name        string // Package name
	Path        string // Fully qualified package path
	Description string // Package description
}

func (*GoPackage) Record() {}

type NPMPackage struct {
	Name        string // Package name
	Description string // Package description
}

func (*NPMPackage) Record() {}

type PyPIPackage struct {
	Name        string // Package name
	Author      string // Package author
	Description string // Package description
}

func (*PyPIPackage) Record() {}

type DNSRecord struct {
	FQDN   string   // Fully qualified domain name
	IPList []net.IP // Associated IP addresses
}

func (*DNSRecord) Record() {}

type EmailRecord struct {
	Addr           string // Email address
	HasValidSyntax bool   // Email address has valid syntax
	HasValidDomain bool   // Email address has valid domain
}

func (*EmailRecord) Record() {}

type SearchResult struct {
	Key     SearchRecordKey
	Records []SearchRecord
}
