# Contribution guidelines

Welcome to the Namigo project! We appreciate your interest in contributing.
This document outlines the process for contributing to the project, including
guidelines for code structure, commands, and more. Please follow these guidelines
to ensure a smooth and efficient collaboration.

## Design

[Click here](docs/search_approach.md) to learn about the `search` implementation.

[Click here](docs/generate_approach.md) to learn about the `generate` implementation.

## Commands

```shell
# Run build
bash scripts/build.sh

# Run tests
bash scripts/test.sh

# Run linters
bash scripts/lint.sh
```

## Code structure

This codebase closely resembles [golang-standards/project-layout].

The TLDR is:

- `cmd` has entry points
- `pkg` has business logic
- `internal` has helpers

[golang-standards/project-layout]: https://github.com/golang-standards/project-layout
