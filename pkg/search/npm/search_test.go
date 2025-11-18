package npm_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search/npm"
)

func TestSearchByAPI(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		json := `{
			"total": 2,
			"objects": [
				{
					"package": {
						"name": "test-package",
						"description": "A test package"
					}
				},
				{
					"package": {
						"name": "another-package",
						"description": "Another package"
					}
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(json))
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}
	results, err := npm.SearchByAPIWithBuilder(client, builder)
	if err != nil {
		t.Fatalf("SearchByAPI failed: %v", err)
	}

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}

	if results[0].Name != "test-package" {
		t.Errorf("expected name 'test-package', got '%s'", results[0].Name)
	}

	if results[0].Description != "A test package" {
		t.Errorf("expected description 'A test package', got '%s'", results[0].Description)
	}
}

func TestSearchByAPIWithBuilder(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		json := `{
			"total": 1,
			"objects": [
				{
					"package": {
						"name": "test-package",
						"description": "A test package"
					}
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(json))
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}

	results, err := npm.SearchByAPIWithBuilder(client, builder)
	if err != nil {
		t.Fatalf("SearchByAPIWithBuilder failed: %v", err)
	}

	if len(results) != 1 {
		t.Errorf("expected 1 result, got %d", len(results))
	}

	expected := model.NPMPackage{
		Name:        "test-package",
		Description: "A test package",
	}

	if results[0] != expected {
		t.Errorf("expected %+v, got %+v", expected, results[0])
	}
}
