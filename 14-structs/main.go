package main

import "fmt"

func main() {
	// declare
	favoriteTeams := teams{"Milwaukee Bucks", favoriteNFLTeam()}
	fmt.Println(favoriteTeams)

	// add
	favoriteTeams = append(favoriteTeams, "BYU Cougars")
	//append does not modify the slice, it creates a new variable with the added on slice

	favoriteTeams.print()

	playStructs()
}

func favoriteNFLTeam() string {
	return "Green Bay Packers"
}

type FamilyMember struct {
	name string
	age  int
	pet  bool
}

func newFamilyMemmber(name string, age int, pet bool) *FamilyMember {
	fm := FamilyMember{name, age, pet}
	return &fm
}

func playStructs() {
	fmt.Println("============")

	// construct directly
	bob := FamilyMember{"Bob", 17, false}
	alice := FamilyMember{name: "Alice", pet: false}

	// construct using method
	fido := newFamilyMemmber("Fido", 5, true)

	// from another instance
	bobPrime := &bob

	fmt.Println(bob)
	fmt.Println(alice)
	fmt.Println(fido)
	fmt.Println(bobPrime)
	bobPrime.age = 41
	fmt.Println(bobPrime)

}
