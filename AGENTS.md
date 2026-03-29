# Namigo Agentic Guidelines

Namigo leverages specialized, concurrent **Agents** to query name availability and generate content across diverse platforms.

## Architecture

1.  **Agents** (`pkg/`): Independent workers (e.g., `npm`, `dns`, `pypi`) that implement platform-specific logic and interfaces.
2.  **Core** (`internal/core/`): Standardized infrastructure for networking, parallelization (worker pools), and document parsing.
3.  **Models** (`internal/model/`): A common domain language (e.g., `SearchKey`, `Portfolio`) that standardizes IO across all agents.

## Behavioral Cues

### Modern Go Syntax
We prefer **Go 1.22+** idioms. Avoid legacy patterns where modern alternatives exist:
- Use `for range count` instead of standard for-loops for simple repetitions.
- Leverage the `slices` and `maps` packages for collection manipulation.
- Use `errors.Join` for multi-error handling.
- Prefer generics for shared core utilities where type safety is paramount.
- Run `go fix ./...` to automatically update deprecated APIs to modern equivalents.
- Use `golangci-lint run` and `golangci-lint fmt` to maintain code quality and formatting.

### Extensibility
To add a new search or generation extension:
1.  **Logic**: Implement the core logic in a new package under `pkg/search/` or `pkg/generate/`.
2.  **Model**: Ensure the result complies with an existing `internal/model/` or define a new one if necessary.
3.  **CLI**: Register the new command in `cmd/namigo/main.go` and its corresponding action in `cmd/namigo/sub/`.

### Manual Verification
Test agents in isolation using the CLI before aggregating results:
- **Search**: `go run cmd/namigo/main.go search package <term>`
- **DNS**: `go run cmd/namigo/main.go search dns <domain>`
- **Email**: `go run cmd/namigo/main.go search email <address>`

### Concurrency & Resilience
- **Parallelism**: Always use `internal/core/parallel.go` to dispatch concurrent tasks. Do not manually manage goroutine lifecycles for agents.
- **Fail-Fast**: Ensure agents respect timeouts and handle API rate limits gracefully using shared core HTTP clients.
