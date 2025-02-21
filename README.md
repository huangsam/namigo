# Namigo

Your naming pal, written in Go 🐶

> It's a dog. It's a friend. It's a Namigo!

[Click here](./docs/approach.md) to learn more about the implementation details.

<img src="./images/namigo.jpeg" alt="Namigo" width="250px" />

## Getting started

```shell
# Run build
bash scripts/build.sh

# Run linters
bash scripts/lint.sh
```

### Running binary

```shell
# Get help
./bin/namigo help

# Search for package matches
./bin/namigo search package 'hello'

# Search for DNS matches
./bin/namigo search dns 'hello'
```

## Code structure

This codebase closely resembles [golang-standards/project-layout].

The TLDR is:

- `cmd` has entry points
- `pkg` has name generation and discovery logic
- `internal` has helpers

[golang-standards/project-layout]: https://github.com/golang-standards/project-layout
