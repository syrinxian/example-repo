# example-repo

A small throwaway Go module containing a `strings` utility package with
helper functions. Includes a CI workflow that runs `go test ./...`.

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
| `RandomLetters(r *rand.Rand, n int) string` | Returns a string of `n` randomly chosen lowercase ASCII letters (`'a'`–`'z'`); returns `""` when `n <= 0`. Pass a seeded `*rand.Rand` for reproducible results. |

