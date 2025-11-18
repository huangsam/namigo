package search

import (
	"errors"
	"testing"

	"github.com/huangsam/namigo/internal/model"
)

func TestNewSearchPortfolio(t *testing.T) {
	p := NewSearchPortfolio(TextOption)
	if p.option != TextOption {
		t.Errorf("expected TextOption, got %v", p.option)
	}
	if len(p.lineMap) != 5 {
		t.Errorf("expected 5 lineMap entries, got %d", len(p.lineMap))
	}
}

func TestPortfolio_Register(t *testing.T) {
	p := NewSearchPortfolio(TextOption)
	initialLen := len(p.funcs)
	p.Register(func() (model.SearchResult, error) { return model.SearchResult{}, nil })
	if len(p.funcs) != initialLen+1 {
		t.Errorf("expected funcs length to increase")
	}
}

func TestPortfolio_Run_Success(t *testing.T) {
	p := NewSearchPortfolio(TextOption)
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
	p := NewSearchPortfolio(TextOption)
	p.Register(func() (model.SearchResult, error) {
		return model.SearchResult{}, errors.New("test error")
	})
	err := p.Run()
	if err != ErrPorftolioFailure {
		t.Errorf("expected ErrPorftolioFailure, got %v", err)
	}
}

func TestPortfolio_Run_Empty(t *testing.T) {
	p := NewSearchPortfolio(TextOption)
	err := p.Run()
	if err != ErrPorftolioEmpty {
		t.Errorf("expected ErrPorftolioEmpty, got %v", err)
	}
}
