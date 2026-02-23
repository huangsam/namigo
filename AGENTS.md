# Agents & Architecture

Namigo's architecture is built around the concept of **Agents**. Rather than having a monolithic search function, Namigo utilizes specialized agents that know how to query specific platforms, registries, or infrastructure to determine name availability.

## Motivation

When creating a new project, startup, or brand, developers usually need to check name availability across multiple platforms:
1. Is the `.com` domain available?
2. Is the Go module path free on `pkg.go.dev`?
3. Is the NPM package name taken?
4. Is it registered on PyPI?

Checking these manually is tedious. Namigo automates this by dispatching independent agents to perform these checks concurrently, aggregating the results into a unified summary. This agent-based design ensures that Namigo is extensible; if we want to add support for a new registry (like RubyGems or crates.io), we simply implement a new agent.

## Architecture

Namigo's architecture is divided into three main layers:

### 1. Agents (`pkg/search/`, `pkg/generate/`)
Agents are the specialized workers that interface with external systems.
- **Search Agents**: Implement concurrent querying and parsing for specific platforms.
  - `golang`: Scrapes `pkg.go.dev` for existing packages.
  - `npm`: Queries the `registry.npmjs.com` REST API.
  - `pypi`: Interacts with the Python Package Index API.
  - `dns`: Probes standard TLDs for domain availability using DNS resolution.
  - `email`: Validates syntax and checks MX records for email deliverability.
- **Generate Agents**: (e.g., `pkg/generate`) Focus on creative tasks, such as combining user inputs into structured prompts for AI chatbots to brainstorm name ideas.

### 2. Core Utilities (`internal/core/`)
Instead of each agent reinventing the wheel, they rely on a shared core library. This ensures consistent networking and performance.
- **Concurrency**: Worker pools and pipelines for hitting APIs without overwhelming them.
- **Networking**: Standardized HTTP clients, robust retries, and rate limiting.
- **Document Processing**: Shared utilities for parsing HTML/JSON.

### 3. Data Models (`internal/model/`)
A shared vocabulary (`SearchKey`, `GoPackage`, `NPMPackage`, etc.) ensures that regardless of how an agent retrieves its data (REST API, direct DNS probe, or HTML scraping), the resulting format is returned in a predictable, strongly-typed domain model.

### Orchestration
The CLI (`cmd/namigo/`) acts as the conductor. It reads the user's input, instantiates the requested agents, and dispatches them. A `Portfolio` struct aggregates the asynchronous results from the agents and formats them for the user's terminal output.
