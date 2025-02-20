package util

import (
	"fmt"

	"github.com/huangsam/namigo/internal/model"
)

func PrintResults(results any, label string, format func(any) string) {
	switch res := results.(type) {
	case []model.GoPackageResult:
		if len(res) > 0 {
			fmt.Printf("%d %s results found:\n", len(res), label)
			for _, r := range res {
				fmt.Println(format(r))
			}
		}
	case []model.NPMPackageResult:
		if len(res) > 0 {
			fmt.Printf("%d %s results found:\n", len(res), label)
			for _, r := range res {
				fmt.Println(format(r))
			}
		}
	case []model.PyPIPackageResult:
		if len(res) > 0 {
			fmt.Printf("%d %s results found:\n", len(res), label)
			for _, r := range res {
				fmt.Println(format(r))
			}
		}
	case []model.DNSResult:
		if len(res) > 0 {
			fmt.Printf("%d %s results found:\n", len(res), label)
			for _, r := range res {
				fmt.Println(format(r))
			}
		}
	}
}
