package sub

import (
	"encoding/json"
	"fmt"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
)

// displayResults prints results.
func displayResults(results any, formatter search.ResultFormatter, format model.OutputFormat) {
	switch res := results.(type) {
	case []model.GoPackage:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, format)
		}
	case []model.NPMPackage:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, format)
		}
	case []model.PyPIPackage:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, format)
		}
	case []model.DNSRecord:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, format)
		}
	case []model.EmailRecord:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, format)
		}
	}
}

// displayResultsHelper prints results by data type and output mode.
func displayResultsHelper[T any](results []T, formatter search.ResultFormatter, format model.OutputFormat) {
	switch format {
	case model.JSONFormat:
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
	case model.TextFormat:
		for _, r := range results {
			fmt.Println(formatter.Format(r))
		}
	}
}

// checkPortfolio checks for any errors.
func checkPortfolio(ptf *search.Portfolio) error {
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
