# AI Agent Instructions

This document provides instructions for AI agents working with this codebase.

## Project Purpose

`inotify-proxy` is a tool that detects file changes in Docker Containers or Virtual Machines. It's designed to run for extended periods and includes a garbage collector to manage memory usage.

## Project Structure

The project is structured as a standard Go application.

- `inotify-proxy.go`: The main application file.
- `internal/`: Contains the internal packages of the application.
- `go.mod`: The Go module definition file.
- `go.sum`: The Go module checksum file.
- `Makefile`: Contains build and test automation.
- `.github/workflows/`: Contains the GitHub Actions workflows.

## Standards

- **Go Fmt**: All Go code should be formatted with `go fmt`.
- **Go Vet**: All Go code should be checked with `go vet`.
- **Go Lint**: All Go code should be checked with `golint`.

## Setup Project

To set up the project, you need to have Go installed. Then, you can run the following commands:

```bash
go get -v -t -d ./...
```

## Testing

To run the tests, you can use the following command:

```bash
go test -v ./...
```

## Build Project

To build the project, you can use the following command:

```bash
go build -v .
```

---

## Git Branch names

Branch names should follow the format `type/description`, where `type` indicates the nature of the work (e.g., `feature`, `bugfix`, `hotfix`, `chore`) and `description` is a short, descriptive name of the change.

Use english words, lowercase, and hyphens to separate words. Avoid using spaces or special characters like hashes.

## Git Commit Message Instructions

This project recommends using the [Conventional Commit](https://www.conventionalcommits.org/) format for all commit messages. This helps keep the commit history readable and enables automated tools for changelogs and releases.

### Commit Message Structure

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

- **type**: The kind of change (e.g., `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`)
- **optional scope**: A section of the codebase affected (e.g., `cache`, `command`, `docs`)
- **description**: Short summary of the change (imperative, lower case, no period)

### Examples

- `feat: add user login functionality`
- `fix(cache): correct total price calculation`
- `docs: update README with installation steps`

### Optional Body

Use the body to provide additional context about the change.

### Optional Footer

Use the footer to reference issues or describe breaking changes.

```
BREAKING CHANGE: changes the API of the cache command

Closes #123
```
