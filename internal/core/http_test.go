package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRESTAPIQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test response"))
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}

	data, err := RESTAPIQuery(client, builder)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(data) != "test response" {
		t.Errorf("expected 'test response', got '%s'", string(data))
	}
}

func TestRESTAPIQuery_BadStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}

	_, err := RESTAPIQuery(client, builder)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
