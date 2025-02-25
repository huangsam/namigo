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

// SearchByProbe searches for email records via nameserver lookups.
//
// Please note that end-to-end email validation would involve APIs with pricing
// tiers like Abstract does; their free plan limits requests to 1 per second.
// To learn more:
//
// https://docs.abstractapi.com/email-validation
func SearchByProbe(name string, size int) ([]model.EmailRecord, error) {
	result := []model.EmailRecord{}
	for _, domain := range domains {
		if len(result) >= size {
			break
		}
		email := fmt.Sprintf("%s@%s", name, domain)

		var hasValidSyntax bool
		verifyRes, err := verifier.Verify(email)
		if err == nil && verifyRes.Syntax.Valid {
			hasValidSyntax = true
		}

		var hasValidDomain bool
		mxRecords, err := net.LookupMX(domain)
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
