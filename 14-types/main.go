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
}

func favoriteNFLTeam() string {
	return "Green Bay Packers"
}
