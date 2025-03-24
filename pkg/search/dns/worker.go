package dns

import (
	"net"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

// netWorker runs concurrent logic for DNS search.
func netWorker(
	domainChan chan string,
	mu *sync.Mutex,
	result *[]model.DNSRecord,
	errors *[]error,
	maxResults int,
) {
	for domain := range domainChan {
		ips, err := net.LookupIP(domain)
		if err != nil {
			mu.Lock() // Critical section
			*errors = append(*errors, err)
			mu.Unlock()
		}

		mu.Lock() // Critical section
		if len(*result) < maxResults {
			*result = append(*result, model.DNSRecord{FQDN: domain, IPList: ips})
		}
		mu.Unlock()
		if len(*result) >= maxResults {
			break
		}
	}
}
