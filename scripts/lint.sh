#!/bin/bash
set -eu

# default: Run golangci-lint for all packages.
# fix: Run golangci-lint and fix issues.
# format: Run golangci-lint fmt to format the code.
mode="${1:-default}"

case "$mode" in
    default)
        golangci-lint run ;;
    fix)
        golangci-lint run --fix ;;
    format)
        golangci-lint fmt ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
