package email

import (
	"net"
	"testing"

	emailverifier "github.com/AfterShip/email-verifier"
)

// mockVerifier implements Verifier for testing.
type mockVerifier struct {
	result *emailverifier.Result
	err    error
}

func (m *mockVerifier) Verify(_ string) (*emailverifier.Result, error) {
	return m.result, m.err
}

func TestSearchByProbeWithDeps(t *testing.T) {
	v := &mockVerifier{
		result: &emailverifier.Result{
			Syntax: emailverifier.Syntax{Valid: true},
		},
	}
	lookup := func(_ string) ([]*net.MX, error) {
		return []*net.MX{{Host: "mx.example.com"}}, nil
	}

	result, err := SearchByProbeWithDeps("test", 1, v, lookup)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(result) != 1 {
		t.Errorf("expected 1 result, got %d", len(result))
	}
	if !result[0].HasValidSyntax {
		t.Error("expected valid syntax")
	}
	if !result[0].HasValidDomain {
		t.Error("expected valid domain")
	}
}
