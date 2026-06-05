package strings

import (
	"strconv"
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

func TestRandomNumberString(t *testing.T) {
	// Deterministic edge cases
	cases := []struct {
		max int
	}{
		{0},  // only valid result is "0"
		{-5}, // negative max clamped to 0; only valid result is "0"
	}
	for _, c := range cases {
		got := RandomNumberString(c.max)
		n, err := strconv.Atoi(got)
		if err != nil {
			t.Errorf("RandomNumberString(%d) = %q; not a valid integer", c.max, got)
		}
		clampedMax := c.max
		if clampedMax < 0 {
			clampedMax = 0
		}
		if n < 0 || n > clampedMax {
			t.Errorf("RandomNumberString(%d) = %d; out of range [0, %d]", c.max, n, clampedMax)
		}
	}

	// Non-deterministic: run 50 iterations with max=100 and verify range
	for i := 0; i < 50; i++ {
		got := RandomNumberString(100)
		n, err := strconv.Atoi(got)
		if err != nil {
			t.Fatalf("RandomNumberString(100) = %q; not a valid integer", got)
		}
		if n < 0 || n > 100 {
			t.Errorf("RandomNumberString(100) = %d; want value in [0, 100]", n)
		}
	}
}

