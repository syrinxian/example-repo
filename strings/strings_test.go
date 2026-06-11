package strings

import (
	"math/rand"
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

func TestRandomLetters(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	cases := []struct {
		n         int
		wantLen   int
		wantEmpty bool
	}{
		{0, 0, true},
		{-1, 0, true},
		{1, 1, false},
		{5, 5, false},
		{26, 26, false},
	}

	for _, c := range cases {
		got := RandomLetters(r, c.n)
		if c.wantEmpty {
			if got != "" {
				t.Errorf("RandomLetters(r, %d) = %q, want \"\"", c.n, got)
			}
			continue
		}
		if len(got) != c.wantLen {
			t.Errorf("RandomLetters(r, %d) = %q, want length %d", c.n, got, c.wantLen)
			continue
		}
		for i, ch := range got {
			if ch < 'a' || ch > 'z' {
				t.Errorf("RandomLetters(r, %d) = %q, rune %d (%q) is not in range 'a'-'z'", c.n, got, i, ch)
			}
		}
	}

	// Verify that two different seeds produce at least one distinct result across a sample.
	r1 := rand.New(rand.NewSource(1))
	r2 := rand.New(rand.NewSource(99999))
	distinct := false
	for i := 0; i < 50; i++ {
		if RandomLetters(r1, 1) != RandomLetters(r2, 1) {
			distinct = true
			break
		}
	}
	if !distinct {
		t.Error("RandomLetters with two different seeds produced identical results for 50 consecutive calls; expected at least one difference")
	}
}
