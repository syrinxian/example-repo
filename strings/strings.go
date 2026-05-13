// Package strings provides small string utility helpers used by the
// AgentPatron end-to-end smoke test repository.
package strings

// Repeat returns s concatenated n times. If n is <= 0 it returns "".
func Repeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	out := ""
	for i := 0; i < n; i++ {
		out += s
	}
	return out
}
