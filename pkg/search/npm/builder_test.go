package npm_test

import (
	"net/http"
	"testing"

	"github.com/huangsam/namigo/pkg/search/npm"
)

func TestAPIList(t *testing.T) {
	name := "example-package"
	size := 10
	builder := npm.APIList(name, size)
	req, err := builder()
	if err != nil {
		t.Fatalf("APIList() error = %v", err)
	}

	expectedURL := "https://registry.npmjs.com/-/v1/search?size=10&text=example-package"
	if req.URL.String() != expectedURL {
		t.Errorf("APIList() URL = %v, want %v", req.URL.String(), expectedURL)
	}

	if req.Method != http.MethodGet {
		t.Errorf("APIList() method = %v, want %v", req.Method, http.MethodGet)
	}
}
