# Namigo

Your naming pal, written in Go 🐶

## Getting started

```shell
# Run build
bash scripts/build.sh

# Run linters
bash scripts/lint.sh
```

## Code structure

This codebase closely resembles [golang-standards/project-layout].

The TLDR is:

- `cmd` has entry points
- `pkg` has name generation and discovery logic
- `internal` has helpers

[golang-standards/project-layout]: https://github.com/golang-standards/project-layout
