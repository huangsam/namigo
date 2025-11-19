package dns

import (
	"fmt"
	"net"
	"sync"

	"github.com/huangsam/namigo/v2/internal/core"
	"github.com/huangsam/namigo/v2/internal/model"
)

// LookupIPFunc is a function type for IP lookup.
type LookupIPFunc func(string) ([]net.IP, error)

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

// SearchByProbe searches for DNS records via nameserver lookups.
func SearchByProbe(name string, size int) ([]model.DNSRecord, error) {
	if !core.IsValidDomainName(name) {
		return nil, fmt.Errorf("invalid domain name: %s", name)
	}
	return SearchByProbeWithLookup(name, size, net.LookupIP)
}

// SearchByProbeWithLookup searches for DNS records using a custom lookup function.
func SearchByProbeWithLookup(name string, size int, lookup LookupIPFunc) ([]model.DNSRecord, error) {
	domains := []string{"com", "org", "net", "io", "tech", "ai", "me", "shop"}
	domainChan := make(chan string)

	go func() {
		for _, domain := range domains {
			domainChan <- fmt.Sprintf("%s.%s", name, domain)
		}
		close(domainChan)
	}()

	result := make([]model.DNSRecord, 0, size) // Pre-allocate with capacity
	errors := []error{}
	var mu sync.Mutex

	core.StartCommonWorkers(func() {
		for domain := range domainChan {
			ips, err := lookup(domain)
			if err != nil {
				mu.Lock() // Critical section
				errors = append(errors, err)
				mu.Unlock()
			}

			mu.Lock() // Critical section
			if len(result) < size {
				result = append(result, model.DNSRecord{FQDN: domain, IPList: ips})
			}
			mu.Unlock()
			if len(result) >= size {
				break
			}
		}
	})

	if len(result) == 0 && len(errors) > 0 {
		return result, fmt.Errorf("no results with %d errors", len(errors))
	}
	return result, nil
}
