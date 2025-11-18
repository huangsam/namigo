package golang_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search/golang"
)

func TestSearchByScrapeMain(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		html := `
		<html>
		<body>
			<div class="SearchSnippet">
				<h2>testpackage (github.com/test/testpackage)</h2>
				<p>A test package for testing</p>
			</div>
			<div class="SearchSnippet">
				<h2>otherpackage (github.com/other/otherpackage)</h2>
				<p>Another package</p>
			</div>
		</body>
		</html>`
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(html))
	}))
	defer server.Close()

	client := server.Client()
	builder := golang.ScrapeListWithBaseURL("test", server.URL)
	results, err := golang.SearchByScrapeWithBuilder(client, "test", 5, builder)
	if err != nil {
		t.Fatalf("SearchByScrape failed: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}

	if results[0].Name != "testpackage" {
		t.Errorf("expected name 'testpackage', got '%s'", results[0].Name)
	}

	if results[0].Path != "github.com/test/testpackage" {
		t.Errorf("expected path 'github.com/test/testpackage', got '%s'", results[0].Path)
	}
}

func TestSearchByScrapeWithBuilder(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		html := `
		<html>
		<body>
			<div class="SearchSnippet">
				<h2>testpackage (github.com/test/testpackage)</h2>
				<p>A test package for testing</p>
			</div>
		</body>
		</html>`
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(html))
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}

	results, err := golang.SearchByScrapeWithBuilder(client, "test", 5, builder)
	if err != nil {
		t.Fatalf("SearchByScrapeWithBuilder failed: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}

	expected := model.GoPackage{
		Name:        "testpackage",
		Path:        "github.com/test/testpackage",
		Description: "A test package for testing",
	}

	if results[0] != expected {
		t.Errorf("expected %+v, got %+v", expected, results[0])
	}
}
