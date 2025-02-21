# Namigo

[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/huangsam/namigo/ci.yml)](https://github.com/huangsam/namigo/actions)
[![License](https://img.shields.io/github/license/huangsam/namigo)](https://github.com/huangsam/namigo/blob/main/LICENSE)

Your naming pal, written in Go 🐶

> It's a dog. It's a friend. It's a Namigo!

[Click here](./docs/approach.md) to learn more about the implementation details.

<img src="./images/namigo.jpeg" alt="Namigo" width="250px" />

## Motivation

Choosing the right name for your projects, packages, or even DNS entries
can be surprisingly challenging. We often find ourselves stuck, brainstorming
for hours, only to come up with names that are either already taken or don't
quite capture the essence of our work.

Namigo was born out of this very frustration. It aims to streamline the naming
process by providing a simple and efficient tool to search for available names
across different domains, such as package repositories and DNS. Whether you're
a seasoned developer or just starting out, Namigo is designed to be your friendly
companion, helping you find the perfect name quickly and easily. 

Essentially, Namigo seeks to alleviate the "naming block" by offering a practical
solution to explore and validate potential names, allowing you to focus on building
rather than just brainstorming.

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
- `pkg` has business logic
- `internal` has helpers

[golang-standards/project-layout]: https://github.com/golang-standards/project-layout
