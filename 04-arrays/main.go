package main

import "fmt"

func main() {

	// byte array
	byteArr := []byte("my_array")
	fmt.Println("byte array =", byteArr)
	str1 := string(byteArr[:])
	fmt.Println("back to string = ", str1)

	// empty array
	playArrays()

	// slices
	playSlices()

}

func favoriteNFLTeam() string {
	return "Green Bay Packers"
}
