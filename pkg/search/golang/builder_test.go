package golang_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/huangsam/namigo/pkg/search/golang"
)

func TestScrapeList(t *testing.T) {
	name := "example-package"
	builder := golang.ScrapeList(name)
	req, err := builder()
	if err != nil {
		t.Fatalf("ScrapeList() error = %v", err)
	}

	expectedPrefix := "https://pkg.go.dev/search"
	if !strings.HasPrefix(req.URL.String(), expectedPrefix) {
		t.Errorf("ScrapeList() URL prefix = %v, want %v", req.URL.String(), expectedPrefix)
	}

	expectedParam := "q=" + name
	if !strings.Contains(req.URL.RawQuery, expectedParam) {
		t.Errorf("ScrapeList() query parameter missing = %v, want %v", req.URL.RawQuery, expectedParam)
	}

	if req.Method != http.MethodGet {
		t.Errorf("ScrapeList() method = %v, want %v", req.Method, http.MethodGet)
	}
}
