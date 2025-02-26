# Namigo

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/huangsam/namigo/ci.yml)](https://github.com/huangsam/namigo/actions)
[![License](https://img.shields.io/github/license/huangsam/namigo)](https://github.com/huangsam/namigo/blob/main/LICENSE)

Your naming pal, written in Go 🐶

> It's a dog. It's a friend. It's a Namigo!

<img src="./images/namigo.jpeg" alt="Namigo" width="250px" />

## Motivation

Finding the right name for projects or packages is often harder than it seems.
Namigo solves this "naming block" by providing a simple tool to quickly search
for available names across package repositories and DNS, letting you focus on
building. By automating the search process, Namigo saves you time and
effort, ensuring that you can find the perfect name without the usual hassle.

Additionally, Namigo can generate new names via AI prompts to chatbots, offering
creative and unique suggestions. This feature leverages the power of artificial
intelligence to provide you with a variety of name options that you might not have
considered. Whether you need a name for a new project, package, or domain, Namigo
is your go-to naming pal.

## Design

[Click here](docs/search_approach.md) to learn about the `search` implementation.

[Click here](docs/generate_approach.md) to learn about the `generate` implementation.

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
# Install binary
go install github.com/huangsam/namigo/cmd/namigo@latest

# Get help
namigo help

# Generate a prompt to get names from AI
namigo generate prompt

# Search for package matches
namigo search package 'hello'

# Search for DNS matches
namigo search dns 'hello'

# Search for email matches
namigo search email 'hello'
```

## Code structure

This codebase closely resembles [golang-standards/project-layout].

The TLDR is:

- `cmd` has entry points
- `pkg` has business logic
- `internal` has helpers

[golang-standards/project-layout]: https://github.com/golang-standards/project-layout
