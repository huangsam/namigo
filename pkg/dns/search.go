package dns

import (
	"fmt"
	"net"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

func SearchByProbe(name string, max int) ([]model.DNSResult, error) {
	domains := []string{"com", "org", "net", "io", "tech", "ai", "me", "shop"}
	domainChan := make(chan string)

	go func() {
		for _, domain := range domains {
			domainChan <- domain
		}
		close(domainChan)
	}()

	result := []model.DNSResult{}
	resultCount := 0
	errorCount := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for domain := range domainChan {
				fullDomain := fmt.Sprintf("%s.%s", name, domain)
				ips, err := net.LookupIP(fullDomain)
				if err != nil {
					errorCount++
				}
				mu.Lock()
				if resultCount < max {
					result = append(result, model.DNSResult{FQDN: fullDomain, IPList: ips})
					resultCount++
				}
				mu.Unlock()
				if len(result) >= max {
					return
				}
			}
		}()
	}

	wg.Wait()
	if errorCount == len(domains) {
		return result, fmt.Errorf("queries failed for %v", domains)
	}
	return result, nil
}
