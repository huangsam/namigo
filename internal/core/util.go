// Package core provides utility functions for the application.
package core

// dismiss ignores the error from a function/method
func dismiss(f func() error) {
	_ = f()
}

// IsValidDomainName validates a domain name for DNS lookups.
func IsValidDomainName(name string) bool {
	if len(name) == 0 || len(name) > 63 {
		return false
	}

	// Check each character
	for i, char := range name {
		// Must be alphanumeric or hyphen
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') &&
			(char < '0' || char > '9') && char != '-' {
			return false
		}
		// Cannot start or end with hyphen
		if char == '-' && (i == 0 || i == len(name)-1) {
			return false
		}
	}

	return true
}
