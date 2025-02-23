package search

import (
	"errors"
	"testing"

	"github.com/huangsam/namigo/internal/model"
)

func TestPortfolioErrors(t *testing.T) {
	tests := []struct {
		name      string
		portfolio Portfolio
		wantErrs  []error
	}{
		{
			name:      "No errors and no results",
			portfolio: Portfolio{},
			wantErrs:  []error{},
		},
		{
			name: "Golang error",
			portfolio: Portfolio{
				Errs: PortfolioErrors{
					Golang: errors.New("golang error"),
				},
			},
			wantErrs: []error{errors.New("golang error")},
		},
		{
			name: "NPM error",
			portfolio: Portfolio{
				Errs: PortfolioErrors{
					NPM: errors.New("npm error"),
				},
			},
			wantErrs: []error{errors.New("npm error")},
		},
		{
			name: "PyPI error",
			portfolio: Portfolio{
				Errs: PortfolioErrors{
					PyPI: errors.New("pypi error"),
				},
			},
			wantErrs: []error{errors.New("pypi error")},
		},
		{
			name: "DNS error",
			portfolio: Portfolio{
				Errs: PortfolioErrors{
					DNS: errors.New("dns error"),
				},
			},
			wantErrs: []error{errors.New("dns error")},
		},
		{
			name: "Multiple errors",
			portfolio: Portfolio{
				Errs: PortfolioErrors{
					Golang: errors.New("golang error"),
					NPM:    errors.New("npm error"),
					PyPI:   errors.New("pypi error"),
					DNS:    errors.New("dns error"),
				},
			},
			wantErrs: []error{
				errors.New("golang error"),
				errors.New("npm error"),
				errors.New("pypi error"),
				errors.New("dns error"),
			},
		},
		{
			name: "No errors with results",
			portfolio: Portfolio{
				Results: PortfolioResults{
					Golang: []model.GoPackage{
						{}, {}, {}, // Some fake results
					},
				},
			},
			wantErrs: []error{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErrs := tt.portfolio.Errors(); !equalErrors(gotErrs, tt.wantErrs) {
				t.Errorf("Portfolio.Errors() = %v, want %v", gotErrs, tt.wantErrs)
			}
		})
	}
}

func equalErrors(a, b []error) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Error() != b[i].Error() {
			return false
		}
	}
	return true
}
