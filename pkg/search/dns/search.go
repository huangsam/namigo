package dns

import (
	"fmt"
	"net"
	"sync"

	"github.com/huangsam/namigo/internal/core"
	"github.com/huangsam/namigo/internal/model"
)

// LookupIPFunc is a function type for IP lookup.
type LookupIPFunc func(string) ([]net.IP, error)

// SearchByProbe searches for DNS records via nameserver lookups.
func SearchByProbe(name string, size int) ([]model.DNSRecord, error) {
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

	result := []model.DNSRecord{}
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
