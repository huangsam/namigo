package email

import (
	"fmt"
	"net"
	"time"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/huangsam/namigo/internal/model"
)

const probeInterval = 100 * time.Millisecond // Adjust as needed

var (
	verifier = emailverifier.NewVerifier()
	domains  = []string{"gmail.com", "outlook.com", "yahoo.com"}
)

// Verifier is an interface for email verification.
type Verifier interface {
	Verify(email string) (*emailverifier.Result, error)
}

// LookupMXFunc is a function type for MX lookup.
type LookupMXFunc func(string) ([]*net.MX, error)

// SearchByProbe searches for email records via nameserver lookups.
//
// Please note that end-to-end email validation would involve APIs with pricing
// tiers like Abstract does; their free plan limits requests to 1 per second.
// To learn more:
//
// https://docs.abstractapi.com/email-validation
func SearchByProbe(name string, size int) ([]model.EmailRecord, error) {
	return SearchByProbeWithDeps(name, size, verifier, net.LookupMX)
}

// SearchByProbeWithDeps searches for email records using dependencies.
func SearchByProbeWithDeps(name string, size int, v Verifier, lookup LookupMXFunc) ([]model.EmailRecord, error) {
	result := make([]model.EmailRecord, 0, size) // Pre-allocate with capacity
	for _, domain := range domains {
		if len(result) >= size {
			break
		}
		email := fmt.Sprintf("%s@%s", name, domain)

		var hasValidSyntax bool
		verifyRes, err := v.Verify(email)
		if err == nil && verifyRes.Syntax.Valid {
			hasValidSyntax = true
		}

		var hasValidDomain bool
		mxRecords, err := lookup(domain)
		if err == nil && len(mxRecords) > 0 {
			hasValidDomain = true
		}

		emailRecord := model.EmailRecord{
			Addr:           email,
			HasValidSyntax: hasValidSyntax,
			HasValidDomain: hasValidDomain,
		}
		result = append(result, emailRecord)

		// Email infrastructure often enforces rate limiting, for
		// each entity request. Therefore, we apply a probe
		// interval so that we don't abuse ToS
		time.Sleep(probeInterval)
	}
	return result, nil
}
