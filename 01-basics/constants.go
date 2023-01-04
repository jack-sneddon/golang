package main

import "fmt"

func playConstants() {
	// constants
	fmt.Println("\n\n***constants")
	const myConstNum1 = 1000

	// assigning a const from calculation
	const myConstNum2 = myConstNum1 / 10
	fmt.Printf("%v %v\n", myConstNum1, myConstNum2)
}
