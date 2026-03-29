# Contributing to Namigo 🐶

Thanks for your interest! We aim for a lean, high-velocity contribution process.

## Pull Requests

1. **Branch**: `git checkout -b <type>/<name>` (e.g., `feature/my-extra-search`)
2. **Implement**: Add tests and follow [Namigo Agentic Guidelines](AGENTS.md).
3. **Verify**: Run `bash scripts/test.sh` and `bash scripts/lint.sh`.
4. **Submit**: Open a PR with a concise description of your changes.

## Issues

For bugs or feature requests, [open an issue](https://github.com/huangsam/namigo/issues) with reproduction steps and expected outcomes.

## Architecture & Design

Namigo follows the standard [Go Project Layout](https://github.com/golang-standards/project-layout):
- `cmd/`: CLI entry points.
- `pkg/`: Public search and generation logic.
- `internal/`: Private core infrastructure and data models.
