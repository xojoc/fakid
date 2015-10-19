// Written by http://xojoc.pw. Public Domain.

package main

type country int

const (
	usa country = iota
	ncountries
)

func (l country) String() string {
	return map[country]string{
		usa: "USA",
	}[l]
}
