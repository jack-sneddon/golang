package main

import "fmt"

func main() {
	// declare
	favoriteTeams := []string{"Milwaukee Bucks", favoriteNFLTeam()}
	fmt.Println(favoriteTeams)

	// add
	favoriteTeams = append(favoriteTeams, "BYU Cougars")
	//append does not modify the slice, it creates a new variable with the added on slice
	for i, team := range favoriteTeams {
		fmt.Println(i, team)
	}

}

func favoriteNFLTeam() string {
	return "Green Bay Packers"
}
