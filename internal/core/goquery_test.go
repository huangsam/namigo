package core

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDocumentPipeline_Execute(t *testing.T) {
	html := `<html><body><h1>Test</h1></body></html>`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(html))
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}

	pipeline := NewDocumentPipeline(client, builder)
	doc, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	title := doc.Find("h1").Text()
	if title != "Test" {
		t.Errorf("expected 'Test', got '%s'", title)
	}
}

func TestDocumentPipeline_Execute_WithHandlers(t *testing.T) {
	html := `<html><body><h1>Test</h1></body></html>`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(html))
	}))
	defer server.Close()

	client := server.Client()
	builder := func() (*http.Request, error) {
		return http.NewRequest("GET", server.URL, nil)
	}

	handlerCalled := false
	handler := func(_ *http.Response) error {
		handlerCalled = true
		return nil
	}

	pipeline := NewDocumentPipeline(client, builder, handler)
	_, err := pipeline.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !handlerCalled {
		t.Error("handler was not called")
	}
}
