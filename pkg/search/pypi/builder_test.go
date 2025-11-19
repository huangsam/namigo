package pypi_test

import (
	"net/http"
	"testing"

	"github.com/huangsam/namigo/v2/pkg/search/pypi"
)

func TestAPIList(t *testing.T) {
	builder := pypi.APIList()
	req, err := builder()
	if err != nil {
		t.Fatalf("APIList() error = %v", err)
	}

	expectedURL := "https://pypi.org/simple/"
	if req.URL.String() != expectedURL {
		t.Errorf("APIList() URL = %v, want %v", req.URL.String(), expectedURL)
	}

	expectedHeader := "application/vnd.pypi.simple.v1+json"
	if req.Header.Get("Accept") != expectedHeader {
		t.Errorf("APIList() Accept header = %v, want %v", req.Header.Get("Accept"), expectedHeader)
	}

	if req.Method != http.MethodGet {
		t.Errorf("APIList() method = %v, want %v", req.Method, http.MethodGet)
	}
}

func TestAPIDetail(t *testing.T) {
	pkg := "example-package"
	builder := pypi.APIDetail(pkg)
	req, err := builder()
	if err != nil {
		t.Fatalf("APIDetail() error = %v", err)
	}

	expectedURL := "https://pypi.org/pypi/" + pkg + "/json"
	if req.URL.String() != expectedURL {
		t.Errorf("APIDetail() URL = %v, want %v", req.URL.String(), expectedURL)
	}

	if req.Method != http.MethodGet {
		t.Errorf("APIDetail() method = %v, want %v", req.Method, http.MethodGet)
	}
}
