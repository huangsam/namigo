# Agents

This document describes the various agents implemented in Namigo, based on the code in `cmd`, `internal`, and `pkg` directories. Agents are responsible for performing specific search or generation tasks.

## Overview

Namigo uses a modular architecture with agents for different types of searches and generation. The agents are orchestrated through the CLI in `cmd/namigo/sub/` and utilize core utilities from `internal/core/` and data models from `internal/model/`.

## Search Agents

Search agents query external services to check name availability across different platforms. Each agent implements a specific search strategy and returns structured results.

### Golang Agent (`pkg/search/golang/`)

- **Purpose**: Searches for available Go package names on pkg.go.dev.
- **Method**: Web scraping using goquery to parse HTML from pkg.go.dev search results.
- **Key Function**: `SearchByScrape(name, size)` - Scrapes search snippets and filters packages containing the search term in name or path.
- **Data Model**: Returns `model.GoPackage` with name, path, and description.
- **Dependencies**: Uses `internal/core` for HTTP requests and document processing.

### NPM Agent (`pkg/search/npm/`)

- **Purpose**: Searches for available NPM package names on npmjs.com.
- **Method**: API queries to registry.npmjs.com for package listings.
- **Key Function**: `SearchByAPI(name, size)` - Queries the NPM registry API and parses JSON responses.
- **Data Model**: Returns `model.NPMPackage` with name and description.
- **Dependencies**: Uses `internal/core` for REST API queries and `internal/model/extern` for API response structures.

### PyPI Agent (`pkg/search/pypi/`)

- **Purpose**: Searches for available Python package names on pypi.org.
- **Method**: API queries to pypi.org, first listing projects then fetching details for each.
- **Key Function**: `SearchByAPI(name, size)` - Lists projects matching the prefix, then fetches detailed info for each using concurrent workers.
- **Data Model**: Returns `model.PyPIPackage` with name, author, and description.
- **Dependencies**: Uses `internal/core` for REST API queries and concurrent processing, `internal/model/extern` for API response structures.

### DNS Agent (`pkg/search/dns/`)

- **Purpose**: Checks DNS availability for domain names across common TLDs.
- **Method**: DNS lookups using Go's net package to resolve IP addresses.
- **Key Function**: `SearchByProbe(name, size)` - Probes multiple domains (e.g., name.com, name.org) concurrently using workers.
- **Data Model**: Returns `model.DNSRecord` with FQDN and associated IP addresses.
- **Dependencies**: Uses `internal/core` for concurrent worker management.

### Email Agent (`pkg/search/email/`)

- **Purpose**: Validates email address syntax and domain availability.
- **Method**: Email syntax verification and MX record lookups.
- **Key Function**: `SearchByProbe(name, size)` - Checks syntax using email-verifier library and validates domains via MX lookups, with rate limiting.
- **Data Model**: Returns `model.EmailRecord` with address, syntax validity, and domain validity.
- **Dependencies**: Uses external `email-verifier` library and Go's net package for MX lookups.

## Generate Agent

### Prompt Generator (`pkg/generate/`)

- **Purpose**: Generates AI prompts for creative name brainstorming.
- **Method**: Template-based prompt construction using embedded template.
- **Key Function**: `Prompt(purpose, theme, demographics, interests, size, length)` - Fills a template with user inputs to create a structured prompt for AI chatbots.
- **Data Model**: Returns a formatted string prompt.
- **Dependencies**: Uses Go's text/template package and embedded project.template file.

## Internal Components

### Core (`internal/core/`)

Provides shared utilities for HTTP requests, concurrent processing, input handling, and data pipelines. Includes functions like `RESTAPIQuery`, `StartCommonWorkers`, and `NewDocumentPipeline`.

### Model (`internal/model/`)

Defines data structures and enums for search results, including `SearchKey` enum (GoKey, NPMKey, PyPIKey, DNSKey, EmailKey) and record types for each agent. Also includes external API response models in `extern/`.

## Orchestration

Search agents are orchestrated in `cmd/namigo/sub/search.go` using `SearchRunner` and `SearchPortfolio` from `pkg/search/portfolio.go`. The portfolio aggregates results from multiple agents and handles formatting/output.</content>
<parameter name="filePath">/Users/samhuang/Playground/projects/namigo/AGENTS.md
