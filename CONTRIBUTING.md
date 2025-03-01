# Welcome to Namigo! 🐶

Welcome to the Namigo project! Thanks for your interest in contributing.
This document outlines the process for contributing to the project, including
guidelines for reporting issues, submitting pull requests, and coding style.

Additionally, you will find information on the design documents and the overall
code structure to help you get started with understanding the project.

## Reporting issues

If you find a bug or have a feature request, please open an issue on GitHub. When reporting a bug, please include:

* A clear and descriptive title
* Steps to reproduce the bug
* The expected behavior
* The actual behavior
* Your operating system and version

## Pull request process

1.  Fork the repository
2.  Create a feature branch (`git checkout -b feature/my-feature`)
3.  Make your changes
4.  Run tests and linters (`bash scripts/test.sh`, `bash scripts/lint.sh`)
5.  Commit your changes
6.  Push to your branch (`git push origin feature/my-feature`)
7.  Open a pull request

Please ensure your pull request adheres to the following:

* Include tests for new features and bug fixes
* Update documentation as needed
* Follow the coding style guidelines

## Design docs

[Click here](docs/search_approach.md) to learn about the `search` implementation.

[Click here](docs/generate_approach.md) to learn about the `generate` implementation.

## Code structure

This codebase closely resembles [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

The TLDR is:

- `cmd` has entry points
- `pkg` has business logic
- `internal` has helpers
