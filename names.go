// Written by http://xojoc.pw. Public Domain.

package main

type sex int

const (
	male sex = iota
	female
)

func (s sex) String() string {
	switch s {
	case male:
		return "Male"
	case female:
		return "Female"
	}
	panic("unreachable")
}

var firstNames [ncountries][female + 1][]string
var lastNames [ncountries][]string
