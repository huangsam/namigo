package search

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/huangsam/namigo/v2/internal/model"
)

func TestNewSearchPortfolio(t *testing.T) {
	var buf bytes.Buffer
	p := NewSearchPortfolio(TextOption, &buf)
	if p.option != TextOption {
		t.Errorf("expected TextOption, got %v", p.option)
	}
	if len(p.lineMap) != 5 {
		t.Errorf("expected 5 lineMap entries, got %d", len(p.lineMap))
	}
}

func TestPortfolio_Register(t *testing.T) {
	var buf bytes.Buffer
	p := NewSearchPortfolio(TextOption, &buf)
	initialLen := len(p.funcs)
	p.Register(func() (model.SearchResult, error) { return model.SearchResult{}, nil })
	if len(p.funcs) != initialLen+1 {
		t.Errorf("expected funcs length to increase")
	}
}

func TestPortfolio_Run_Success(t *testing.T) {
	var buf bytes.Buffer
	p := NewSearchPortfolio(TextOption, &buf)
	p.Register(func() (model.SearchResult, error) {
		return model.SearchResult{Key: model.GoKey, Records: []model.SearchRecord{&model.GoPackage{Name: "test"}}}, nil
	})
	err := p.Run()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(p.resultMap) != 1 {
		t.Errorf("expected 1 result, got %d", len(p.resultMap))
	}
}

func TestPortfolio_Run_Error(t *testing.T) {
	var buf bytes.Buffer
	p := NewSearchPortfolio(TextOption, &buf)
	p.Register(func() (model.SearchResult, error) {
		return model.SearchResult{}, errors.New("test error")
	})
	err := p.Run()
	if err != ErrPorftolioFailure {
		t.Errorf("expected ErrPorftolioFailure, got %v", err)
	}
}

func TestPortfolio_Run_Empty(t *testing.T) {
	var buf bytes.Buffer
	p := NewSearchPortfolio(TextOption, &buf)
	err := p.Run()
	if err != ErrPorftolioEmpty {
		t.Errorf("expected ErrPorftolioEmpty, got %v", err)
	}
}

func TestPortfolio_Display_Text(t *testing.T) {
	var buf bytes.Buffer
	p := NewSearchPortfolio(TextOption, &buf)

	// Add some test data
	p.resultMap = map[model.SearchKey][]model.SearchRecord{
		model.GoKey: {&model.GoPackage{Name: "test", Path: "github.com/test", Description: "A test package"}},
	}

	p.Display()

	output := buf.String()
	if !strings.Contains(output, "🍺 Prepare PlainText results") {
		t.Errorf("expected display header, got: %s", output)
	}
	if !strings.Contains(output, "test") {
		t.Errorf("expected package name in output, got: %s", output)
	}
}

func TestPortfolio_Display_JSON(t *testing.T) {
	var buf bytes.Buffer
	p := NewSearchPortfolio(JSONOption, &buf)

	// Add some test data
	p.resultMap = map[model.SearchKey][]model.SearchRecord{
		model.GoKey: {&model.GoPackage{Name: "test", Path: "github.com/test", Description: "A test package"}},
	}

	p.Display()

	output := buf.String()
	if !strings.Contains(output, "🍺 Prepare JSON results") {
		t.Errorf("expected display header, got: %s", output)
	}
	if !strings.Contains(output, `"label": "Golang"`) {
		t.Errorf("expected JSON label in output, got: %s", output)
	}
}
