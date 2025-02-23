package dns

import (
	"fmt"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

// SearchByProbe searches for DNS records via nameserver lookups.
func SearchByProbe(name string, max int) ([]model.DNSRecord, error) {
	domains := []string{"com", "org", "net", "io", "tech", "ai", "me", "shop"}
	domainChan := make(chan string)

	go func() {
		for _, domain := range domains {
			domainChan <- fmt.Sprintf("%s.%s", name, domain)
		}
		close(domainChan)
	}()

	result := []model.DNSRecord{}
	errorCount := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go netWorker(domainChan, &wg, &mu, &result, &errorCount, max)
	}

	wg.Wait()
	if len(result) == 0 && errorCount > 0 {
		return result, fmt.Errorf("no results with %d errors", errorCount)
	}
	return result, nil
}
