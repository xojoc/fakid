// Written by http://xojoc.pw. Public Domain.

package main

import (
	"math/rand"
	"strconv"
	"strings"
)

// https://en.wikipedia.org/wiki/Disposable_email_address

var emailDomains = []string{
	"mailinator.com",
}

// The more unique the better.
func newEmail(first, last string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z' ||
			r >= 'A' && r <= 'Z' ||
			r >= '0' && r <= '9' ||
			strings.ContainsRune("._-@", r):
			return r
		default:
			return -1
		}
	},
		strings.ToLower(first+"."+last+strconv.Itoa(rand.Intn(100))+"@"+emailDomains[rand.Intn(len(emailDomains))]))
}
