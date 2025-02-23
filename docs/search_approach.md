# Search Approach

This document outlines a methodology for searching and discovering information
related to online presence and software package availability.
It covers DNS, email address verification, and software package metadata from
popular repositories for Go, JavaScript, Python, Rust, and Java.

## DNS

- **`dig`:**
    - Retrieve DNS records (A, MX, TXT, NS).
    - Use `+trace`, `+short`, and analyze TTL/DNSSEC.
- **`whois`:**
    - Get domain registration info.
    - Check different `whois` servers and historical data.
- **TLD Variation:**
    - Check `.io`, `.ai`, etc. for availability/patterns.

## Email

- **Lookup:**
    - Verify email existence.
    - Use verification services, check social profiles, headers, breaches.
- **Provider Variation:**
    - Check `@yahoo.com`, `@outlook.com`, etc.
    - Use provider APIs, analyze spam scores, check for disposable emails.

## Go

- **pkg.go.dev:**
    - Scrape metadata, popularity, dependencies.
- **index.golang.org:**
    - Search via API, analyze results.

## JavaScript (npm)

- **npmjs.com:**
    - Scrape metadata, downloads, vulnerabilities.
- **registry.npmjs.com/:package:**
    - Get package info via API.
- **registry.npmjs.com/-/v1/search:**
    - Search via API, analyze results.

## Python (PyPI)

- **pypi.org:**
    - Scrape metadata, downloads, vulnerabilities.
- **pypi.org/pypi/:package/json:**
    - Get package info via API.

## Rust (Crates.io)

- **crates.io:**
    - Scrape metadata, downloads, vulnerabilities.
- **crates.io/api/v1/crates/:crate:**
    - Get crate info via API.
- **crates.io/api/v1/crates?q=:**
    - Search via API, analyze results.

## Java

- **Maven/Gradle:**
    - Scrape metadata, dependencies, build info.
    - Use APIs, analyze build files, Javadoc.
    - Check GitHub/GitLab, use `jdeps`, vulnerability tools.
