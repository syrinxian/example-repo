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

// Reverse returns s with its runes in reversed order, correctly handling
// multi-byte UTF-8 characters.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
