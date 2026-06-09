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

func TestRandomLetter(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	// Call RandomLetter 100 times and verify each result is a single lowercase letter.
	cases := make([]struct{ call int }, 100)
	for i := range cases {
		cases[i].call = i
	}
	for _, c := range cases {
		got := RandomLetter(r)
		if len(got) != 1 {
			t.Errorf("call %d: RandomLetter() = %q, want length 1", c.call, got)
			continue
		}
		ch := rune(got[0])
		if ch < 'a' || ch > 'z' {
			t.Errorf("call %d: RandomLetter() = %q, rune %q is not in range 'a'-'z'", c.call, got, ch)
		}
	}

	// Verify that two different seeds produce at least one distinct letter across a sample.
	r1 := rand.New(rand.NewSource(1))
	r2 := rand.New(rand.NewSource(99999))
	distinct := false
	for i := 0; i < 50; i++ {
		if RandomLetter(r1) != RandomLetter(r2) {
			distinct = true
			break
		}
	}
	if !distinct {
		t.Error("RandomLetter with two different seeds produced identical letters for 50 consecutive calls; expected at least one difference")
	}
}

