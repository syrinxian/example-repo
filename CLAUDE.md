# CLAUDE.md

This file provides guidance to Claude-based coding agents working with this
repository.

## Repository overview

`example-repo` is a small throwaway Go module used by the AgentPatron
end-to-end smoke test. It contains a single `strings` package with four
exported helper functions.

## Module information

| Property | Value |
|----------|-------|
| **Module path** | `github.com/syrinxian/example-repo` |
| **Minimum Go version** | 1.22 |

## Package layout

```
.
├── go.mod
├── strings/
│   ├── strings.go        # Exported helper functions
│   └── strings_test.go   # Table-driven unit tests
├── CLAUDE.md             # This file
├── README.md
└── LICENSE.md
```

## Build & test commands

```bash
# Build all packages
go build ./...

# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with race detector
go test -race ./...
```

## Package `strings`

Import path: `github.com/syrinxian/example-repo/strings`

### Exported functions

| Signature | Description |
|-----------|-------------|
| `Repeat(s string, n int) string` | Returns `s` concatenated `n` times; returns `""` when `n <= 0`. |
| `Reverse(s string) string` | Returns `s` with its runes in reversed order; correctly handles multi-byte UTF-8 characters. |
| `Append(s, suffix string) string` | Returns `s` with `suffix` concatenated onto the end. |
| `RickyRoll(s string) string` | Ignores the input string and returns the YouTube URL for the classic Rick Astley "Never Gonna Give You Up" music video. |

## Coding conventions

- **Table-driven tests**: All tests use `[]struct{ ... }` case tables and a
  single loop; follow this pattern when adding new tests.
- **UTF-8 safety**: Functions that iterate over string characters should
  operate on `[]rune`, not `[]byte`, to handle multi-byte characters
  correctly.
- **No external dependencies**: The module has no third-party dependencies;
  keep it that way unless there is a compelling reason.
- **Go standard formatting**: Run `gofmt` (or `goimports`) before committing.
  There is no separate linter config — standard `go vet ./...` is sufficient.
