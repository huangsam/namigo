package sub

import (
	"encoding/json"
	"fmt"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
)

// displayResults prints results.
func displayResults(results any, formatter search.Formatter, mode model.OutputMode) {
	switch res := results.(type) {
	case []model.GoPackage:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, mode)
		}
	case []model.NPMPackage:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, mode)
		}
	case []model.PyPIPackage:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, mode)
		}
	case []model.DNSRecord:
		if len(res) > 0 {
			displayResultsHelper(res, formatter, mode)
		}
	}
}

// displayResultsHelper prints results by data type and output mode.
func displayResultsHelper[T any](results []T, formatter search.Formatter, mode model.OutputMode) {
	switch mode {
	case model.JSONMode:
		type wrapper struct {
			Label   string `json:"label"`
			Results []T    `json:"results"`
		}
		label := formatter.Label()
		data, err := json.MarshalIndent(&wrapper{Label: label, Results: results}, "", "  ")
		if err != nil {
			fmt.Printf("Cannot print %s for %s: %v\n", mode, label, err)
			return
		}
		fmt.Printf("%s\n", data)
	case model.TextMode:
		for _, r := range results {
			fmt.Println(formatter.FormatResult(r))
		}
	}
}
