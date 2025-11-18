package dns

import (
	"net"
	"testing"
)

func TestSearchByProbeWithLookup(t *testing.T) {
	lookup := func(domain string) ([]net.IP, error) {
		if domain == "test.com" {
			return []net.IP{net.ParseIP("1.2.3.4")}, nil
		}
		return nil, nil
	}

	result, err := SearchByProbeWithLookup("test", 10, lookup) // Increase size to get all
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	found := false
	for _, r := range result {
		if r.FQDN == "test.com" && len(r.IPList) > 0 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find test.com with IPs")
	}
}
