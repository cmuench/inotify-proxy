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

## Git Commits Standards

This project follows the "Conventional Commit" standard.

- **Feature branches**: Should be prefixed with `feature/`.
- **Bugfix branches**: Should be prefixed with `bugfix/`.
- **General fixes**: Should be prefixed with `fix/`.
- **Pull Request Title**: The title of a Pull Request should explain the purpose of the changes in a short manner.

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
