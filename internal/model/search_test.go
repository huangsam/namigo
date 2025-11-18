package model

import "testing"

func TestSearchKey_String(t *testing.T) {
	tests := []struct {
		key      SearchKey
		expected string
	}{
		{UnknownKey, "Unknown"},
		{GoKey, "Golang"},
		{NPMKey, "NPM"},
		{PyPIKey, "PyPI"},
		{DNSKey, "DNS"},
		{EmailKey, "Email"},
	}

	for _, test := range tests {
		if result := test.key.String(); result != test.expected {
			t.Errorf("expected %s, got %s", test.expected, result)
		}
	}
}
