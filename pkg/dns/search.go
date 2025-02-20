package dns

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/huangsam/namigo/internal/model"
)

func SearchByProbe(name string) ([]model.DNSResult, error) {
	var wg sync.WaitGroup
	domains := []string{"com", "org", "net", "io", "tech", "ai", "me", "shop"}
	domainChan := make(chan string)

	go func() {
		for _, domain := range domains {
			domainChan <- domain
		}
		close(domainChan)
	}()

	result := []model.DNSResult{}
	errorCount := 0

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
				result = append(result, model.DNSResult{FQDN: fullDomain, IPList: ips})
			}
		}()
	}

	wg.Wait()
	if errorCount == len(domains) {
		return result, errors.New("all DNS queries failed")
	}
	return result, nil
}
