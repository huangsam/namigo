# Welcome to Namigo! 🐶

Welcome to the Namigo project! We appreciate your interest in contributing.
This document outlines the process for contributing to the project, including
guidelines for reporting issues, submitting pull requests, and coding style.

## Reporting Issues

If you find a bug or have a feature request, please open an issue on GitHub. When reporting a bug, please include:

* A clear and descriptive title
* Steps to reproduce the bug
* The expected behavior
* The actual behavior
* Your operating system and version

## Pull Request Process

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

## Development Setup

To set up your development environment:

```shell
# Run build
bash scripts/build.sh

# Run tests
bash scripts/test.sh

# Run linters
bash scripts/lint.sh
```
