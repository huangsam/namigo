package golang_test

import (
	"net/http"
	"net/http/httptest"
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

func TestSearchByScrape(t *testing.T) {
	html := `
	<html>
	<body>
		<div class="SearchSnippet">
			<h2>testpackage (github.com/test/testpackage)</h2>
			<p>A test package</p>
		</div>
	</body>
	</html>`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(html))
	}))
	defer server.Close()

	builder := golang.ScrapeListWithBaseURL("test", server.URL)
	client := server.Client()
	result, err := golang.SearchByScrapeWithBuilder(client, "test", 1, builder)
	if err != nil {
		t.Fatalf("SearchByScrapeWithBuilder() error = %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 result, got %d", len(result))
	}
	if result[0].Name != "testpackage" {
		t.Errorf("expected name 'testpackage', got '%s'", result[0].Name)
	}
}
