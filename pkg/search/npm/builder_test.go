package npm_test

import (
	"net/http"
	"net/http/httptest"
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

func TestSearchByAPIWithClient(t *testing.T) {
	jsonResponse := `{
		"objects": [
			{
				"package": {
					"name": "test-package",
					"description": "A test package"
				}
			}
		]
	}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(jsonResponse))
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}

	result, err := npm.SearchByAPIWithBuilder(client, builder)
	if err != nil {
		t.Fatalf("SearchByAPIWithBuilder() error = %v", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 result, got %d", len(result))
	}
	if result[0].Name != "test-package" {
		t.Errorf("expected name 'test-package', got '%s'", result[0].Name)
	}
}
