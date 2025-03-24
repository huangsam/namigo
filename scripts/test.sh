#!/bin/bash
set -eu

mode="${1:-default}"

selector=(
    "github.com/huangsam/namigo/pkg/..."
    "github.com/huangsam/namigo/internal/..."
)

case "$mode" in
    "default")
        go test "${selector[@]}" ;;
    "bench")
        go test -bench=. "${selector[@]}" ;;
    "cover")
        go test -cover "${selector[@]}" ;;
    *)
        echo "Invalid mode '$mode' detected" && exit 1 ;;
esac
