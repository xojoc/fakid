// Written by http://xojoc.pw. Public Domain.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var bloodTypes = []string{"O", "A", "B", "AB"}
var bloodSign = []string{"-", "+"}

func bloodType() string {
	return bloodTypes[rand.Intn(len(bloodTypes))] + bloodSign[rand.Intn(len(bloodSign))]
}

var likesList = []string{"Ice cream", "Kitties", "Staying up all night"}
func likes() string {
	return likesList[rand.Intn(len(likesList))]
}

func hates() string {
	return ""
}

var colors = []string{"red", "blue", "green"}
func color() string {
	return colors[rand.Intn(len(colors))]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cntry := country(rand.Intn(int(ncountries)))
	sex := sex(rand.Intn(2))
	first := firstNames[cntry][sex][rand.Intn(len(firstNames[cntry][sex]))]
	last := lastNames[cntry][rand.Intn(len(lastNames[cntry]))]
	profession := professions[rand.Intn(len(professions))]

	fmt.Println("First name: " + first)
	fmt.Println("Last name: " + last)
	fmt.Println("Nickname: " + nickname(first, last))
	fmt.Println("Sex: ", sex)
	fmt.Println("Country: ", cntry)
	fmt.Println("Profession: ", profession)
	fmt.Println("Email: ", newEmail(first, last))
	fmt.Println("Passphrase: ", passphrase())
	fmt.Println("Blood type: ", bloodType())
	fmt.Println("Likes: ", likes())
	fmt.Println("Hates: ", hates())
	fmt.Println("Favorite color: ", color())
}
