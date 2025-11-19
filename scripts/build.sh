#!/bin/bash
set -eu

root="$(pwd)"
cmd="namigo"

go build -o "$root/bin/$cmd" "github.com/huangsam/namigo/v2/cmd/$cmd"
