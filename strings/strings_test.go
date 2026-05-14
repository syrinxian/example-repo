package strings

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	cases := []struct {
		in   string
		n    int
		want string
	}{
		{"", 3, ""},
		{"ab", 0, ""},
		{"ab", -1, ""},
		{"ab", 3, "ababab"},
	}
	for _, c := range cases {
		if got := Repeat(c.in, c.n); got != c.want {
			t.Errorf("Repeat(%q, %d) = %q, want %q", c.in, c.n, got, c.want)
		}
	}
}

func TestReverse(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},               // empty string
		{"hello", "olleh"},     // ASCII
		{"héllo", "olléh"},     // multi-byte UTF-8
		{"racecar", "racecar"}, // palindrome
	}
	for _, c := range cases {
		if got := Reverse(c.in); got != c.want {
			t.Errorf("Reverse(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAppend(t *testing.T) {
	cases := []struct {
		s, suffix, want string
	}{
		{"", "", ""},
		{"hello", "", "hello"},
		{"", "world", "world"},
		{"hello", " world", "hello world"},
		{"こんにちは", "世界", "こんにちは世界"},
	}
	for _, c := range cases {
		if got := Append(c.s, c.suffix); got != c.want {
			t.Errorf("Append(%q, %q) = %q, want %q", c.s, c.suffix, got, c.want)
		}
	}
}

func TestReplaceFirst(t *testing.T) {
	cases := []struct {
		s, old, new, want string
	}{
		// old not present
		{"hello", "xyz", "!", "hello"},
		// empty s
		{"", "a", "b", ""},
		// single occurrence
		{"hello world", "world", "Go", "hello Go"},
		// multiple occurrences — only first replaced
		{"aaa", "a", "b", "baa"},
		// empty old (inserts at start)
		{"hello", "", "X", "Xhello"},
		// Unicode
		{"こんにちは世界", "世界", "Go", "こんにちはGo"},
		// replacement is empty (deletion of first occurrence)
		{"foobar", "foo", "", "bar"},
	}
	for _, c := range cases {
		if got := ReplaceFirst(c.s, c.old, c.new); got != c.want {
			t.Errorf("ReplaceFirst(%q, %q, %q) = %q, want %q", c.s, c.old, c.new, got, c.want)
		}
	}
}
