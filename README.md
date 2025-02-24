# Namigo

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/huangsam/namigo/ci.yml)](https://github.com/huangsam/namigo/actions)
[![License](https://img.shields.io/github/license/huangsam/namigo)](https://github.com/huangsam/namigo/blob/main/LICENSE)

Your naming pal, written in Go 🐶

> It's a dog. It's a friend. It's a Namigo!

<img src="./images/namigo.jpeg" alt="Namigo" width="250px" />

## Motivation

Finding the right name for projects or packages is often harder than it seems.
Namigo solves this "naming block" by providing a simple tool to quickly search
for available names across package repositories and DNS, letting you focus
on building.

## Design docs

[Click here](docs/generate_approach.md) to learn about the `generate` implementation.

[Click here](docs/search_approach.md) to learn about the `search` implementation.

## Getting started

```shell
# Run build
bash scripts/build.sh

# Run tests
bash scripts/test.sh

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

# Search for email matches
./bin/namigo search email 'hello'
```

## Code structure

This codebase closely resembles [golang-standards/project-layout].

The TLDR is:

- `cmd` has entry points
- `pkg` has business logic
- `internal` has helpers

[golang-standards/project-layout]: https://github.com/golang-standards/project-layout
