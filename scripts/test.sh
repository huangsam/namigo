#!/bin/bash
set -eu

# default: Run tests for all packages with caching.
# bench: Run benchmarks for all packages.
# cover: Run tests and report coverage for all packages.
mode="${1:-default}"

selector=(
    "github.com/huangsam/go-trial/pkg/..."
    "github.com/huangsam/go-trial/internal/..."
)

case "$mode" in
    default)
        go test "${selector[@]}" ;;
    bench)
        go test -bench=. "${selector[@]}" ;;
    cover)
        go test -cover "${selector[@]}" ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
