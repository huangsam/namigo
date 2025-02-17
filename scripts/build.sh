#!/bin/bash
set -eu

root="$(git rev-parse --show-toplevel)"
cmd="namigo"

go build -o "$root/bin/$cmd" "github.com/huangsam/namigo/cmd/$cmd"
