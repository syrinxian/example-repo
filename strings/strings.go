// Package strings provides small string utility helpers.
package strings

import (
	"math/rand"
	"strconv"
)

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

// Append returns s with suffix concatenated onto the end.
func Append(s, suffix string) string {
	return s + suffix
}

// RandomNumberString returns a random integer in the range [0, max] as a
// decimal string. If max is less than 0 it is treated as 0.
func RandomNumberString(max int) string {
	if max < 0 {
		max = 0
	}
	return strconv.Itoa(rand.Intn(max + 1))
}

// RandomLetters returns a string of n randomly chosen lowercase ASCII letters
// ('a'–'z'). When n <= 0 it returns "".
// Pass a seeded *rand.Rand for reproducible results; passing nil panics.
func RandomLetters(r *rand.Rand, n int) string {
	if n <= 0 {
		return ""
	}
	buf := make([]byte, n)
	for i := range buf {
		buf[i] = byte('a' + r.Intn(26))
	}
	return string(buf)
}

// firstNames is the pool from which FirstName draws.
var firstNames = []string{
	"Alice", "Bob", "Carol", "David", "Eve",
	"Frank", "Grace", "Hank", "Ivy", "Jack",
	"Karen", "Leo", "Mia", "Nate", "Olivia",
	"Pete", "Quinn", "Rose", "Sam", "Tina",
}

// FirstName returns a randomly chosen first name from a built-in list.
// Pass a seeded *rand.Rand for reproducible results; passing nil panics.
func FirstName(r *rand.Rand) string {
	return firstNames[r.Intn(len(firstNames))]
}
