package strings

import (
	"math/rand"
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

func TestRandomLetters(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	cases := []struct {
		n        int
		wantLen  int
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

