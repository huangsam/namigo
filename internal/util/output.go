package util

import (
	"fmt"

	"encoding/json"

	"github.com/huangsam/namigo/internal/model"
)

type OutputMode int

const (
	TextMode OutputMode = iota
	JSONMode
)

func (o OutputMode) String() string {
	switch o {
	case TextMode:
		return "PlainText"
	case JSONMode:
		return "JSON"
	default:
		return "Unknown"
	}
}

func PrintResults(results any, label string, format func(any) string, mode OutputMode) {
	switch res := results.(type) {
	case []model.GoPackageResult:
		if len(res) > 0 {
			printResults(res, label, format, mode)
		}
	case []model.NPMPackageResult:
		if len(res) > 0 {
			printResults(res, label, format, mode)
		}
	case []model.PyPIPackageResult:
		if len(res) > 0 {
			printResults(res, label, format, mode)
		}
	case []model.DNSResult:
		if len(res) > 0 {
			printResults(res, label, format, mode)
		}
	}
}

func printResults[T any](results []T, label string, format func(any) string, mode OutputMode) {
	switch mode {
	case JSONMode:
		jsonData, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Printf("Cannot print %s for %s: %v\n", mode, label, err)
			return
		}
		fmt.Printf("%s: %s\n", label, jsonData)
	case TextMode:
		fmt.Printf("%d %s results found:\n", len(results), label)
		for _, r := range results {
			fmt.Println(format(r))
		}
	}
}
