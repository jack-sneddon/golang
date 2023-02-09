package main

import "fmt"

func playArrays() {
	fmt.Println("======Arrays of int")

	// int array with fixed size of 3
	var mySetOfNumbers [3]int
	mySetOfNumbers[0] = 100
	mySetOfNumbers[1] = 101
	mySetOfNumbers[2] = 102
	fmt.Printf("\nmySetOfNumbers = %v", mySetOfNumbers)
	fmt.Printf("\nmySetOfNumber[0] %v", mySetOfNumbers[0])
	fmt.Printf("\nmySetOfNumber[1] %v", mySetOfNumbers[1])
	fmt.Printf("\nmySetOfNumber[2] %v", mySetOfNumbers[2])

	// split an array
	x := []int{5, 2, 3, -9, 4, 2}

	// [5,2,4]
	y := x[:len(x)/2]
	// [-9,4,2]
	z := x[len(x)/2:]

	fmt.Printf("\n\n***Split Array \norig: %v\nx = %v  \ny=%v ", x, y, z)

	// declare -
	fmt.Println("\n\n***Arrays of strings")
	favoriteTeams := []string{"Milwaukee Bucks", favoriteNFLTeam()}
	fmt.Printf("\n%v", favoriteTeams)

	// add
	favoriteTeams = append(favoriteTeams, "BYU Cougars")
	//append does not modify the slice, it creates a new variable with the added on slice
	for i, team := range favoriteTeams {
		fmt.Println(i, team)
	}

	// multi-dimentional array
	fmt.Println("======Multi-dimentional Arrays")
	// multi dim int array with fixed size of 3, 2
	var multiArray [3][2]int
	multiArray[0][0] = 0
	multiArray[0][1] = 1
	multiArray[1][0] = 2
	multiArray[1][1] = 3
	multiArray[2][0] = 4
	multiArray[2][1] = 5
	fmt.Printf("\nmulti=dim array: %v\n", multiArray)

}
