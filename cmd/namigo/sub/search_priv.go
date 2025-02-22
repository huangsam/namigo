package sub

import (
	"encoding/json"
	"fmt"

	"github.com/huangsam/namigo/internal/model"
	"github.com/huangsam/namigo/pkg/search"
)

// display prints results.
func display(results any, formatter search.Formatter, mode model.OutputMode) {
	switch res := results.(type) {
	case []model.GoPackageResult:
		if len(res) > 0 {
			displayTypeMode(res, formatter, mode)
		}
	case []model.NPMPackageResult:
		if len(res) > 0 {
			displayTypeMode(res, formatter, mode)
		}
	case []model.PyPIPackageResult:
		if len(res) > 0 {
			displayTypeMode(res, formatter, mode)
		}
	case []model.DNSResult:
		if len(res) > 0 {
			displayTypeMode(res, formatter, mode)
		}
	}
}

// displayTypeMode prints results based on type and output mode.
func displayTypeMode[T any](results []T, formatter search.Formatter, mode model.OutputMode) {
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
			fmt.Println(formatter.Format(r))
		}
	}
}
