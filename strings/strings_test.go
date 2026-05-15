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

func TestRickRoll(t *testing.T) {
	cases := []string{
		"",
		"hello world",
		"こんにちは",
		"https://example.com",
		"never gonna give you up",
	}
	for _, c := range cases {
		if got := RickRoll(c); got != RickRollURL {
			t.Errorf("RickRoll(%q) = %q, want %q", c, got, RickRollURL)
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

