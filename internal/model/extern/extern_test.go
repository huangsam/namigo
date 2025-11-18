package extern

import (
	"encoding/json"
	"testing"
)

func TestNPMAPIListResponse_Unmarshal(t *testing.T) {
	jsonData := `{
		"objects": [
			{
				"package": {
					"name": "test-package",
					"description": "A test package"
				}
			}
		],
		"total": 1
	}`

	var resp NPMAPIListResponse
	err := json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal NPMAPIListResponse: %v", err)
	}

	if len(resp.Objects) != 1 {
		t.Errorf("Expected 1 object, got %d", len(resp.Objects))
	}

	if resp.Objects[0].Package.Name != "test-package" {
		t.Errorf("Expected name 'test-package', got '%s'", resp.Objects[0].Package.Name)
	}

	if resp.Total != 1 {
		t.Errorf("Expected total 1, got %d", resp.Total)
	}
}

func TestPyPIAPIListResponse_Unmarshal(t *testing.T) {
	jsonData := `{
		"projects": [
			{
				"name": "test-pypi"
			}
		]
	}`

	var resp PyPIAPIListResponse
	err := json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal PyPIAPIListResponse: %v", err)
	}

	if len(resp.Projects) != 1 {
		t.Errorf("Expected 1 project, got %d", len(resp.Projects))
	}

	if resp.Projects[0].Name != "test-pypi" {
		t.Errorf("Expected name 'test-pypi', got '%s'", resp.Projects[0].Name)
	}
}

func TestPyPIAPIDetailResponse_Unmarshal(t *testing.T) {
	jsonData := `{
		"info": {
			"author": "Test Author",
			"summary": "A summary",
			"description": "A description",
			"version": "1.0.0"
		}
	}`

	var resp PyPIAPIDetailResponse
	err := json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		t.Fatalf("Failed to unmarshal PyPIAPIDetailResponse: %v", err)
	}

	if resp.Info.Author != "Test Author" {
		t.Errorf("Expected author 'Test Author', got '%s'", resp.Info.Author)
	}

	if resp.Info.Summary != "A summary" {
		t.Errorf("Expected summary 'A summary', got '%s'", resp.Info.Summary)
	}
}
