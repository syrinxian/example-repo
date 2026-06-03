# example-repo

Throwaway repository used by the AgentPatron end-to-end smoke test
(`make e2e` in the AgentPatron repo). Contains a tiny Go module with a
`strings` package the coding agent can extend, plus a CI workflow that
runs `go test ./...`.

Not intended for production use.

## Package `strings`

**Module path:** `github.com/syrinxian/example-repo`  
**Minimum Go version:** 1.22

Import the package with:

```go
import "github.com/syrinxian/example-repo/strings"
```

### Exported functions

| Signature | Description |
|-----------|-------------|
| `Repeat(s string, n int) string` | Returns `s` concatenated `n` times; returns `""` when `n <= 0`. |
| `Reverse(s string) string` | Returns `s` with its runes in reversed order, correctly handling multi-byte UTF-8 characters. |
| `Append(s, suffix string) string` | Returns `s` with `suffix` concatenated onto the end. |

