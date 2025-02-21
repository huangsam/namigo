package dns

import (
	"net"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

// worker is concurrent logic for DNS search.
func worker(
	domainChan chan string,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
	result *[]model.DNSResult,
	errorCount *int,
	maxResults int,
) {
	defer wg.Done()
	for domain := range domainChan {
		ips, err := net.LookupIP(domain)
		if err != nil {
			*errorCount++
		}
		mu.Lock()
		if len(*result) < maxResults {
			*result = append(*result, model.DNSResult{FQDN: domain, IPList: ips})
		}
		mu.Unlock()
		if len(*result) >= maxResults {
			return
		}
	}
}
