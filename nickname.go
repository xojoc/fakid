// Written by http://xojoc.pw. Public Domain.

package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Reservoir sampling http://www.cs.umd.edu/~samir/498/vitter.pdf
func nickname(first, last string) string {
	s := []rune(first + last)
	n := min(len(s), 4+rand.Intn(4))
	nick := make([]rune, n)
	copy(nick, s[:n])
	for i := n; i < len(s); i++ {
		j := rand.Intn(i)
		if j < n {
			nick[j] = s[i]
		}
	}
	return strings.ToLower(string(nick) + strconv.Itoa(rand.Intn(100)))
}
