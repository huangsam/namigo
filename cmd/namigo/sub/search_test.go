package sub_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/huangsam/namigo/cmd/namigo/sub"
	"github.com/huangsam/namigo/pkg/search"
	"github.com/urfave/cli/v3"
)

func createMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.Path, "/search") && r.URL.Query().Get("q") == "test":
			// Mock pkg.go.dev search response
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`
				<div class="SearchSnippet">
					<h2>testing (testing)</h2>
					<p>Package testing provides support for automated testing of Go packages.</p>
				</div>
				<div class="SearchSnippet">
					<h2>require (github.com/stretchr/testify/require)</h2>
					<p>Package require implements the same assertions as the assert package but stops test execution.</p>
				</div>
			`))
		case strings.Contains(r.URL.Path, "/-/v1/search") && r.URL.Query().Get("text") == "test":
			// Mock npm registry response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"objects": [
					{
						"package": {
							"name": "test-package",
							"description": "A test package for npm"
						}
					},
					{
						"package": {
							"name": "testing-lib",
							"description": "Library for testing"
						}
					}
				]
			}`))
		case r.URL.Path == "/simple/":
			// Mock PyPI simple API response
			w.Header().Set("Content-Type", "application/vnd.pypi.simple.v1+json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"projects": [
					{"name": "test-package"},
					{"name": "testing-utils"}
				]
			}`))
		case strings.HasPrefix(r.URL.Path, "/pypi/test-package/json"):
			// Mock PyPI detail response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"info": {
					"name": "test-package",
					"author": "Test Author",
					"summary": "A test package"
				}
			}`))
		case strings.HasPrefix(r.URL.Path, "/pypi/testing-utils/json"):
			// Mock PyPI detail response for second package
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"info": {
					"name": "testing-utils",
					"author": "Test Author 2",
					"summary": "Testing utilities"
				}
			}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
}

type mockTransport struct {
	serverURL string
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Rewrite the URL to point to our test server
	req.URL.Scheme = "http"
	req.URL.Host = strings.TrimPrefix(t.serverURL, "http://")
	return http.DefaultTransport.RoundTrip(req)
}

func TestSearchRunner_RunPackageSearch(t *testing.T) {
	tests := []struct {
		name         string
		searchTerm   string
		maxSize      int
		outputFormat search.FormatOption
	}{
		{
			name:         "Valid search with text format",
			searchTerm:   "test",
			maxSize:      5,
			outputFormat: search.TextOption,
		},
		{
			name:         "Valid search with JSON format",
			searchTerm:   "test",
			maxSize:      5,
			outputFormat: search.JSONOption,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := createMockServer()
			defer mockServer.Close()

			mockClient := &http.Client{
				Transport: &mockTransport{serverURL: mockServer.URL},
				Timeout:   10 * time.Second,
			}
			var buf bytes.Buffer
			runner := sub.NewSearchRunnerWithClient(&buf, mockClient)

			err := runner.RunPackageSearch(tt.searchTerm, tt.maxSize, tt.outputFormat)

			output := buf.String()
			// Verify that search indicators were printed (output was redirected)
			if !strings.Contains(output, "🔍 Search for") {
				t.Errorf("RunPackageSearch() output missing search indicator, got: %v", output)
			}

			// Should not have network errors with mock server
			if err != nil {
				t.Errorf("RunPackageSearch() returned unexpected error: %v", err)
			}

			// Verify we got some results
			if !strings.Contains(output, "🍺 Prepare") {
				t.Errorf("RunPackageSearch() output missing display message, got: %v", output)
			}
		})
	}
}

// TestSearchRunner_RunDNSSearch removed - integration testing moved to unit tests
// Use pkg/search/dns tests for DNS search functionality testing

// TestSearchRunner_RunEmailSearch removed - integration testing moved to unit tests
// Use pkg/search/email tests for email search functionality testing

func TestSearchPackageAction(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Missing search term",
			args:    []string{},
			wantErr: true,
			errMsg:  "missing search term",
		},
		// Removed integration test that makes real network calls
		// Use TestSearchRunner_RunPackageSearch for functional testing with mocks
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cli.Command{
				Commands: []*cli.Command{
					{
						Name:   "search",
						Action: sub.SearchPackageAction,
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "size", Value: 5},
							&cli.StringFlag{Name: "format", Value: "text"},
						},
					},
				},
			}

			args := append([]string{"namigo", "search"}, tt.args...)
			err := cmd.Run(context.Background(), args)

			if (err != nil) != tt.wantErr {
				t.Errorf("SearchPackageAction() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("SearchPackageAction() error message = %v, want containing %v", err.Error(), tt.errMsg)
				}
			}
		})
	}
}

func TestNewSearchRunnerWithClient(t *testing.T) {
	var buf bytes.Buffer
	httpClient := &http.Client{}
	runner := sub.NewSearchRunnerWithClient(&buf, httpClient)

	if runner == nil {
		t.Errorf("expected non-nil runner")
	}
}

func TestSearchDNSAction(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Missing search term",
			args:    []string{},
			wantErr: true,
			errMsg:  "missing search term",
		},
		// Removed integration test that makes real DNS lookups
		// Use TestSearchRunner_RunDNSSearch for functional testing
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cli.Command{
				Commands: []*cli.Command{
					{
						Name:   "dns",
						Action: sub.SearchDNSAction,
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "size", Value: 5},
							&cli.StringFlag{Name: "format", Value: "text"},
						},
					},
				},
			}

			args := append([]string{"namigo", "dns"}, tt.args...)
			err := cmd.Run(context.Background(), args)

			if (err != nil) != tt.wantErr {
				t.Errorf("SearchDNSAction() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("SearchDNSAction() error message = %v, want containing %v", err.Error(), tt.errMsg)
				}
			}
		})
	}
}

func TestSearchEmailAction(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Missing search term",
			args:    []string{},
			wantErr: true,
			errMsg:  "missing search term",
		},
		// Removed integration test that makes real email validation calls
		// Use TestSearchRunner_RunEmailSearch for functional testing
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cli.Command{
				Commands: []*cli.Command{
					{
						Name:   "email",
						Action: sub.SearchEmailAction,
						Flags: []cli.Flag{
							&cli.IntFlag{Name: "size", Value: 5},
							&cli.StringFlag{Name: "format", Value: "text"},
						},
					},
				},
			}

			args := append([]string{"namigo", "email"}, tt.args...)
			err := cmd.Run(context.Background(), args)

			if (err != nil) != tt.wantErr {
				t.Errorf("SearchEmailAction() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errMsg) {
					t.Errorf("SearchEmailAction() error message = %v, want containing %v", err.Error(), tt.errMsg)
				}
			}
		})
	}
}
