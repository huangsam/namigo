package model

import "net"

const (
	// NoAuthor is a fallback for author.
	NoAuthor = "No author"

	// NoDescription is a fallback for description.
	NoDescription = "No description"
)

// SearchKey is an enum for search keys.
type SearchKey int

const (
	// UnknownKey is a fallback for unknown search keys.
	UnknownKey SearchKey = iota

	// GoKey is a key for Go packages.
	GoKey

	// NPMKey is a key for NPM packages.
	NPMKey

	// PyPIKey is a key for PyPI packages.
	PyPIKey

	// DNSKey is a key for DNS records.
	DNSKey

	// EmailKey is a key for email records.
	EmailKey
)

func (k SearchKey) String() string {
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

// SearchRecord is an union type for all search record values.
type SearchRecord interface{ record() }

// GoPackage is a struct for Go package search results.
type GoPackage struct {
	Name        string // Package name
	Path        string // Fully qualified package path
	Description string // Package description
}

func (*GoPackage) record() {}

// NPMPackage is a struct for NPM package search results.
type NPMPackage struct {
	Name        string // Package name
	Description string // Package description
}

func (*NPMPackage) record() {}

// PyPIPackage is a struct for PyPI package search results.
type PyPIPackage struct {
	Name        string // Package name
	Author      string // Package author
	Description string // Package description
}

func (*PyPIPackage) record() {}

// DNSRecord is a struct for DNS search results.
type DNSRecord struct {
	FQDN   string   // Fully qualified domain name
	IPList []net.IP // Associated IP addresses
}

func (*DNSRecord) record() {}

// EmailRecord is a struct for email search results.
type EmailRecord struct {
	Addr           string // Email address
	HasValidSyntax bool   // Email address has valid syntax
	HasValidDomain bool   // Email address has valid domain
}

func (*EmailRecord) record() {}

// SearchResult is a collection of search records with the original key.
type SearchResult struct {
	Key     SearchKey
	Records []SearchRecord
}

// SearchRender is a collection of search results with a rendered label.
type SearchRender struct {
	Label  string         `json:"label"`
	Result []SearchRecord `json:"result"`
}
