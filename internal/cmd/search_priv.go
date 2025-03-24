package cmd

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
)

var (
	ErrPorftolioEmpty   = errors.New("portfolio collection empty")
	ErrPorftolioFailure = errors.New("portfolio collection failure")
)

// DisplayResults displays results based on the specified parameters.
func DisplayResults[T any](results []T, formatter search.ResultFormatter, format model.FormatOption) {
	if len(results) == 0 {
		return
	}
	switch format {
	case model.JSONOption:
		type wrapper struct {
			Label   string `json:"label"`
			Results []T    `json:"results"`
		}
		label := formatter.Label()
		data, err := json.MarshalIndent(&wrapper{Label: label, Results: results}, "", "  ")
		if err != nil {
			fmt.Printf("Cannot print %s for %s: %v\n", format, label, err)
			return
		}
		fmt.Printf("%s\n", data)
	case model.TextOption:
		for _, r := range results {
			fmt.Println(formatter.Format(r))
		}
	}
}

// CheckPortfolio checks for any portfolio errors.
func CheckPortfolio(ptf *search.Portfolio) error {
	if errs := ptf.Errors(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("💀 Error: %s\n", err)
		}
		return ErrPorftolioFailure
	} else if ptf.Size() == 0 {
		return ErrPorftolioEmpty
	}
	return nil
}
