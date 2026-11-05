# go-playground

A Go playground project for learning and experimenting with the Go language.

## Project Structure

| Directory | Purpose |
|-----------|---------|
| `jar/`    | Experiments with extracting Java JAR file contents (manifests, entries, etc.). |
| `jar/`    | Experiments with Java JAR file contents extraction (manifests, entries, etc.). |
| `lang/`   | Experiments with core Go language features — slices, maps, named returns, nil handling, and more. |
| `rt/`     | Experiments with the `runtime` standard library package. |
| `util/`   | Miscellaneous utility functions and helpers. |

## Goals

- Understand idiomatic Go patterns and practices.
- Explore core language constructs through small, focused experiments.
- Experiment with external packages (e.g., JAR file handling via `jar`).
- Learn the standard library (`runtime`, etc.).

## Tooling (!!!this tools should not be launched by agents!!!)

- **staticcheck** — Static analysis; run before committing to catch issues early.
- `go test ./...` — Run all tests across every package.

## Conventions

- One purpose per file; keep experiments small and self-contained.
- Always include a corresponding `_test.go` file with examples/tests.
