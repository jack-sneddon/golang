package main

import "fmt"

func main() {
	// call a function
	favoriteNBATeam := favNBATeam()
	fmt.Println(favoriteNBATeam)

	// call a function in a local package
	favoriteNFLTeam := favoriteNFLTeam()
	fmt.Println(favoriteNFLTeam)
}

func favNBATeam() string {
	return "Milwaukee Bucks"
}
