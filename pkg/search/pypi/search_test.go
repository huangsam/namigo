package pypi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/internal/model/extern"
)

// testTransport rewrites request URLs to point to the test server.
type testTransport struct {
	baseURL string
}

func (t *testTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Parse the base URL
	base, err := url.Parse(t.baseURL)
	if err != nil {
		return nil, err
	}
	// Rewrite the request URL to use the test server's scheme and host
	req.URL.Scheme = base.Scheme
	req.URL.Host = base.Host
	return http.DefaultTransport.RoundTrip(req)
}

func TestSearchByAPI(t *testing.T) {
	// Mock server responses
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/simple/":
			// Return list of projects
			listResp := extern.PyPIAPIListResponse{
				Projects: []struct {
					LastSerial int    `json:"_last-serial"`
					Name       string `json:"name"`
				}{
					{Name: "test-package"},
					{Name: "another-package"},
					{Name: "test-lib"},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(listResp); err != nil {
				t.Errorf("failed to encode response: %v", err)
			}
		default:
			if strings.HasPrefix(r.URL.Path, "/pypi/") && strings.HasSuffix(r.URL.Path, "/json") {
				// Extract package name from /pypi/{pkg}/json
				pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
				if len(pathParts) >= 2 {
					pkg := pathParts[1]
					detailResp := extern.PyPIAPIDetailResponse{
						Info: struct {
							Author      string `json:"author"`
							Description string `json:"description"`
							Summary     string `json:"summary"`
							Version     string `json:"version"`
						}{
							Summary: "A test package",
							Author:  "Test Author",
						},
					}
					if pkg == "another-package" {
						detailResp.Info.Summary = ""
						detailResp.Info.Author = ""
					}
					w.Header().Set("Content-Type", "application/json")
					if err := json.NewEncoder(w).Encode(detailResp); err != nil {
						t.Errorf("failed to encode response: %v", err)
					}
				} else {
					w.WriteHeader(http.StatusNotFound)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}
	}))
	defer server.Close()

	// Create client with transport that rewrites URLs
	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: &testTransport{baseURL: server.URL},
	}

	results, err := SearchByAPIWithClient(client, "test", 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := map[string]model.PyPIPackage{
		"test-package": {Name: "test-package", Description: "A test package", Author: "Test Author"},
		"test-lib":     {Name: "test-lib", Description: "A test package", Author: "Test Author"},
	}

	if len(results) != len(expected) {
		t.Errorf("expected %d results, got %d", len(expected), len(results))
	}

	for _, pkg := range results {
		exp, ok := expected[pkg.Name]
		if !ok {
			t.Errorf("unexpected package %s", pkg.Name)
		}
		if pkg != exp {
			t.Errorf("expected %+v, got %+v", exp, pkg)
		}
	}
}

func TestSearchByAPI_NoResults(t *testing.T) {
	// Test with no matching packages
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/simple/" {
			listResp := extern.PyPIAPIListResponse{
				Projects: []struct {
					LastSerial int    `json:"_last-serial"`
					Name       string `json:"name"`
				}{
					{Name: "other-package"},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(listResp); err != nil {
				t.Errorf("failed to encode response: %v", err)
			}
		}
	}))
	defer server.Close()

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: &testTransport{baseURL: server.URL},
	}

	results, err := SearchByAPIWithClient(client, "test", 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(results) != 0 {
		t.Errorf("expected 0 results, got %d", len(results))
	}
}

func TestSearchByAPI_Error(t *testing.T) {
	// Test with server error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &http.Client{
		Timeout:   5 * time.Second,
		Transport: &testTransport{baseURL: server.URL},
	}

	_, err := SearchByAPIWithClient(client, "test", 2)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
